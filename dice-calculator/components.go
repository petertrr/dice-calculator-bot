package dicecalculator

import "github.com/bwmarrin/discordgo"

func Components() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
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
				discordgo.Button{
					Label:    "AC",
					CustomID: "AC",
					Style:    discordgo.DangerButton,
				},
				discordgo.Button{
					Label:    "dx",
					CustomID: "empty-dx",
					Style:    discordgo.SecondaryButton,
					// this button is TODO
					Disabled: true,
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
				discordgo.Button{
					Label:    "+",
					CustomID: "+",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					Label:    "(",
					CustomID: "empty-lpar",
					Style:    discordgo.SecondaryButton,
					// this button is TODO
					Disabled: true,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "7",
					CustomID: "roll-7",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "8",
					CustomID: "roll-8",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "9",
					CustomID: "roll-9",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "-",
					CustomID: "-",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					Label:    ")",
					CustomID: "empty-rpar",
					Style:    discordgo.SecondaryButton,
					// this button is TODO
					Disabled: true,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "4",
					CustomID: "roll-4",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "5",
					CustomID: "roll-5",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "6",
					CustomID: "roll-6",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "*",
					CustomID: "*",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					Label:    "0",
					CustomID: "roll-0",
					Style:    discordgo.SuccessButton,
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "1",
					CustomID: "roll-1",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "2",
					CustomID: "roll-2",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "3",
					CustomID: "roll-3",
					Style:    discordgo.SuccessButton,
				},
				discordgo.Button{
					Label:    "Roll!",
					CustomID: "roll",
					Style:    discordgo.PrimaryButton,
				},
			},
		},
	}
}
