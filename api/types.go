package api

import (
	"time"
)

type Message struct {
	Delivered bool      `json:"delivered"`
	FromMe    bool      `json:"fromMe"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}

type Chat struct {
	Id          int    `json:"id"`
	DisplayName string `json:"displayName"`
}
