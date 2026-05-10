//go:build integration

package integration

import (
	"context"
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/classifier"
	openaiclient "github.com/PritOriginal/zero-agency-test/internal/client/openai"
	"github.com/PritOriginal/zero-agency-test/internal/config"
	"github.com/PritOriginal/zero-agency-test/internal/shared/tags"
	"github.com/stretchr/testify/suite"
)

type ClassifierSuite struct {
	suite.Suite
	Cfg *config.Config
}

func (st *ClassifierSuite) SetupSuite() {
	st.Cfg = config.MustLoadPath("../../../configs/config.yaml")
}

func TestClassifierSuite(t *testing.T) {
	suite.Run(t, new(ClassifierSuite))
}

func (st *ClassifierSuite) TestClassify() {
	tests := []struct {
		name      string
		userInput string
		wantTag   string
	}{
		{
			name:      "Ok-" + tags.Chat,
			userInput: "Привет, как дела?",
			wantTag:   tags.Chat,
		},
		{
			name:      "Ok-" + tags.SupportUrgency,
			userInput: "Верните мои деньги!!",
			wantTag:   tags.SupportUrgency,
		},
		{
			name:      "Ok-" + tags.Feedback,
			userInput: "Отличный сервис",
			wantTag:   tags.Feedback,
		},
		{
			name:      "Ok-" + tags.InfoRequest,
			userInput: "Как осуществить запрос с помощью API?",
			wantTag:   tags.InfoRequest,
		},
	}
	for _, tt := range tests {
		st.Run(tt.name, func() {
			log := slog.New(slog.DiscardHandler)
			openAiClient := openaiclient.New(st.Cfg.OpenAI.Model, st.Cfg.OpenAI.URL, st.Cfg.OpenAI.ApiKey)
			classifierService := classifier.New(log, openAiClient, []string{
				tags.InfoRequest,
				tags.SupportUrgency,
				tags.Chat,
				tags.Feedback,
			})

			tag, err := classifierService.Classify(context.Background(), tt.userInput)
			st.NoError(err)
			st.Equal(tt.wantTag, tag)
		})
	}
}
