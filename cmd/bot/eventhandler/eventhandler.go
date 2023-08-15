package eventhandler

import (
	"github.com/DillonAd/d4bot/cmd/bot/eventhandler/command"
)

func Handlers() map[string]interface{} {
	return map[string]interface{}{
		command.RollCommandName: command.Roll,
	}
}
