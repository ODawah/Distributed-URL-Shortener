package models

type URL struct {
	URL string `json:"url" binding:"required"`
	ID  string `json:"id"`
}
