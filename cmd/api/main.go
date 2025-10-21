package main

import (
	"log"

	"github.com/bccfilkom-be/postify/internal/bootstrap"
)

func main() {
	err := bootstrap.StartApp()

	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
