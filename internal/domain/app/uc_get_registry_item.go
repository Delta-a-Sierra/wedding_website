package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
)

func (a App) GetRegistryItem(ctx context.Context, id uuid.UUID) (entities.RegistryItem, error) {
	item, err := a.repo.GetRegistryItem(ctx, id)
	if err != nil {
		return entities.RegistryItem{}, fmt.Errorf("a.repo.GetRegistryItems: %w", err)
	}
	return item, nil
}
