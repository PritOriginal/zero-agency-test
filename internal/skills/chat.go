package skills

import "log/slog"

type ChatSkill struct {
	log *slog.Logger
}

func NewChatSkill(log *slog.Logger) *ChatSkill {
	return &ChatSkill{
		log: log,
	}
}

func (s *ChatSkill) Execute(userInput string) (string, error) {
	s.log.Debug("Запущен навык: ChatSkill")
	return "Привет! Я ИИ-помощник, чем могу помочь?", nil
}
