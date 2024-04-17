package repositories

import (
	"database/sql"
	"fmt"
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
	connStr := "postgresql://root:password@localhost:5432/registry?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("db.Ping: %w", err)
	}
	p.db = db

	return nil
}

// TODO: finish implemeting postgress
func (p *Postgress) seed() error {
	_, err := p.db.Exec("DROP TABLE IF EXISTS items;")
	if err != nil {
		return fmt.Errorf("p.db.Exec: %w", err)
	}

	_, err = p.db.Exec("CREATE TABLE items(ID INT PRIMARY KEY NOT NULL,name VARCHAR(100) NOT NULL,description VARCHAR(255) NOT NULL);link VARCHAR(255) NOT NULL")
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
	return nil
}
