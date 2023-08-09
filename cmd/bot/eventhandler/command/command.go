package command

import (
	"context"
	"strings"

	"github.com/DillonAd/d4bot/cmd/otel"
	"github.com/bwmarrin/discordgo"
)

func IsCommand(ctx context.Context, m *discordgo.MessageCreate, commandName string) bool {
	_, span := otel.StartSpan(context.Background(), "IsCommand")
	defer span.End()

	content := strings.ToLower(m.Content)
	prefix := strings.Split(content, " ")[0]
	commandFormat := "!" + commandName
	return prefix == commandFormat
}
