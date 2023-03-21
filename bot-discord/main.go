package main

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
	common "github.com/petertrr/dice-calc-bot/bot-common"
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
		{
			Name:        "dice-roll",
			Description: "Evaluate dice notation expression",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "expression",
					Description: "dice expression (e.g. `d20+5`)",
					Required:    true,
				},
			},
		},
	}
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	if Token == "" {
		log.Panicf("Token should be provided in order to access Discord API")
	}
}

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	c := common.CommonBotContext {}
	c.Setup()

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("ERROR: MainInterfaceHandler has crashed", r)
			}
		}()
		MainInterfaceHandler(c.Roller, s, i)
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
		log.Println("INFO: registered command [", v, "]")
		registeredCommands[i] = cmd
	}

	common.WaitForGracefulShutdown()
	discord.Close()
}
