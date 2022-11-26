package dicecalculator

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MainInterfaceHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Type == discordgo.InteractionApplicationCommand {
		log.Println("Received interaction event", i.Interaction)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    "",
				Flags:      discordgo.MessageFlagsEphemeral,
				Components: components(),
				// discordgo.TextInput{
				// Label: "Text",
				// },
			},
		})
		if err != nil {
			log.Println("ERROR: ", err)
		}
	} else if i.Interaction.Type == discordgo.InteractionMessageComponent {
		log.Println("Received interaction event", i.Interaction)
		if strings.HasPrefix(i.MessageComponentData().CustomID, "roll-") {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: i.Message.Content + "+" + strings.TrimPrefix(i.MessageComponentData().CustomID, "roll-"),
				},
			})
			if err != nil {
				log.Println("ERROR: ", err)
			}
		}
	} else {
		log.Println("Received unhandled interaction event", i.Interaction)
	}
}

func components() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Take a look inside",
					Style:    discordgo.LinkButton,
					Disabled: false,
					URL:      "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
					Emoji: discordgo.ComponentEmoji{
						Name: "🤷",
					},
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "d4",
					CustomID: "roll-d4",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "d6",
					CustomID: "roll-d6",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "d8",
					CustomID: "roll-d8",
					Style:    discordgo.SuccessButton,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "d10",
					CustomID: "roll-d10",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "d12",
					CustomID: "roll-d12",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "d20",
					CustomID: "roll-d20",
					Style:    discordgo.SuccessButton,
				},
			},
		},
	}
}
