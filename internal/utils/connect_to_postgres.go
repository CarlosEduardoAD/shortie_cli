package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() (*gorm.DB, error) {
	err := LoadEnvWithPath()

	if err != nil {
		return nil, err
	}

	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("host=ep-broad-bush-a5eyo86v.us-east-2.aws.neon.tech user=CarlosEduardoAD password=%s dbname=shorties port=5432 sslmode=require TimeZone=America/Sao_Paulo", password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
