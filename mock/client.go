package mock

import "github.com/alexdavid/sigma"

type mockClient struct {
}

func NewClient() (sigma.Client, error) {
	return &mockClient{}, nil
}

func (c *mockClient) Attachments(id int) ([]string, error) {
	return []string{}, nil
}

func (c *mockClient) Chats() ([]sigma.Chat, error) {
	return []sigma.Chat{}, nil
}

func (c *mockClient) Close() {
}

func (c *mockClient) Messages(query sigma.MessagesQuery) ([]sigma.Message, error) {
	return []sigma.Message{}, nil
}

func (c *mockClient) SendMessage(chatId int, message string) error {
	return nil
}
