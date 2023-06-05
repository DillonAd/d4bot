package eventhandler

import (
	"context"
	"fmt"
	"strings"

	"github.com/DillonAd/d4bot/cmd/otel"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func shoresy(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, span := otel.StartSpan(context.Background(), "eventhandler/shoresy")
	defer span.End()

	span.SetAttributes(
		attribute.String("username", m.Author.Username),
	)

	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.ToLower(m.Content)
	if !strings.Contains(content, "shoresy") {
		span.AddEvent("invalid event")
		return
	}
	msg := fmt.Sprintf("%s may like to win, but they don't hate to lose", m.Author.Username)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		span.SetStatus(codes.Error, "error sending message")
		span.RecordError(err)
	}
}
