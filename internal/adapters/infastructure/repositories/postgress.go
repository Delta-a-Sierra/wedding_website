package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Postgress struct {
	db          *sql.DB
	weddingInfo entities.WeddingInfo
}

func NewPostgresRepo() *Postgress {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	time, _ := time.Parse(layout, "May 4, 2024 at 1:00pm (GMT)")
	return &Postgress{
		weddingInfo: entities.WeddingInfo{
			Date:     time,
			Location: "15 bedford road, clapham, london, sw8 2hz",
		},
	}
}

func (p *Postgress) Connect() error {
	username := os.Getenv("pgusername")
	password := os.Getenv("pgpassword")
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:5432/registry?sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("db.Ping: %w", err)
	}
	p.db = db
	if err = p.bootstrapTables(); err != nil {
		return fmt.Errorf("p.bootstrapTables: %w", err)
	}
	if os.Getenv("ww_environment") == "dev" {
		if err = p.seed(); err != nil {
			return fmt.Errorf("p.seed: %w", err)
		}
	}

	return nil
}

func (p *Postgress) bootstrapTables() error {
	fmt.Println("bootstrapping tables")
	_, err := p.db.Exec("CREATE TABLE IF NOT EXISTS items(id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name VARCHAR(100) NOT NULL, description VARCHAR(255) NOT NULL, link VARCHAR(255) NOT NULL);")
	if err != nil {
		return fmt.Errorf("p.db.Exec: %w", err)
	}
	_, err = p.db.Exec("CREATE TABLE IF NOT EXISTS guests(id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name VARCHAR(100) NOT NULL, email VARCHAR(255) NOT NULL);")
	if err != nil {
		return fmt.Errorf("p.db.Exec: %w", err)
	}
	fmt.Println("bootstrapped tables")
	return nil
}

func (p *Postgress) seed() error {
	fmt.Println("seeding db")
	_, err := p.db.Exec("DROP TABLE IF EXISTS items;")
	if err != nil {
		return fmt.Errorf("p.db.Exec: %w", err)
	}

	_, err = p.db.Exec("CREATE TABLE items(id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name VARCHAR(100) NOT NULL, description VARCHAR(255) NOT NULL, link VARCHAR(255) NOT NULL, purchased BOOLEAN NOT NULL DEFAULT FALSE);")
	if err != nil {
		return fmt.Errorf("p.db.Exec: %w", err)
	}
	registryItems := []entities.RegistryItem{
		{ID: uuid.New(), Name: "1 Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "2 Towels", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "3 Toaster", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "4 Kettle", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "5 Dyson Hoover", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "6 Herman Miller Office Chair", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "7 Le Creuset Cast Iron Signature Square Skillet", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "8 Towels", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "9 Toaster", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "10 Kettle", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "11 Dyson Hoover", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
		{ID: uuid.New(), Name: "12 Herman Miller Office Chair", Description: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim"},
	}
	for _, item := range registryItems {
		query := "INSERT INTO items (name, description, link)VALUES ($1, $2, $3);"
		_, err = p.db.Query(query, item.Name, item.Description, item.Link)
		if err != nil {
			return fmt.Errorf("p.db.Query: %w", err)
		}
	}
	fmt.Println("seeded db")
	return nil
}

func (p *Postgress) GetRegistryItems(ctx context.Context) ([]entities.RegistryItem, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, description, link, purchased FROM items ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("p.db.Query: %w", err)
	}
	items, err := parseRowsToItems(rows)
	if err != nil {
		return nil, fmt.Errorf("parseRowsToItems: %w", err)
	}
	return items, nil
}

func (p *Postgress) GetRegistryItemsPageAll(ctx context.Context, itemsPerPage int, page int) ([]entities.RegistryItem, error) {
	return p.getRegistryItemsPage(ctx, itemsPerPage, page, "")
}

func (p *Postgress) GetRegistryItemsPageNotPurchased(ctx context.Context, itemsPerPage int, page int) ([]entities.RegistryItem, error) {
	return p.getRegistryItemsPage(ctx, itemsPerPage, page, "not-purchased")
}

func (p *Postgress) getRegistryItemsPage(ctx context.Context, itemsPerPage int, page int, pageType string) ([]entities.RegistryItem, error) {
	var query string
	switch pageType {
	case "not-purchased":
		query = "SELECT id, name, description, link, purchased FROM items WHERE purchased=FALSE ORDER BY name OFFSET ($1-1)*$2 LIMIT $2;"
	default:
		query = "SELECT id, name, description, link, purchased FROM items ORDER BY name OFFSET ($1-1)*$2 LIMIT $2;"
	}
	rows, err := p.db.QueryContext(ctx, query, page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("p.db.QueryContext: %w", err)
	}
	items, err := parseRowsToItems(rows)
	if err != nil {
		return nil, fmt.Errorf("parseRowsToItems: %w", err)
	}
	return items, nil
}

func (p *Postgress) GetRegistryItemsNotPurchased(ctx context.Context) ([]entities.RegistryItem, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, description, link, purchased FROM items WHERE purchased = FALSE ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("p.db.Query: %w", err)
	}
	items, err := parseRowsToItems(rows)
	if err != nil {
		return nil, fmt.Errorf("parseRowsToItems: %w", err)
	}
	return items, nil
}

func (p *Postgress) GetRegistryItem(ctx context.Context, id uuid.UUID) (entities.RegistryItem, error) {
	row := p.db.QueryRowContext(ctx, "SELECT id, name, description, link, purchased FROM items ORDER BY name")
	var name, description, link string
	var purchased bool
	if err := row.Scan(&id, &name, &description, &link, &purchased); err != nil {
		return entities.RegistryItem{}, fmt.Errorf("rows.Scan: %w", err)
	}
	return entities.RegistryItem{ID: id, Name: name, Description: description, Link: link, Purchased: purchased}, nil
}

func (p *Postgress) DeclareRegistryItemPurchase(ctx context.Context, id uuid.UUID, purchaserID uuid.UUID) error {
	query := "UPDATE items SET purchased = TRUE WHERE id = $1;"
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("p.db.QueryContext: %w", err)
	}
	return nil
}

func (p *Postgress) AddRegistryItem(context.Context, entities.RegistryItem) error {
	return nil
}

func (p *Postgress) EditRegistryItem(context.Context, entities.RegistryItem) error {
	return nil
}

func (p *Postgress) DeleteRegistryItem(context.Context, uuid.UUID) error {
	return nil
}

func (p *Postgress) SearchRegistry(searchString string) ([]entities.RegistryItem, error) {
	query := `SELECT id, name, description, link, purchased FROM items WHERE name LIKE LOWER('%k%') ORDER BY name;`
	rows, err := p.db.QueryContext(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("p.db.QueryContext: %w", err)
	}
	items, err := parseRowsToItems(rows)
	if err != nil {
		return nil, fmt.Errorf("parseRowsToItems: %w", err)
	}
	return items, nil
}

func (p *Postgress) SearchRegistryNotPurchased(searchString string) ([]entities.RegistryItem, error) {
	return p.SearchRegistry(searchString)
}

func (p *Postgress) AddGuest(ctx context.Context, guest entities.Guest) (entities.Guest, error) {
	return entities.Guest{}, nil
}

func (p *Postgress) GetWeddingDate(context.Context) (time.Time, error) {
	return p.weddingInfo.Date, nil
}

func (p *Postgress) GetWeddingLocation(context.Context) (string, error) {
	return p.weddingInfo.Location, nil
}

func (p *Postgress) GetWeddingInfo(context.Context) (entities.WeddingInfo, error) {
	return p.weddingInfo, nil
}

func (p *Postgress) GetGuestList(context.Context) ([]entities.Guest, error) {
	return nil, nil
}

func (p *Postgress) GetGuestByID(context.Context, uuid.UUID) (entities.Guest, error) {
	return entities.Guest{}, nil
}

func (p *Postgress) GetGuestByEmail(context.Context, string) (entities.Guest, error) {
	return entities.Guest{}, nil
}

func (p *Postgress) RSVP(context.Context, entities.Guest) error {
	return nil
}

func parseRowsToItems(rows *sql.Rows) ([]entities.RegistryItem, error) {
	fmt.Println("parsing rows")
	var items []entities.RegistryItem
	for rows.Next() {
		var id uuid.UUID
		var name, description, link string
		var purchased bool
		if err := rows.Scan(&id, &name, &description, &link, &purchased); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		fmt.Println("addpening item")
		items = append(items, entities.RegistryItem{ID: id, Name: name, Description: description, Link: link, Purchased: purchased})
	}
	return items, nil
}
