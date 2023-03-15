package main

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
		var err error
		commandName := i.Interaction.ApplicationCommandData().Name
		log.Println("Received interaction event", i.Interaction)
		if commandName == "dice-calc" {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    "",
					Flags:      discordgo.MessageFlagsEphemeral,
					Components: Components(),
				},
			})
		} else if commandName == "dice-roll" {
			var expression string
			for _, option := range i.Interaction.ApplicationCommandData().Options {
				if option.Name == "expression" {
					expression = option.StringValue()
					break
				}
			}
			if expression == "" {
				log.Panicln("ERROR: cannot roll an empty expression")
			}
			respondWithRollResult(expression, roller, s, i)
		}
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
					Content:    formula,
					Components: i.Message.Components,
				},
			})
		} else if interactionId == "roll" {
			expression := i.Message.Content
			log.Println("Rolling ", expression)
			// todo: also delete original ephemeral message. Caveat: ephemeral messages cannot
			//  be deleted like normal messages.
			respondWithRollResult(expression, roller, s, i)
		} else if interactionId == "AC" {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					// Fixme: sending a placeholder string because empty strings are not allowed by Discord API
					Content:    emptyEmbedContentPlaceholder,
					Components: i.Message.Components,
				},
			})
		} else if interactionId == BackspaceCommandId {
			originalMessage := i.Message.Content
			trimmedMessage := originalMessage
			if originalMessage != emptyEmbedContentPlaceholder {
				trimmedMessage = originalMessage[:len(originalMessage)-1]
				if trimmedMessage == "" {
					trimmedMessage = emptyEmbedContentPlaceholder
				}
			}
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    trimmedMessage,
					Components: i.Message.Components,
				},
			})
		} else if interactionId == "-" || interactionId == "+" || interactionId == "*" {
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    i.Message.Content + interactionId,
					Components: i.Message.Components,
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

func respondWithRollResult(
	expression string,
	roller parser.Antrl4BasedRoller,
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) error {
	log.Println("Rolling ", expression)
	rollResult, _, rollerErr := roller.Roll(expression)
	var response string
	if rollerErr != nil {
		log.Println("ERROR: ", rollerErr)
		response = "Invalid query (" + rollerErr.Error() + ")"
	} else {
		log.Println("INFO: ", rollResult)
		response = rollResult.String()
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "<@" + i.Interaction.Member.User.ID + "> is rolling " + expression + ": " + response,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Users:       []string{i.Interaction.Member.User.ID},
				RepliedUser: true,
			},
		},
	})
	return err
}
