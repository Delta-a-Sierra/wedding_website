package entities

import (
	"github.com/google/uuid"
)

type Guest struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Attending bool
}
