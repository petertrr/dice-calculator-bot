package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/justinian/dice"
	common "github.com/petertrr/dice-calc-bot/bot-common"
	"github.com/petertrr/dice-calc-bot/parser"
)

var (
	Token  string
	roller parser.Antrl4BasedRoller
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	if Token == "" {
		log.Panicf("Token should be provided in order to access Telegram API")
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Println("error creating Telegram session,", err)
		return
	}

	c := common.CommonBotContext{}
	c.Setup()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go process(updates, bot)

	common.WaitForGracefulShutdown()
	bot.StopReceivingUpdates()
}

func process(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		processUpdate(update, bot)
	}
}

func processUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: error processing update", r)
		}
	}()
	// processing
	text := update.Message.Text
	var result dice.RollResult
	if text != "" {
		result, _, _ = roller.Roll(text)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "@"+update.SentFrom().UserName+" rolled "+result.String())
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
