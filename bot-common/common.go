package common

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/petertrr/dice-calc-bot/parser"
)

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
