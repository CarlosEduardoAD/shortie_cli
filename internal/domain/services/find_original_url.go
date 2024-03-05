package services

import (
	shortUrlRepository "github.com/CarlosEduardoAD/filler/internal/infrastructure/persistence/postgres"
)

type FindOriginalUrl struct {
	ShortUrlRepository *shortUrlRepository.ShortUrlPostgres
}

func NewFindOriginalUrl(shortUrlRepository *shortUrlRepository.ShortUrlPostgres) *FindOriginalUrl {
	return &FindOriginalUrl{
		ShortUrlRepository: shortUrlRepository,
	}
}

func (c *FindOriginalUrl) Execute(code string) (string, error) {
	shortUrl, err := c.ShortUrlRepository.Find(code)

	if err != nil {
		return "", err
	}

	return shortUrl.OriginalUrl, nil
}
