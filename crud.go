package main

import (
	"P9/crud_gin/router"
	"log"
)

func main() {
	r := router.Router()

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
