package dicecalculator

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/petertrr/dice-calc-bot/parser"
)

const emptyEmbedContentPlaceholder = "<Enter an expression>"

/**
 * @param roller fixme: no public interface for roller is exported from `dice`
 */
func MainInterfaceHandler(
	roller parser.Antrl4BasedRoller,
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	if i.Interaction.Type == discordgo.InteractionApplicationCommand {
		log.Println("Received interaction event", i.Interaction)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    "",
				Flags:      discordgo.MessageFlagsEphemeral,
				Components: Components(),
			},
		})
		if err != nil {
			log.Println("ERROR: ", err)
		}
	} else if i.Interaction.Type == discordgo.InteractionMessageComponent {
		log.Println("Received interaction event", i.Interaction)
		if strings.HasPrefix(i.MessageComponentData().CustomID, "roll-") {
			formula := i.Message.Content
			if formula == emptyEmbedContentPlaceholder {
				formula = ""
			}
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
			rollResult, _, err := roller.Roll(i.Message.Content)
			var response string
			if err != nil {
				log.Println("ERROR: ", err)
				response = "Invalid query (" + err.Error() + ")"
			} else {
				log.Println("INFO: ", rollResult)
				response = rollResult.String()
			}
			// todo: also delete original ephemeral message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Result: " + response,
				},
			})
		} else if i.MessageComponentData().CustomID == "AC" {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					// Fixme: sending a placeholder string because empty strings are not allowed by Discord API
					Content: emptyEmbedContentPlaceholder,
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
