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
		interactionId := i.MessageComponentData().CustomID
		var err error = nil
		if strings.HasPrefix(interactionId, "roll-") {
			formula := i.Message.Content
			if formula == emptyEmbedContentPlaceholder {
				formula = ""
			}
			formula += strings.TrimPrefix(interactionId, "roll-")
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: formula,
				},
			})
		} else if interactionId == "roll" {
			log.Println("Rolling ", i.Message.Content)
			rollResult, _, rollerErr := roller.Roll(i.Message.Content)
			var response string
			if rollerErr != nil {
				log.Println("ERROR: ", rollerErr)
				response = "Invalid query (" + rollerErr.Error() + ")"
			} else {
				log.Println("INFO: ", rollResult)
				response = rollResult.String()
			}
			// todo: also delete original ephemeral message. Caveat: ephemeral messages cannot
			//  be deleted like normal messages.
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Result: " + response,
				},
			})
		} else if interactionId == "AC" {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					// Fixme: sending a placeholder string because empty strings are not allowed by Discord API
					Content: emptyEmbedContentPlaceholder,
				},
			})
		} else if interactionId == "-" || interactionId == "+" || interactionId == "*" {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: i.Message.Content + interactionId,
				},
			})
		}
		if err != nil {
			log.Println("ERROR: ", err)
		}
	} else {
		log.Println("Received unhandled interaction event", i.Interaction)
	}
}
