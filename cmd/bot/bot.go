package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/DillonAd/d4bot/cmd/bot/eventhandler"
	"github.com/DillonAd/d4bot/cmd/config"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func New(ctx context.Context, config *config.Bot) (*Bot, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	if err != nil {
		return nil, err
	}
	bot := &Bot{
		session: session,
	}

	bot.registerEventHandlers()

	return bot, nil
}

func (b *Bot) Start() {
	err := b.session.Open()
	if err != nil {
		log.Printf("error starting bot: %v", err)
	}
}

func (b *Bot) registerEventHandlers() {
	for name, handler := range eventhandler.EventHandlers() {
		fmt.Printf("registering event handler: %s\n", name)
		_ = b.session.AddHandler(handler)
	}
}

func (b *Bot) Close() {
	err := b.session.Close()
	if err != nil {
		log.Printf("error closing bot: %v", err)
	}
}
