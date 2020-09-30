package main

import (
	"github.com/gbrlmza/lana-bechallenge-checkout/pkg/rest"
	"log"
	"net/http"
)

func main() {
	// Init Router
	r := rest.Router()

	// Start Server
	log.Fatal(http.ListenAndServe(":80", r))
}
