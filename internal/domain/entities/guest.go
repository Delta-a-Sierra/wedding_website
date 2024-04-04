package entities

import (
	"github.com/google/uuid"
)

type Guest struct {
	ID        uuid.UUID
	name      string
	Email     string
	Attending bool
}
