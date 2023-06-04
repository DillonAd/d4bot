package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/DillonAd/d4bot/cmd/bot"
	"github.com/DillonAd/d4bot/cmd/config"
	"github.com/DillonAd/d4bot/cmd/health"
	"github.com/DillonAd/d4bot/cmd/otel"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	shutdownTracing := otel.InitTracing(ctx, config.Tracing.OtelEndpoint)
	defer shutdownTracing()

	ready := make(chan bool)
	healthDone := health.Init(ctx, config, ready)

	bot, err := bot.New(ctx, config)
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
