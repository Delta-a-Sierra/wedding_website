package app

import (
	"context"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a *App) GetWeddingInfo(ctx context.Context) (entities.WeddingInfo, error) {
	return entities.WeddingInfo{}, nil
}
