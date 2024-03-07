package services

import (
	shortUrlRepository "github.com/CarlosEduardoAD/shortie_cli/internal/infrastructure/persistence/redis"
)

type CreateShortUrlCache struct {
	shortUrlRepository shortUrlRepository.ShortUrlRedis
}

func NewCreateShortUrlCache(shortUrlRepository shortUrlRepository.ShortUrlRedis) *CreateShortUrlCache {
	return &CreateShortUrlCache{
		shortUrlRepository: shortUrlRepository,
	}
}

func (s *CreateShortUrlCache) CreateShortUrlCache(code, originalUrl string) error {
	err := s.shortUrlRepository.CreateCache(code, originalUrl)

	if err != nil {
		return err
	}

	return nil
}
