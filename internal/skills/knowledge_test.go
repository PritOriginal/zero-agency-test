package skills_test

import (
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/skills"
	"github.com/stretchr/testify/require"
)

func TestKnowledgeSkill_Execute(t *testing.T) {
	log := slog.New(slog.DiscardHandler)

	expected := "Вот что я нашел в документации: ..."

	s := skills.NewKnowledgeSkill(log)
	got, gotErr := s.Execute("test")
	require.Equal(t, expected, got)
	require.NoError(t, gotErr)
}
