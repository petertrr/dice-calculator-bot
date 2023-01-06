package dicecalculator

import "github.com/bwmarrin/discordgo"

func Components() []discordgo.MessageComponent {
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