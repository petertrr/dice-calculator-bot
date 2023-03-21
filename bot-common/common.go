package common

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/petertrr/dice-calc-bot/parser"
)

// var (
// 	Token string
// )

// func setup() {
// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// 	if Token == "" {
// 		log.Panicf("Token should be provided in order to access Discord API")
// 	}
// }

type CommonBotContext struct {
	Roller *parser.Antrl4BasedRoller
}

func (c *CommonBotContext) Setup() {
	rand.Seed(time.Now().UnixNano())
	roller := parser.NewAntrl4BasedRoller(
		func(x int) int { return rand.Intn(x) + 1 },
	)
	c.Roller = &roller
}

func WaitForGracefulShutdown() {
	log.Println("Bot is now running.")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Bot is shutting down.")
}
