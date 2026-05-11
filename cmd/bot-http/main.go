package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PritOriginal/zero-agency-test/internal/app/bot"
	httpclient "github.com/PritOriginal/zero-agency-test/internal/client/http"
	"github.com/PritOriginal/zero-agency-test/internal/config"
	slogger "github.com/PritOriginal/zero-agency-test/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logger, err := slogger.SetupLogger(cfg.Env)
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	httpClient := httpclient.New(cfg.OpenAI.Model, cfg.OpenAI.URL, cfg.OpenAI.ApiKey)
	b := bot.New(logger, httpClient)

	go func() {
		b.Run()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done

	b.Stop()
}
