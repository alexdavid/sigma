package sigma

import (
	"time"
)

func (c *realClient) Chats() ([]Chat, error) {
	rows, err := c.runSQL(`
		SELECT chat.ROWID, display_name, handle.id, COALESCE(MAX(message.date),0) as last_activity
		FROM chat
		LEFT JOIN chat_handle_join ON chat.ROWID = chat_handle_join.chat_id
		LEFT JOIN handle ON chat_handle_join.handle_id = handle.ROWID
		LEFT JOIN chat_message_join ON chat_message_join.chat_id = chat.ROWID
		LEFT JOIN message ON chat_message_join.message_id = message.ROWID
		GROUP BY chat.ROWID
		ORDER BY last_activity DESC
	`)
	if err != nil {
		return []Chat{}, err
	}
	defer rows.Close()

	chats := []Chat{}

	for rows.Next() {
		var id int
		var displayName string
		var handleId string
		var lastActivity int
		err = rows.Scan(&id, &displayName, &handleId, &lastActivity)
		if err != nil {
			return []Chat{}, err
		}
		if displayName == "" {
			displayName = handleId
		}
		var lastActivityTime time.Time
		if lastActivity > 0 {
			lastActivityTime = cocoaTimestampToTime(lastActivity)
		}
		chats = append(chats, Chat{
			Id:           id,
			DisplayName:  displayName,
			LastActivity: lastActivityTime,
		})
	}
	return chats, nil
}
