package command

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestIsCommand(t *testing.T) {
	// Assemble
	cases := []struct {
		Name    string
		Command *discordgo.MessageCreate
		IsValid bool
	}{
		{
			Name: "valid",
			Command: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "!roll 4d8",
				},
			},
			IsValid: true,
		},
		{
			Name: "invalid",
			Command: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "! roll 4d8",
				},
			},
			IsValid: false,
		},
		{
			Name: "valid ignore case",
			Command: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "!ROLL 4d8",
				},
			},
			IsValid: true,
		},
		{
			Name: "valid standalone",
			Command: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "!roll",
				},
			},
			IsValid: true,
		},
		{
			Name: "invalid standalone with suffix",
			Command: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "!rolldice",
				},
			},
			IsValid: false,
		},
	}

	for _, c := range cases {
		// Act
		result := IsCommand(context.Background(), c.Command, "roll")

		if (c.IsValid && !result) || (!c.IsValid && result) {
			t.Errorf("%s - expected %t, but got %t", c.Name, c.IsValid, result)
			continue
		}
	}
}
