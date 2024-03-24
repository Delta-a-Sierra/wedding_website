package domain

import (
	context "context"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

type RegistryRepo interface {
	GetRegistryItems(context.Context) ([]entities.RegistryItem, error)
}
