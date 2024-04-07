package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a *App) GetGuestList(ctx context.Context) ([]entities.Guest, error) {
	guests, err := a.repo.GetGuestList(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetGuestList: %w", err)
	}
	return guests, nil
}
