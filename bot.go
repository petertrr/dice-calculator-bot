package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	dicecalculator "github.com/petertrr/dice-calc-bot/dice-calculator"
	"github.com/petertrr/dice-calc-bot/parser"
)

var (
	Token string
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "dice-calc",
			Description: "Show the dice calculator interface",
		},
	}
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	if Token == "" {
		log.Panicf("Token should be provided in order to access Discord API")
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	roller := parser.Antrl4BasedRoller{}
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("ERROR: MainInterfaceHandler has crashed", r)
			}
		}()
		dicecalculator.MainInterfaceHandler(roller, s, i)
	})

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
