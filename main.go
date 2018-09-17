package main

import (
	"github.com/joho/godotenv"
	"log"
	"advanced-computer-lab/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.Start()
}
