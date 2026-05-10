package skills

import "log/slog"

type KnowledgeSkill struct {
	log *slog.Logger
}

func NewKnowledgeSkill(log *slog.Logger) *KnowledgeSkill {
	return &KnowledgeSkill{
		log: log,
	}
}

func (s *KnowledgeSkill) Execute(userInput string) (string, error) {
	s.log.Debug("Запущен навык: KnowledgeSkill")
	s.log.Debug("Делаю поиск по базе...")
	return "Вот что я нашел в документации: ...", nil
}
