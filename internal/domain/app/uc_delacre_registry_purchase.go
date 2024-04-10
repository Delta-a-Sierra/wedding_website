package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (a *App) DelcareRegistryPurchase(ctx context.Context, id uuid.UUID, guestID uuid.UUID) error {
	if err := a.repo.DeclareRegistryItemPurchase(ctx, id, guestID); err != nil {
		return fmt.Errorf("a.repo.DeclareRegistryItemPurchase: %w", err)
	}
	return nil
}
