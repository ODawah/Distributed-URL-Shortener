package models

type URL struct {
	ID  string `json:"id" binding:"required"`
	URL string `json:"url"`
}
