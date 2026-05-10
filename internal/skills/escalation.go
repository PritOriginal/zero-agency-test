package skills

import "log/slog"

type EscalationSkill struct {
	log *slog.Logger
}

func NewEscalationSkill(log *slog.Logger) *EscalationSkill {
	return &EscalationSkill{
		log: log,
	}
}

func (s *EscalationSkill) Execute(userInput string) (string, error) {
	s.log.Debug("Запущен навык: EscalationSkill")
	s.log.Debug("Зову человека!")
	return "Переключаю на оператора", nil
}
