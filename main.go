package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/DillonAd/d4bot/cmd/bot"
	"github.com/DillonAd/d4bot/cmd/health"
	"github.com/DillonAd/d4bot/cmd/otel"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	OtelEndpoint         string `split_words:"true"`
	OtelExporterInsecure bool   `split_words:"true"`
	BotToken             string `split_words:"true"`
}

func main() {
	var s Specification
	err := envconfig.Process("d4bot", &s)
	if err != nil {
		fmt.Println(err.Error())
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	shutdownTracing := otel.InitTracing(ctx, s.OtelEndpoint, s.OtelExporterInsecure)
	defer shutdownTracing()

	ready := make(chan bool)
	healthDone := health.Init(ctx, ready)

	bot, err := bot.New(ctx, s.BotToken)
	if err != nil {
		panic(err)
	}

	go bot.Start()
	defer bot.Close()
	ready <- true

	<-shutdown
	fmt.Println("starting shutdown")

	cancel()
	fmt.Println("cancelled context")

	<-healthDone
}
