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

func ReportSuccess(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, span := otel.StartSpan(context.Background(), "command/ReportSuccess")
	defer span.End()

	err := s.MessageReactionAdd(m.ChannelID, m.Message.ID, "✅")
	if err != nil {
		otel.SpanError(span, err)
	}
}

func ReportFailure(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, span := otel.StartSpan(context.Background(), "command/ReportFailure")
	defer span.End()

	err := s.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
	if err != nil {
		otel.SpanError(span, err)
	}
}
