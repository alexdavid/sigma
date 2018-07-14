package api

import (
	"time"
)

type Message struct {
	Delivered bool
	FromMe    bool
	Text      string
	Time      time.Time
}

type Chat struct {
	Id          int
	DisplayName string
}
