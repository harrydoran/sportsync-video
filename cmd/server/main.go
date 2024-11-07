package main

import (
	"log"
	"net/http"

	"github.com/harrydoran/sportsync-video/internal/server"
)

func main() {
	log.Println("Starting SportSync Video server...")

	srv := server.New()
	if err := srv.Start(":8080"); err != http.ErrServerClosed {
		log.Fatalf("Server error: %v\n", err)
	}
}
