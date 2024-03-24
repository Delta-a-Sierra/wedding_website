package main

import (
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/infastructure/repositories"
	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/Delta-a-Sierra/wedding_website/pkg/config"
)

func main() {
	configHandler, err := config.NewHandler()
	if err != nil {
		panic(fmt.Errorf("config.NewHandler: %w", err))
	}
	conf, err := configHandler.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("configHandler.LoadConfig: %w", err))
	}
	repo := repositories.NewInMemoryRepo()
	app := app.NewApp(repo)
	if conf.PresentationMethod == config.HTMXPM {
		htmx.Start(app)
	}
}
