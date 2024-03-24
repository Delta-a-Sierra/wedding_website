package repositories

import (
	"context"
	"time"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

type InMemory struct {
	date          time.Time
	registryItems []entities.RegistryItem
}

func NewInMemoryRepo() *InMemory {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	time, _ := time.Parse(layout, "May 4, 2024 at 1:00pm (GMT)")
	return &InMemory{
		date: time,
		registryItems: []entities.RegistryItem{
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
			{Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim", Price: 120.50},
		},
	}
}

func (i *InMemory) GetWeddingDate(context.Context) (time.Time, error) {
	return i.date, nil
}

func (i *InMemory) GetRegistryItems(context.Context) ([]entities.RegistryItem, error) {
	return i.registryItems, nil
}
