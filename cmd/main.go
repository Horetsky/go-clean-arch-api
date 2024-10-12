package main

import (
	"log"
	"seeker/internal/app"
)

func main() {
	err := app.Start()

	if err != nil {
		log.Fatalf("Failed to start application, %s", err)
	}
}
