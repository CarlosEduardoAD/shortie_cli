package commands

import (
	"strings"

	"github.com/CarlosEduardoAD/filler/internal/domain/services"
	postgres "github.com/CarlosEduardoAD/filler/internal/infrastructure/persistence/postgres"
	redis "github.com/CarlosEduardoAD/filler/internal/infrastructure/persistence/redis"
	"github.com/CarlosEduardoAD/filler/internal/utils"
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
		return "", err
	}

	createSHortUrlCache := services.NewCreateShortUrlCache(*shortUrlRepositoryRedis)
	err = createSHortUrlCache.CreateShortUrlCache(code, shorturl.OriginalUrl)

	if err != nil {
		return "", err
	}

	return shorturl.OriginalUrl, nil
}
