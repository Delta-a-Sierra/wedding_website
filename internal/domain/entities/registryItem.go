package entities

import "github.com/google/uuid"

type RegistryItem struct {
	Name, Description, Link string
	ID                      uuid.UUID
	Price                   float64
}
