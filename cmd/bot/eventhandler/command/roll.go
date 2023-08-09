package command

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/DillonAd/d4bot/cmd/otel"
	"github.com/bwmarrin/discordgo"
)

var RollCommandName string = "roll"

func Roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	spanCtx, span := otel.StartSpan(context.Background(), "command/roll")
	defer span.End()

	otel.AddHandlerAttributes(span, m)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !IsCommand(spanCtx, m, RollCommandName) {
		return
	}

	dieCount, diceSides, err := getDiceData(m.Content)
	if err != nil {
		helpResponse := "Invalid roll.\nFormat: `!roll {numberOfDice}d{numberOfDiceSides}`\nExample: `!roll 5d8`"
		_, err := s.ChannelMessageSend(m.ChannelID, helpResponse)
		if err != nil {
			otel.SpanError(span, err)
		}
		return
	}

	total := 0
	results := make([]string, dieCount)
	for i := 0; i < dieCount; i++ {
		result := rand.Intn(diceSides) + 1
		results[i] = strconv.Itoa(result)
		total += result
	}

	response := fmt.Sprintf("(`%s`)=`%d`", strings.Join(results, "`+`"), total)
	_, err = s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		otel.SpanError(span, err)
	}
}

func getDiceData(input string) (int, int, error) {
	r, _ := regexp.Compile(`\d+d\d+`)
	match := r.FindString(input)
	if match == "" {
		return -1, -1, fmt.Errorf("")
	}
	commandParts := strings.Split(match, "d")
	dieCount, _ := strconv.Atoi(commandParts[0])
	diceSides, _ := strconv.Atoi(commandParts[1])
	return dieCount, diceSides, nil
}
