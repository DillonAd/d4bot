package event

import (
	"context"
	"fmt"

	"github.com/DillonAd/d4bot/cmd/bot/eventhandler/event/haiku"
	"github.com/DillonAd/d4bot/cmd/otel"
	"github.com/bwmarrin/discordgo"
)

const EventNameHaiku = "haiku"

func Haiku(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, span := otel.StartSpan(context.Background(), "event/haiku")
	defer span.End()

	otel.AddHandlerAttributes(span, m)

	if m.Author.ID == s.State.User.ID {
		return
	}

	response, err := haiku.Format(m.Content)
	if err != nil {
		if err == haiku.ErrNotAHaiku {
			return
		}

		span.RecordError(err)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("㊗️ _Haiku Detected_ ㊗️\n>>> %s", response))
	if err != nil {
		otel.SpanError(span, err)
	}
}
