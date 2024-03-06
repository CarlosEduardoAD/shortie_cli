package shorturl

type ShortUrlActions interface {
	Create(shortUrl *ShortUrl) error
	Find(code string) (*ShortUrl, error)
	CreateCache(code, originalUrl string) error
	FindCache(orignialUrl string) (string, error)
}
