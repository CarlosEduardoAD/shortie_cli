package cryptography

import "github.com/teris-io/shortid"

func GenerateShortCode() (string, error) {
	shortCode, err := shortid.Generate()

	if err != nil {
		return "", err
	}

	return shortCode, nil
}
