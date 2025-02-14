package models

import "time"

type RequestData struct {
	ShortID   string    `json:"shortID"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"IP"`
}
