package postgres

import (
	shorturl "github.com/CarlosEduardoAD/filler/internal/domain/models/shortUrl"
	"gorm.io/gorm"
)

type ShortUrlPostgres struct {
	db *gorm.DB
}

func NewShortUrlPostgres(db *gorm.DB) *ShortUrlPostgres {
	return &ShortUrlPostgres{
		db: db,
	}
}

func (s *ShortUrlPostgres) Create(shortUrl *shorturl.ShortUrl) error {
	result := s.db.Create(shortUrl)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ShortUrlPostgres) Find(code string) (*shorturl.ShortUrl, error) {
	var shortUrl *shorturl.ShortUrl

	result := s.db.Where("url_code = ?", code).First(&shortUrl)

	if result.Error != nil {
		return nil, result.Error
	}

	return shortUrl, nil
}
