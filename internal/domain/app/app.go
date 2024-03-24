package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

type repo interface {
	domain.WeddingInfoRepo
	domain.RegistryRepo
}

type App struct {
	repo repo
}

func NewApp(repo repo) *App {
	return &App{
		repo: repo,
	}
}

func (a App) GetCountdown(ctx context.Context) (entities.Countdown, error) {
	weddingDate, err := a.repo.GetWeddingDate(ctx)
	if err != nil {
		return entities.Countdown{}, fmt.Errorf("a.rep.GetWeddingDate: %w", err)
	}
	return entities.CreateCountDownFromDateTime(weddingDate), nil
}

func (a App) GetRegistryItems(ctx context.Context) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItems: %w", err)
	}
	return items, nil
}
