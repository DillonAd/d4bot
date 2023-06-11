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
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	config_otel_endpoint := os.Getenv("OTEL_ENDPOINT")
	config_otel_insecure := os.Getenv("OTEL_EXPORTER_INSECURE") == "true"
	config_bot_token := os.Getenv("BOT_TOKEN")

	shutdownTracing := otel.InitTracing(ctx, config_otel_endpoint, config_otel_insecure)
	defer shutdownTracing()

	ready := make(chan bool)
	healthDone := health.Init(ctx, ready)

	bot, err := bot.New(ctx, config_bot_token)
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
