package main

import (
	"fmt"
	"log"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/infastructure/repositories"
	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/Delta-a-Sierra/wedding_website/pkg/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("gotdotenv.Load: ", err.Error())
	}
	configHandler, err := config.NewHandler()
	if err != nil {
		panic(fmt.Errorf("config.NewHandler: %w", err))
	}
	conf, err := configHandler.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("configHandler.LoadConfig: %w", err))
	}
	repo := repositories.NewPostgresRepo()
	if err := repo.Connect(); err != nil {
		log.Fatalf("repo.Connect: %s", err.Error())
	}
	app := app.NewApp(repo)
	if conf.PresentationMethod == config.HTMXPM {
		htmx.Start(app)
	}
}
