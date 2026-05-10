package skills_test

import (
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/skills"
	"github.com/stretchr/testify/require"
)

func TestFeedbackSkill_Execute(t *testing.T) {
	log := slog.New(slog.DiscardHandler)

	expected := "Спасибо за ваш отзыв!"

	s := skills.NewFeedbackSkill(log)
	got, gotErr := s.Execute("test")
	require.Equal(t, expected, got)
	require.NoError(t, gotErr)
}
