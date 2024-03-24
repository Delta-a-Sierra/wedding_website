package domain

import (
	"context"
	"time"
)

type WeddingInfoRepo interface {
	GetWeddingDate(context.Context) (time.Time, error)
}
