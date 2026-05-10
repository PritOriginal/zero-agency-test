package skills_test

import (
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/skills"
	"github.com/stretchr/testify/require"
)

func TestEscalationSkill_Execute(t *testing.T) {
	log := slog.New(slog.DiscardHandler)

	expected := "Переключаю на оператора"

	s := skills.NewEscalationSkill(log)
	got, gotErr := s.Execute("test")
	require.Equal(t, expected, got)
	require.NoError(t, gotErr)
}
