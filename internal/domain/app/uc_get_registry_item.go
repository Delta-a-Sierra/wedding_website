package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a App) GetRegistryItem(ctx context.Context) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItems: %w", err)
	}
	return items, nil
}
