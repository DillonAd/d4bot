package bot

import (
	"context"

	"github.com/DillonAd/d4bot/cmd/config"
)

type Bot struct {
}

func New(ctx context.Context, config *config.Config) (*Bot, error) {
	return &Bot{}, nil
}

func (b *Bot) Start() {
}

func (b *Bot) Close() {
}
