package sigma

import (
	"io"
	"time"
)

// Client interface represents a sigma client (either real or mock)
type Client interface {
	// Chats returns a list of chat threads
	Chats() ([]Chat, error)

	// Messages returns a list of messages for a given chat id
	Messages(chatID int, filter MessageFilter) ([]Message, error)

	// SendMessage sends the given string message to the specified chat id
	SendMessage(chatID int, message string) error

	// SendMedia sends a new empty message with arbitrary data as an attachment to
	// a given chat id
	//
	// `name` is used for the file name that will used on disk which is visible if
	// the sender or receiver choose to save the image
	//
	// Due to current limitations the name will be prefixed with a random id
	SendMedia(chatID int, name string, data io.Reader) error

	// Attachments returns an array of file paths of attachments on the
	// specified message id
	Attachments(messageID int) ([]string, error)

	// Close destructs the client
	Close()
}

// Chat represents an active conversation with another user
type Chat struct {
	ID           int       `json:"id"`
	DisplayName  string    `json:"displayName"`
	LastActivity time.Time `json:"lastActivity"`
}

// Message represents a single message in a chat
type Message struct {
	ID        int       `json:"id"`
	Delivered bool      `json:"delivered"`
	FromMe    bool      `json:"fromMe"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}

// MessageFilter is used to help paginate message results
type MessageFilter struct {
	AfterTime time.Time //(optional) get messages that occured after specified time.Time)
	BeforeID  int       // (optional) get messages before the specified message id
	Limit     int       // the maximum number of messages to return
}
