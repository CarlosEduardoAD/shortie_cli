package services

import (
	"fmt"
	"time"

	shorturl "github.com/CarlosEduardoAD/shortie_cli/internal/domain/models/shortUrl"
	"github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/cryptography"
	shortUrlRepository "github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/persistence/postgres"
)

type CreateShortUrl struct {
	ShortUrlRepository *shortUrlRepository.ShortUrlPostgres
}

func NewCreateShortUrl(shortUrlRepository *shortUrlRepository.ShortUrlPostgres) *CreateShortUrl {
	return &CreateShortUrl{
		ShortUrlRepository: shortUrlRepository,
	}
}

func (c *CreateShortUrl) Execute(url string) (*shorturl.ShortUrl, error) {
	shortCode, err := cryptography.GenerateShortCode()

	if err != nil {
		return nil, err
	}

	shortUrl := shorturl.NewShortUrl(url, fmt.Sprintf("http://localhost:9090/%s", shortCode), shortCode, time.Time{})

	err = c.ShortUrlRepository.Create(shortUrl)

	if err != nil {
		return nil, err
	}

	newUrl := shorturl.NewShortUrl(shortUrl.OriginalUrl, shortUrl.ShortUrl, shortUrl.UrlCode, time.Time{})

	return newUrl, nil
}
