package app

import (
	"context"
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

func (a App) GetCountdown(ctx context.Context) (entities.Countdown, error) {
	weddingDate, err := a.repo.GetWeddingDate(ctx)
	if err != nil {
		return entities.Countdown{}, fmt.Errorf("a.rep.GetWeddingDate: %w", err)
	}
	return entities.CreateCountDownFromDateTime(weddingDate), nil
}
