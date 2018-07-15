package api

import (
	"time"
)

func GetMessages(chatId int, since time.Time) ([]Message, error) {
	rows, err := runSQL(`
		SELECT message.ROWID, message.date, message.text, message.is_sent, message.is_from_me
		FROM message
		LEFT JOIN chat_message_join ON message.ROWID = chat_message_join.message_id
		WHERE chat_message_join.chat_id = ?
		AND date > ?
		ORDER BY date ASC
	`, chatId, timeToCocoaTimestamp(since))
	if err != nil {
		return []Message{}, err
	}
	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var id int
		var timestamp int
		var text string
		var isSent bool
		var isFromMe bool
		err = rows.Scan(&id, &timestamp, &text, &isSent, &isFromMe)
		if err != nil {
			return []Message{}, err
		}
		messages = append(messages, Message{
			Id:        id,
			Time:      cocoaTimestampToTime(timestamp),
			Text:      text,
			Delivered: isSent,
			FromMe:    isFromMe,
		})
	}

	return messages, nil
}
