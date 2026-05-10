package router

import (
	"context"
	"fmt"
	"log/slog"
)

type Skill interface {
	Execute(userInput string) (string, error)
}

type Classifier interface {
	Classify(ctx context.Context, userInput string) (string, error)
}

type Router struct {
	skills     map[string]Skill
	log        *slog.Logger
	classifier Classifier
}

func New(log *slog.Logger, c Classifier) *Router {
	return &Router{
		skills:     make(map[string]Skill),
		log:        log,
		classifier: c,
	}
}

func (r *Router) RegisterSkill(tag string, skill Skill) {
	r.skills[tag] = skill
}

func (r *Router) Route(userInput string) (string, error) {
	const op = "router.Router.Message"

	tag, err := r.classifier.Classify(context.Background(), userInput)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	r.log.Debug("ИИ определил категорию", slog.String("category", string(tag)))

	skill, ok := r.skills[tag]
	if !ok {
		return "", fmt.Errorf("%s: подходящий навык отсутствует", op)
	}
	resp, err := skill.Execute(userInput)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return resp, nil
}
