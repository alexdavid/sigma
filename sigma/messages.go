package sigma

import "database/sql"

type MessagesQuery struct {
	ChatId   int
	BeforeId int
	Limit    int
}

func Messages(query MessagesQuery) ([]Message, error) {
	rows, err := normalizeMessagesQuery(query)
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

const queryStart = `
  SELECT message.ROWID, message.date, message.text, message.is_sent, message.is_from_me
  FROM message
  LEFT JOIN chat_message_join ON message.ROWID = chat_message_join.message_id
  WHERE chat_message_join.chat_id = ?
`
const queryHasBeforeId = `
  AND message.ROWID < ?
`
const queryEnd = `
  ORDER BY date DESC
  LIMIT ?
`

func normalizeMessagesQuery(query MessagesQuery) (*sql.Rows, error) {
	if query.Limit == 0 {
		query.Limit = 20
	}
	if query.BeforeId != 0 {
		return runSQL(
			queryStart+queryHasBeforeId+queryEnd,
			query.ChatId,
			query.BeforeId,
			query.Limit,
		)
	}
	return runSQL(
		queryStart+queryEnd,
		query.ChatId,
		query.Limit,
	)
}
