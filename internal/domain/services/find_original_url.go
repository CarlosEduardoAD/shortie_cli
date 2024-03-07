package services

import (
	shorturl "github.com/CarlosEduardoAD/shortie_cli/internal/domain/models/shortUrl"
	shortUrlRepository "github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/persistence/postgres"
)

type FindOriginalUrl struct {
	ShortUrlRepository *shortUrlRepository.ShortUrlPostgres
}

func NewFindOriginalUrl(shortUrlRepository *shortUrlRepository.ShortUrlPostgres) *FindOriginalUrl {
	return &FindOriginalUrl{
		ShortUrlRepository: shortUrlRepository,
	}
}

func (c *FindOriginalUrl) Execute(code string) (*shorturl.ShortUrl, error) {
	shortUrl, err := c.ShortUrlRepository.Find(code)

	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}
