package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a *App) RSVP(ctx context.Context, guest entities.Guest) error {
	if err := a.repo.RSVP(ctx, guest); err != nil {
		return fmt.Errorf("a.repo.RSVP: %w", err)
	}
	return nil
}
