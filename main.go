package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("[MAIN] :: Starting server on port 8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
}
