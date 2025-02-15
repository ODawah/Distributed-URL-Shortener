// services/url_service.go
package services

import (
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
	"github.com/jaevor/go-nanoid"
	"regexp"
)

func IsValidURL(url string) bool {
	// Define regex pattern for validating websites
	pattern := `^www\.[a-zA-Z0-9-]+\.(com|net|org|edu|gov|io|co|uk|us|info|biz|tv|me)$`

	// Compile regex
	re := regexp.MustCompile(pattern)

	// Match the input website against the regex
	return re.MatchString(url)
}

// Shorten generates a short URL ID and saves it in Redis
func Shorten(url *models.URL) error {

	canonicID, err := nanoid.Standard(8)
	if err != nil {
		return err
	}
	url.ID = canonicID()

	err = persistence.RedisInsert(url)
	if err != nil {
		return err
	}
	return nil
}

// GetURL retrieves a URL from Redis
func GetURL(id string) *models.URL {

	url := persistence.RedisGet(id)
	if url == nil {
		return nil
	}

	return url
}
