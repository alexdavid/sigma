package sigma

import (
	"fmt"
	"os/exec"
)

func SendMessage(chatId int, message string) error {
	handleId, err := getHandleIdFromChatId(chatId)
	if err != nil {
		return err
	}
	cmd := exec.Command("osascript", "-e", applescript, message, handleId)
	return cmd.Run()
}

// Modified from https://github.com/bboyairwreck/PieMessage
const applescript = `
on run {msgText, handle}
	tell application "Messages"
		set serviceID to id of 1st service whose service type = iMessage
		send msgText to buddy handle of service id serviceID
	end tell
end run
`

// Applescript needs the message handle id so look it up:
func getHandleIdFromChatId(chatId int) (string, error) {
	rows, err := runSQL(`
		SELECT handle.id
		FROM chat_handle_join
		LEFT JOIN handle ON chat_handle_join.handle_id = handle.ROWID
		WHERE chat_id = ?
	`, chatId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	found := rows.Next()
	if !found {
		return "", fmt.Errorf("Could not find handle of chat %d", chatId)
	}

	var handeId string
	err = rows.Scan(&handeId)
	return handeId, err
}
