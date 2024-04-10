package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
)

type InMemory struct {
	weddingInfo   entities.WeddingInfo
	registryItems []entities.RegistryItem
	guestList     []entities.Guest
}

func NewInMemoryRepo() *InMemory {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	time, _ := time.Parse(layout, "May 4, 2024 at 1:00pm (GMT)")
	return &InMemory{
		weddingInfo: entities.WeddingInfo{
			Date:     time,
			Location: "15 bedford road, clapham, london, sw8 2hz",
		},
		registryItems: []entities.RegistryItem{
			{ID: uuid.New(), Name: "Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
			{ID: uuid.New(), Name: "Towels", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
			{ID: uuid.New(), Name: "Toaster", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
			{ID: uuid.New(), Name: "Kettle", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
			{ID: uuid.New(), Name: "Dyson Hoover", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
			{ID: uuid.New(), Name: "Herman Miller Office Chair", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		},
	}
}

func (i *InMemory) GetWeddingDate(context.Context) (time.Time, error) {
	return i.weddingInfo.Date, nil
}

func (i *InMemory) GetWeddingLocation(context.Context) (string, error) {
	return i.weddingInfo.Location, nil
}

func (i *InMemory) GetWeddingInfo(context.Context) (entities.WeddingInfo, error) {
	return i.weddingInfo, nil
}

func (i *InMemory) GetRegistryItems(context.Context) ([]entities.RegistryItem, error) {
	return i.registryItems, nil
}

func (i *InMemory) GetRegistryItemsNotPurchased(context.Context) ([]entities.RegistryItem, error) {
	var items []entities.RegistryItem
	for _, item := range i.registryItems {
		if item.Purchased {
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

func (i *InMemory) GetRegistryItem(ctx context.Context, id uuid.UUID) (entities.RegistryItem, error) {
	for _, item := range i.registryItems {
		if item.GetID() == id {
			return item, nil
		}
	}
	return entities.RegistryItem{}, domain.ErrRegistryItemNotFound
}

func (i *InMemory) DeclareRegistryItemPurchase(ctx context.Context, id uuid.UUID, guestID uuid.UUID) error {
	item, err := i.GetRegistryItem(ctx, id)
	if err != nil {
		return fmt.Errorf("i.GetRegistryItem: %w", err)
	}

	item.Purchased = true
	item.Purchaser = guestID

	if err := i.EditRegistryItem(ctx, item); err != nil {
		return fmt.Errorf("i.EditRegistryItem: %w", err)
	}
	return nil
}

func (i *InMemory) EditRegistryItem(ctx context.Context, item entities.RegistryItem) error {
	for j, it := range i.registryItems {
		if it.GetID() != item.GetID() {
			continue
		}
		i.registryItems[j] = item
		return nil
	}
	return domain.ErrRegistryItemNotFound
}

func (i *InMemory) AddRegistryItem(ctx context.Context, item entities.RegistryItem) error {
	item.ID = uuid.New()
	i.registryItems = append(i.registryItems, item)
	return nil
}

func (i *InMemory) DeleteRegistryItem(ctx context.Context, id uuid.UUID) error {
	for j, it := range i.registryItems {
		if it.GetID() != id {
			continue
		}

		if len(i.registryItems) < j+1 {
			i.registryItems = append(i.registryItems[:j], i.registryItems[j+1:]...)
		}
		i.registryItems = i.registryItems[:j]

		return nil
	}
	return domain.ErrRegistryItemNotFound
}

func (i *InMemory) GetGuestList(context.Context) ([]entities.Guest, error) {
	return i.guestList, nil
}

func (i *InMemory) AddGuest(ctx context.Context, guest entities.Guest) (entities.Guest, error) {
	g, err := i.GetGuestByEmail(ctx, guest.Email)
	if err == nil {
		return g, nil
	}
	guest.ID = uuid.New()
	i.guestList = append(i.guestList, guest)
	return guest, nil
}

func (i *InMemory) GetGuestByID(ctx context.Context, ID uuid.UUID) (entities.Guest, error) {
	for _, guest := range i.guestList {
		if guest.ID == ID {
			return guest, nil
		}
	}
	return entities.Guest{}, domain.ErrGuestNotFound
}

func (i *InMemory) GetGuestByEmail(ctx context.Context, email string) (entities.Guest, error) {
	for _, guest := range i.guestList {
		if guest.Email == email {
			return guest, nil
		}
	}
	return entities.Guest{}, domain.ErrGuestNotFound
}

func (i *InMemory) RSVP(ctx context.Context, guest entities.Guest) error {
	for j, g := range i.guestList {
		if g.ID != guest.ID {
			continue
		}
		g.Attending = true
		fmt.Println("Attending")
		i.guestList[j] = g
		return nil
	}
	return domain.ErrGuestNotFound
}

func (i *InMemory) SearchRegistry(searchString string) ([]entities.RegistryItem, error) {
	var items []entities.RegistryItem
	for _, item := range i.registryItems {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchString)) {
			items = append(items, item)
		}
	}
	fmt.Println("items", items)
	return items, nil
}

func (i *InMemory) SearchRegistryNotPurchased(searchString string) ([]entities.RegistryItem, error) {
	var items []entities.RegistryItem
	for _, item := range i.registryItems {
		if item.Purchased {
			continue
		}
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchString)) {
			items = append(items, item)
		}
	}
	fmt.Println("items", items)
	return items, nil
}
