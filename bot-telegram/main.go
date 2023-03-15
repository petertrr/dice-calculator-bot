package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/justinian/dice"
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
	rand.Seed(time.Now().UnixNano())
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Println("error creating Telegram session,", err)
		return
	}

	roller = parser.NewAntrl4BasedRoller(
		func(x int) int { return rand.Intn(x) + 1 },
	)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go process(updates, bot)

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Bot is shutting down.")
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
