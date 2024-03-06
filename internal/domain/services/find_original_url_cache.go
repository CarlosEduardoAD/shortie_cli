package services

import shortUrlRepository "github.com/CarlosEduardoAD/filler/internal/infrastructure/persistence/redis"

type FindOriginalUrlCache struct {
	shortUrlRepository shortUrlRepository.ShortUrlRedis
}

func NewFindOriginalUrlCache(shortUrlRepository shortUrlRepository.ShortUrlRedis) *FindOriginalUrlCache {
	return &FindOriginalUrlCache{
		shortUrlRepository: shortUrlRepository,
	}
}

func (s *FindOriginalUrlCache) FindOriginalUrlCache(code string) (string, error) {
	originalUrl, err := s.shortUrlRepository.FindCache(code)

	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
