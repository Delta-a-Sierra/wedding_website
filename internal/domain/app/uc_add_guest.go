package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a *App) AddGuest(ctx context.Context, guest entities.Guest) (entities.Guest, error) {
	guest, err := a.repo.AddGuest(ctx, guest)
	if err != nil {
		return entities.Guest{}, fmt.Errorf("a.repo.AddGuest: %w", err)
	}
	return guest, nil
}
