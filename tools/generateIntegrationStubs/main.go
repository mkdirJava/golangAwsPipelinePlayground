package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/api"
	"github.com/Yamashou/gqlgenc/clientgenv2"
	"github.com/Yamashou/gqlgenc/config"
	"github.com/Yamashou/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Fatal(err)
	}

	// Get Plugins
	clientGen := api.AddPlugin(clientgenv2.New(cfg.Query, cfg.Client, cfg.Generate))

	// Generate Client
	err = generator.Generate(ctx, cfg, clientGen)
	if err != nil {
		log.Fatal(err)
	}
}
