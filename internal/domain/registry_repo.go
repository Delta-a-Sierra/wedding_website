package domain

import (
	context "context"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
)

type RegistryRepo interface {
	GetRegistryItems(context.Context) ([]entities.RegistryItem, error)
	GetRegistryItem(context.Context, uuid.UUID) (entities.RegistryItem, error)
	DeclareRegistryItemPurchase(context.Context, uuid.UUID, uuid.UUID) error
	AddRegistryItem(context.Context, entities.RegistryItem) error
	EditRegistryItem(context.Context, entities.RegistryItem) error
	DeleteRegistryItem(context.Context, uuid.UUID) error
}
