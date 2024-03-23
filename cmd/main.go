package main

import (
	"fmt"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx"
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
	if conf.PresentationMethod == config.HTMXPM {
		htmx.Start()
	}
}
