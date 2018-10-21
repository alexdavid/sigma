package mock

import (
	"fmt"
	"time"

	"github.com/alexdavid/sigma"
)

type mockClient struct {
	lastId int
	chats  map[int][]sigma.Message
}

func NewClient() (sigma.Client, error) {
	c := &mockClient{
		lastId: 0,
		chats:  map[int][]sigma.Message{},
	}

	for chatId, mockThread := range mockChats {
		for _, message := range mockThread {
			c.appendMessageToChat(chatId+1, message)
		}
	}

	return c, nil
}

func (c *mockClient) appendMessageToChat(chatId int, template sigma.Message) {
	messages, ok := c.chats[chatId]
	if !ok {
		messages = []sigma.Message{}
	}
	c.lastId++
	c.chats[chatId] = append(messages, sigma.Message{
		Id:        c.lastId,
		FromMe:    template.FromMe,
		Text:      template.Text,
		Time:      time.Now(),
		Delivered: true,
	})
}

func (c *mockClient) Attachments(messageId int) ([]string, error) {
	return []string{}, nil
}

func (c *mockClient) Chats() ([]sigma.Chat, error) {
	results := []sigma.Chat{}
	for chatId := range c.chats {
		var lastActivity time.Time
		messages := c.chats[chatId]
		if len(messages) > 0 {
			lastActivity = messages[len(messages)-1].Time
		}
		results = append(results, sigma.Chat{
			Id:           chatId,
			DisplayName:  fmt.Sprintf("Chat %d", chatId),
			LastActivity: lastActivity,
		})
	}
	return results, nil
}

func (c *mockClient) Close() {
}

func (c *mockClient) Messages(query sigma.MessagesQuery) ([]sigma.Message, error) {
	messages, ok := c.chats[query.ChatId]
	if !ok {
		return []sigma.Message{}, fmt.Errorf("Chat id %d doesn't exist", query.ChatId)
	}
	return messages, nil
}

func (c *mockClient) SendMessage(chatId int, message string) error {
	_, ok := c.chats[chatId]
	if !ok {
		return fmt.Errorf("Chat id %d doesn't exist", chatId)
	}
	// Emulate a delay since applescript is slow to send
	time.Sleep(800 * time.Millisecond)
	c.appendMessageToChat(chatId, sigma.Message{Text: message, FromMe: true})
	return nil
}
