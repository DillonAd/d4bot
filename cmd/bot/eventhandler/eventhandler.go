package eventhandler

import (
	"github.com/DillonAd/d4bot/cmd/bot/eventhandler/command"
	"github.com/DillonAd/d4bot/cmd/bot/eventhandler/event"
)

func Handlers() map[string]interface{} {
	return map[string]interface{}{
		command.CommandNameRoll: command.Roll,
		event.EventNameHaiku:    event.Haiku,
	}
}
