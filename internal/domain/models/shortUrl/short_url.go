package shorturl

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ShortUrl struct {
	Id               string    `json:"id"`
	OriginalUrl      string    `json:"original_url"`
	ShortUrl         string    `json:"short_url"`
	UrlCode          string    `json:"url_code"`
	CreatedAt        time.Time `json:"created_at"`
	LastTimeAccessed time.Time `json:"last_time_accessed"`
}

func NewShortUrl(originalUrl, shortUrl, urlCode string, lastTimeAccessed time.Time) *ShortUrl {
	return &ShortUrl{
		Id:               uuid.NewString(),
		OriginalUrl:      originalUrl,
		ShortUrl:         shortUrl,
		UrlCode:          urlCode,
		CreatedAt:        time.Now(),
		LastTimeAccessed: lastTimeAccessed,
	}
}

func (s *ShortUrl) Validate() error {
	if s.OriginalUrl == "" {
		return errors.New("OriginalUrl is required")
	}

	if s.ShortUrl == "" {
		return errors.New("ShortUrl is required")
	}

	if s.UrlCode == "" {
		return errors.New("urlCode is required")
	}

	if s.LastTimeAccessed.IsZero() {
		return errors.New("LastTimeAccessed is required")
	}

	return nil

}
