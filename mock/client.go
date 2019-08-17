package mock

import (
	"fmt"
	"time"

	"github.com/alexdavid/sigma"
)

type mockClient struct {
	lastID      int
	chats       map[int][]sigma.Message
	attachments map[int][]string
}

// NewClient creates a new mock client for testing or building sigma frontends on non-MacOS systems
func NewClient() (sigma.Client, error) {
	c := &mockClient{
		lastID:      0,
		chats:       map[int][]sigma.Message{},
		attachments: map[int][]string{},
	}

	for chatID, mockThread := range getMockChats() {
		for _, message := range mockThread {
			messageID := c.appendMessageToChat(chatID+1, sigma.Message{
				FromMe: message.FromMe,
				Text:   message.Text,
			})
			c.attachments[messageID] = message.Attachments
		}
	}

	return c, nil
}

func (c *mockClient) appendMessageToChat(chatID int, template sigma.Message) (messageID int) {
	messages, ok := c.chats[chatID]
	if !ok {
		messages = []sigma.Message{}
	}
	c.lastID++
	messageID = c.lastID
	c.chats[chatID] = append(messages, sigma.Message{
		ID:        messageID,
		FromMe:    template.FromMe,
		Text:      template.Text,
		Time:      time.Now(),
		Delivered: true,
	})
	return
}

func (c *mockClient) Attachments(messageID int) ([]string, error) {
	attachments := c.attachments[messageID]
	return attachments, nil
}

func (c *mockClient) Chats() ([]sigma.Chat, error) {
	results := []sigma.Chat{}
	for chatID := range c.chats {
		var lastActivity time.Time
		messages := c.chats[chatID]
		if len(messages) > 0 {
			lastActivity = messages[len(messages)-1].Time
		}
		results = append(results, sigma.Chat{
			ID:           chatID,
			DisplayName:  fmt.Sprintf("Chat %d", chatID),
			LastActivity: lastActivity,
		})
	}
	return results, nil
}

func (c *mockClient) Close() {}

func (c *mockClient) Messages(chatID int, filter sigma.MessageFilter) ([]sigma.Message, error) {
	messages, ok := c.chats[chatID]
	if !ok {
		return []sigma.Message{}, fmt.Errorf("Chat id %d doesn't exist", chatID)
	}
	return messages, nil
}

func (c *mockClient) SendMessage(chatID int, message string) error {
	_, ok := c.chats[chatID]
	if !ok {
		return fmt.Errorf("Chat id %d doesn't exist", chatID)
	}
	// Emulate a delay since applescript is slow to send
	time.Sleep(800 * time.Millisecond)
	c.appendMessageToChat(chatID, sigma.Message{Text: message, FromMe: true})
	return nil
}
