package main

import (
	"context"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/cmd/config"
	"github.com/gbrlmza/lana-bechallenge-checkout/cmd/container"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/rest"
	"log"
	"net/http"
)

func main() {
	// TODO: Explain hexagonal architecture
	// TODO: Explain app wiring
	ctx := context.Background()

	// Config
	cfg := config.Get()

	// Container & service initialization
	container := container.NewContainer(ctx)
	service := checkout.NewService(container)

	// Handler
	handler := rest.NewHandler(service)
	router := handler.RouterInit()

	// Start server
	fmt.Printf("### Environment: %s\n", cfg.Environment)
	fmt.Printf("### Starting server at port: %s\n", cfg.Port)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
