package main

import (
	"log"

	"github.com/CarlosEduardoAD/shortie_cli/pkg/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file => ", err)
	}

	cmd.Execute()
}
