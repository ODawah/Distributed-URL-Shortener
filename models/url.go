package models

import (
	"github.com/jaevor/go-nanoid"
)

type URL struct {
	URL string `json:"url" binding:"required"`
	ID  string `json:"id"`
}

func (Url *URL) Shorten() error {
	canonicID, err := nanoid.Standard(8)
	if err != nil {
		return err
	}
	id := canonicID()
	Url.ID = id
	return nil
}
