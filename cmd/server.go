package main

import (
	"gocipher/api"
	"log"
)

func main() {

	r := api.SetUpRouter()

	log.Println("Server running at port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
