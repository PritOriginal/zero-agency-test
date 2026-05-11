package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PritOriginal/zero-agency-test/internal/app/bot"
	openaiclient "github.com/PritOriginal/zero-agency-test/internal/client/openai"
	"github.com/PritOriginal/zero-agency-test/internal/config"
	slogger "github.com/PritOriginal/zero-agency-test/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logger, err := slogger.SetupLogger(cfg.Env)
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	openAiClient := openaiclient.New(cfg.OpenAI.Model, cfg.OpenAI.URL, cfg.OpenAI.ApiKey)
	b := bot.New(logger, openAiClient)

	go func() {
		b.Run()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done

	b.Stop()
}
