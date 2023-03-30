package main

import (
	"flag"
	"log"
	"runtime/debug"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/justinian/dice"
	common "github.com/petertrr/dice-calc-bot/bot-common"
)

var (
	Token string
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
		log.Println("error creating Telegram session: ", err)
		return
	}

	c := &common.CommonBotContext{}
	c.Setup()

	commands := tgbotapi.NewSetMyCommands(tgbotapi.BotCommand{
		Command:     "roll",
		Description: "Evaluates dice notation, e.g. `2d20+4`.",
	})

	_, err = bot.Request(commands)
	if err != nil {
		log.Println("error registering bot commands: ", err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go process(updates, bot, c)

	common.WaitForGracefulShutdown()
	bot.StopReceivingUpdates()
}

func process(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, ctx *common.CommonBotContext) {
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			// process only Command updates
			processUpdate(update, bot, ctx)
		}
	}
}

func processUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, ctx *common.CommonBotContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("ERROR: error processing update %v\n%v", r, string(debug.Stack()))
		}
	}()

	if update.Message.Command() == "roll" {
		// processing
		text := update.Message.CommandArguments()
		var result dice.RollResult
		if text != "" {
			result, _, _ = ctx.Roller.Roll(text)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "@"+update.SentFrom().UserName+" rolled "+result.String())
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
