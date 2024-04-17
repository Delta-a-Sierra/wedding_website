package app

import (
	"context"
	"fmt"
	"math"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a App) GetRegistryItems(ctx context.Context) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItems: %w", err)
	}
	return items, nil
}

func (a App) GetRegistryItemsNotPurchased(ctx context.Context) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItemsNotPurchased(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItems: %w", err)
	}
	return items, nil
}

func (a App) GetRegistryItemsPageAll(ctx context.Context, count, offset int) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItemsPageAll(ctx, count, offset)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItemsPage: %w", err)
	}
	return items, nil
}

func (a App) GetRegistryItemsPageNotPurchased(ctx context.Context, count, offset int) ([]entities.RegistryItem, error) {
	items, err := a.repo.GetRegistryItemsPageNotPurchased(ctx, count, offset)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetRegistryItemsPage: %w", err)
	}
	return items, nil
}

func (a App) GetPageCount(ctx context.Context) (int, error) {
	i, err := a.GetRegistryItems(ctx)
	if err != nil {
		return 0, fmt.Errorf("a.GetRegistryItems: %w", err)
	}
	var count int
	if len(i) > 0 {
		count = int(math.Ceil(float64(len(i)) / float64(6)))
	}
	return count, nil
}

func (a App) GetPageCountNotPurchased(ctx context.Context) (int, error) {
	i, err := a.GetRegistryItemsNotPurchased(ctx)
	if err != nil {
		return 0, fmt.Errorf("a.GetRegistryItemsFiltered: %w", err)
	}
	var count int
	if len(i) > 0 {
		count = int(math.Ceil(float64(len(i)) / float64(6)))
	}
	return count, nil
}
