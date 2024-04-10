package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a *App) SearchRegistry(ctx context.Context, searchString string) ([]entities.RegistryItem, error) {
	items, err := a.repo.SearchRegistry(searchString)
	if err != nil {
		return nil, fmt.Errorf("a.repo.SearchRegistry: %w", err)
	}
	// if items == nil {
	// 	return a.repo.GetRegistryItems(ctx)
	// }
	return items, nil
}

func (a *App) SearchRegistryNotPurchased(ctx context.Context, searchString string) ([]entities.RegistryItem, error) {
	items, err := a.repo.SearchRegistryNotPurchased(searchString)
	if err != nil {
		return nil, fmt.Errorf("a.repo.SearchRegistryNotPurchased: %w", err)
	}
	// if items == nil {
	// 	return a.repo.GetRegistryItems(ctx)
	// }
	return items, nil
}
