package mock

import (
	"fmt"
	"time"

	"github.com/alexdavid/sigma"
)

type mockClient struct {
	lastId      int
	chats       map[int][]sigma.Message
	attachments map[int][]string
}

func NewClient() (sigma.Client, error) {
	c := &mockClient{
		lastId:      0,
		chats:       map[int][]sigma.Message{},
		attachments: map[int][]string{},
	}

	for chatId, mockThread := range getMockChats() {
		for _, message := range mockThread {
			messageId := c.appendMessageToChat(chatId+1, sigma.Message{
				FromMe: message.FromMe,
				Text:   message.Text,
			})
			c.attachments[messageId] = message.Attachments
		}
	}

	return c, nil
}

func (c *mockClient) appendMessageToChat(chatId int, template sigma.Message) (messageId int) {
	messages, ok := c.chats[chatId]
	if !ok {
		messages = []sigma.Message{}
	}
	c.lastId++
	messageId = c.lastId
	c.chats[chatId] = append(messages, sigma.Message{
		Id:        messageId,
		FromMe:    template.FromMe,
		Text:      template.Text,
		Time:      time.Now(),
		Delivered: true,
	})
	return
}

func (c *mockClient) Attachments(messageId int) ([]string, error) {
	attachments := c.attachments[messageId]
	return attachments, nil
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

func (c *mockClient) Close() {}

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
