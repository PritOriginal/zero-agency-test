package bot

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"syscall"

	"github.com/PritOriginal/zero-agency-test/internal/classifier"
	httpclient "github.com/PritOriginal/zero-agency-test/internal/client/http"
	openaiclient "github.com/PritOriginal/zero-agency-test/internal/client/openai"
	"github.com/PritOriginal/zero-agency-test/internal/config"
	"github.com/PritOriginal/zero-agency-test/internal/router"
	"github.com/PritOriginal/zero-agency-test/internal/shared/tags"
	"github.com/PritOriginal/zero-agency-test/internal/skills"
	"github.com/PritOriginal/zero-agency-test/pkg/logger"
)

type Bot struct {
	log *slog.Logger
	r   *router.Router
}

func NewWithHttpClient(log *slog.Logger, cfg *config.Config) *Bot {
	httpClient := httpclient.New(cfg.OpenAI.Model, cfg.OpenAI.URL, cfg.OpenAI.ApiKey)
	return new(log, httpClient)
}

func NewWithOpenAiClient(log *slog.Logger, cfg *config.Config) *Bot {
	openAiClient := openaiclient.New(cfg.OpenAI.Model, cfg.OpenAI.URL, cfg.OpenAI.ApiKey)
	return new(log, openAiClient)
}

func new(log *slog.Logger, client classifier.Client) *Bot {
	classifierService := classifier.New(log, client, []string{
		tags.InfoRequest,
		tags.SupportUrgency,
		tags.Chat,
		tags.Feedback,
	})

	r := router.New(log, classifierService)
	r.RegisterSkill(tags.InfoRequest, skills.NewKnowledgeSkill(log))
	r.RegisterSkill(tags.SupportUrgency, skills.NewEscalationSkill(log))
	r.RegisterSkill(tags.Chat, skills.NewChatSkill(log))
	r.RegisterSkill(tags.Feedback, skills.NewFeedbackSkill(log))

	return &Bot{
		log: log,
		r:   r,
	}
}

func (b *Bot) Run() {
	pid := os.Getpid()
	b.log.Info("Бот запущен!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> Введите сообщение (Для выхода введите - exit): ")
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "" {
			continue
		}

		if userInput == "exit" {
			err := syscall.Kill(pid, syscall.SIGINT)
			if err != nil {
				log.Printf("Error sending signal: %v", err)
			}
			break
		}

		resp, err := b.r.Route(userInput)
		if err != nil {
			b.log.Error("Произошла ошибка:", logger.Err(err))
			fmt.Println("Бот: Произошла ошибка. Пожалуйста, повторите запрос")
		} else {
			fmt.Println("Бот:", resp)
		}
	}
}

func (b Bot) Stop() {
	b.log.Info("Бот остановлен")
}
