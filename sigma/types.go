package sigma

import (
	"time"
)

type Message struct {
	Id        int       `json:"id"`
	Delivered bool      `json:"delivered"`
	FromMe    bool      `json:"fromMe"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}

type Chat struct {
	Id           int       `json:"id"`
	DisplayName  string    `json:"displayName"`
	LastActivity time.Time `json:"lastActivity"`
}
