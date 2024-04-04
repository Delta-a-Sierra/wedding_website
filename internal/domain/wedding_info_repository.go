package domain

import (
	"context"
	"time"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
)

type WeddingInfoRepo interface {
	GetWeddingDate(context.Context) (time.Time, error)
	GetWeddingLocation(context.Context) (string, error)
	GetWeddingInfo(context.Context) (entities.WeddingInfo, error)
	GetGuestList(context.Context) ([]entities.Guest, error)
	AddGuest(context.Context, entities.Guest) error
	GetGuestByID(context.Context, uuid.UUID) (entities.Guest, error)
	GetGuestByEmail(context.Context, string) (entities.Guest, error)
	RSVP(context.Context, entities.Guest) error
}
