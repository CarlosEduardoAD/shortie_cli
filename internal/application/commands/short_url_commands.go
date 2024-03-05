package commands

import (
	"strings"

	"github.com/CarlosEduardoAD/filler/internal/domain/services"
	"github.com/CarlosEduardoAD/filler/internal/infrastructure/persistence/postgres"
	"github.com/CarlosEduardoAD/filler/internal/utils"
)

func CreateShortUrl(url string) (string, error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		return "", err
	}

	shortUrlRepository := postgres.NewShortUrlPostgres(db)
	shortUrlService := services.NewCreateShortUrl(shortUrlRepository)
	newUrl, err := shortUrlService.Execute(url)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return newUrl.ShortUrl, nil
}

func FindOriginalUrl(short_url string) (string, error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		return "", err
	}

	code := strings.Split(short_url, "/")[3]

	shortUrlRepository := postgres.NewShortUrlPostgres(db)
	shortUrlService := services.NewFindOriginalUrl(shortUrlRepository)
	shorturl, err := shortUrlService.Execute(code)

	if err != nil {
		return "", err
	}

	return shorturl, nil
}
