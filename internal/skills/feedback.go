package skills

import "log/slog"

type FeedbackSkill struct {
	log *slog.Logger
}

func NewFeedbackSkill(log *slog.Logger) *FeedbackSkill {
	return &FeedbackSkill{
		log: log,
	}
}

func (s *FeedbackSkill) Execute(userInput string) (string, error) {
	s.log.Debug("Запущен навык: FeedbackSkill")
	s.log.Debug("Сохраняю в базу данных...")
	return "Спасибо за ваш отзыв!", nil
}
