package api

func GetChats() ([]Chat, error) {
	rows, err := runSQL(`
		SELECT chat.ROWID, display_name, handle.id
		FROM chat
		LEFT JOIN chat_handle_join ON chat.ROWID = chat_handle_join.chat_id
		LEFT JOIN handle ON chat_handle_join.handle_id = handle.ROWID
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
		err = rows.Scan(&id, &displayName, &handleId)
		if err != nil {
			return []Chat{}, err
		}
		if displayName == "" {
			displayName = handleId
		}
		chats = append(chats, Chat{
			Id:          id,
			DisplayName: displayName,
		})
	}
	return chats, nil
}
