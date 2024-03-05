package shorturl

type ShortUrlPostgresActions interface {
	Create(shortUrl *ShortUrl) error
	Find(shortUrl *ShortUrl) (*ShortUrl, error)
}
