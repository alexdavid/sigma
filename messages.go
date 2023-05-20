package sigma

import (
	"database/sql"
)

func (c *realClient) Messages(chatID int, query MessageFilter) ([]Message, error) {
	rows, err := c.normalizeMessagesQuery(chatID, query)
	if err != nil {
		return []Message{}, err
	}
	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var id int
		var timestamp int64
		var text sql.NullString
		var isSent bool
		var isFromMe bool
		err = rows.Scan(&id, &timestamp, &text, &isSent, &isFromMe)
		if err != nil {
			return []Message{}, err
		}
		messages = append(messages, Message{
			ID:        id,
			Time:      cocoaTimestampToTime(timestamp),
			Text:      text.String,
			Delivered: isSent,
			FromMe:    isFromMe,
		})
	}

	return messages, nil
}

const queryStart = `
  SELECT message.ROWID, message.date, message.text, message.is_sent, message.is_from_me
  FROM message
  LEFT JOIN chat_message_join ON message.ROWID = chat_message_join.message_id
  WHERE chat_message_join.chat_id = ?
`
const queryHasBeforeID = `
  AND message.ROWID < ?
`
const queryEnd = `
  ORDER BY date DESC
  LIMIT ?
`
const queryHasAfterID = `
  AND message.ROWID > ?
`

func (c *realClient) normalizeMessagesQuery(chatID int, query MessageFilter) (*sql.Rows, error) {
	if query.Limit == 0 {
		query.Limit = 20
	}
	if query.BeforeID != 0 {
		return c.runSQL(queryStart+queryHasBeforeID+queryEnd, chatID, query.BeforeID, query.Limit)
	}
	if query.AfterID != 0 {
		return c.runSQL(queryStart+queryHasAfterID+queryEnd, chatID, query.AfterID, query.Limit)
	}
	return c.runSQL(queryStart+queryEnd, chatID, query.Limit)
}
