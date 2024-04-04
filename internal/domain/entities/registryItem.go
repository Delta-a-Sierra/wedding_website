package entities

import "github.com/google/uuid"

type RegistryItem struct {
	Name, Description, Link string
	ID                      uuid.UUID
	Purchased               bool
	Purchaser               uuid.UUID
}

func (r RegistryItem) GetID() uuid.UUID {
	return r.ID
}
