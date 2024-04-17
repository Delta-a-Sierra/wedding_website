package domain

import (
	context "context"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
)

type RegistryRepo interface {
	GetRegistryItems(context.Context) ([]entities.RegistryItem, error)
	GetRegistryItemsPage(context.Context, int, int, func(entities.RegistryItem) bool) ([]entities.RegistryItem, error)
	GetRegistryItemsFiltered(context.Context, func(entities.RegistryItem) bool) ([]entities.RegistryItem, error)
	GetRegistryItemsNotPurchased(context.Context) ([]entities.RegistryItem, error)
	GetRegistryItem(context.Context, uuid.UUID) (entities.RegistryItem, error)
	DeclareRegistryItemPurchase(context.Context, uuid.UUID, uuid.UUID) error
	AddRegistryItem(context.Context, entities.RegistryItem) error
	EditRegistryItem(context.Context, entities.RegistryItem) error
	DeleteRegistryItem(context.Context, uuid.UUID) error
	SearchRegistry(string) ([]entities.RegistryItem, error)
	SearchRegistryNotPurchased(string) ([]entities.RegistryItem, error)
}
