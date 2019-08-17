package sigma

import (
	"time"
)

// Client interface represents a sigma client (either real or mock)
type Client interface {
	Attachments(messageID int) ([]string, error)
	Chats() ([]Chat, error)
	Close()
	Messages(chatID int, filter MessageFilter) ([]Message, error)
	SendMessage(chatID int, message string) error
}

// Message is a single message in a chat
type Message struct {
	ID        int       `json:"id"`
	Delivered bool      `json:"delivered"`
	FromMe    bool      `json:"fromMe"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}

// Chat is an active conversation with another user
type Chat struct {
	ID           int       `json:"id"`
	DisplayName  string    `json:"displayName"`
	LastActivity time.Time `json:"lastActivity"`
}

// MessageFilter is used to help paginate message results
type MessageFilter struct {
	BeforeID int // (optional) get messages before the specified message id
	Limit    int // the maximum number of messages to return
}
