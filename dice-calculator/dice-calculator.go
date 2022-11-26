package dicecalculator

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/justinian/dice"
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
			formula := i.Message.Content
			if i.Message.Content != "" {
				formula += "+"
			}
			formula += "1" + strings.TrimPrefix(i.MessageComponentData().CustomID, "roll-")
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: formula,
				},
			})
			if err != nil {
				log.Println("ERROR: ", err)
			}
		} else if i.MessageComponentData().CustomID == "roll" {
			log.Println("Rolling ", i.Message.Content)
			rollResult, _, err := dice.Roll(i.Message.Content)
			var response string
			if err != nil {
				log.Println("ERROR: ", err)
				response = "Invalid query (" + err.Error() + ")"
			} else {
				log.Println("INFO: ", rollResult)
				response = strconv.Itoa(rollResult.Int()) + " (" + rollResult.Description() + ")"
			}
			// todo: also delete original ephemeral message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Result: " + response,
				},
			})
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
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Roll!",
					CustomID: "roll",
					Style:    discordgo.SuccessButton,
				},
			},
		},
	}
}
