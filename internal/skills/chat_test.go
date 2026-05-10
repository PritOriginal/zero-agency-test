package skills_test

import (
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/skills"
	"github.com/stretchr/testify/require"
)

func TestChatSkill_Execute(t *testing.T) {
	log := slog.New(slog.DiscardHandler)

	expected := "Привет! Я ИИ-помощник, чем могу помочь?"

	s := skills.NewChatSkill(log)
	got, gotErr := s.Execute("test")
	require.Equal(t, expected, got)
	require.NoError(t, gotErr)
}
