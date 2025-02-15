// services/url_service.go
package services

import (
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
	"github.com/jaevor/go-nanoid"
)

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
func GetURL(url *models.URL) error {
	fmt.Println("Fetching URL for ID:", url.ID) // Debugging statement

	err := persistence.RedisGet(url)
	if err != nil {
		fmt.Println("Redis lookup failed:", err) // Debugging statement
		return err
	}

	fmt.Println("URL fetched from Redis:", url.URL) // Debugging statement
	return nil
}
