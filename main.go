package main

import (
	"log"

	shorturl "github.com/CarlosEduardoAD/filler/internal/domain/models/shortUrl"
	"github.com/CarlosEduardoAD/filler/internal/utils"
	"github.com/CarlosEduardoAD/filler/pkg/cmd"

	"github.com/joho/godotenv"
)

func main() {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&shorturl.ShortUrl{})

	err = godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file => ", err)
	}

	cmd.Execute()
}
