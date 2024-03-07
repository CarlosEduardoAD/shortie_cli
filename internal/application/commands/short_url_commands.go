package commands

import (
	"log"
	"strings"

	"github.com/CarlosEduardoAD/shortie_cli/internal/domain/services"
	postgres "github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/persistence/postgres"
	redis "github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/persistence/redis"
	"github.com/CarlosEduardoAD/shortie_cli/internal/utils"
)

func CreateShortUrl(url string) (string, error) {
	dbPsql, err := utils.ConnectToPostgres()

	if err != nil {
		return "", err
	}

	shortUrlRepositoryPostgres := postgres.NewShortUrlPostgres(dbPsql)
	shortUrlService := services.NewCreateShortUrl(shortUrlRepositoryPostgres)
	newUrl, err := shortUrlService.Execute(url)

	if err != nil {
		return "", err
	}

	return newUrl.ShortUrl, nil
}

func FindOriginalUrl(short_url string) (string, error) {
	db, err := utils.ConnectToPostgres()

	if err != nil {
		return "", err
	}

	dbRedis := utils.ConnectToRedis()

	shortUrlRepositoryRedis := redis.NewShortUrlRedis(dbRedis)
	findOriginalUrlCache := services.NewFindOriginalUrlCache(*shortUrlRepositoryRedis)

	code := strings.Split(short_url, "/")[3]

	cachedUrl, err := findOriginalUrlCache.FindOriginalUrlCache(code)

	if err == nil && cachedUrl != "" {
		return cachedUrl, nil
	}

	shortUrlRepository := postgres.NewShortUrlPostgres(db)
	shortUrlService := services.NewFindOriginalUrl(shortUrlRepository)
	shorturl, err := shortUrlService.Execute(code)

	if err != nil {
		log.Println("Deu erro 1")
		return "", err
	}

	createSHortUrlCache := services.NewCreateShortUrlCache(*shortUrlRepositoryRedis)
	err = createSHortUrlCache.CreateShortUrlCache(code, shorturl.OriginalUrl)

	if err != nil {
		log.Println("Deu erro 2")
		return "", err
	}

	return shorturl.OriginalUrl, nil
}
