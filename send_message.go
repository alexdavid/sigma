package sigma

import (
	"fmt"
	"os/exec"
)

func (c *realClient) SendMessage(chatId int, message string) error {
	handleId, serviceId, err := c.getHandleAndServiceId(chatId)
	if err != nil {
		return err
	}
	cmd := exec.Command("osascript", "-e", applescript, message, handleId, serviceId)
	return cmd.Run()
}

// Modified from https://github.com/bboyairwreck/PieMessage
const applescript = `
on run {msgText, handleId, serviceId}
	tell application "Messages"
		send msgText to buddy handleId of service id serviceId
	end tell
end run
`

// Applescript needs the message handle & service id so look it up:
func (c *realClient) getHandleAndServiceId(chatId int) (handleId string, serviceId string, err error) {
	rows, err := c.runSQL(`
		SELECT handle.id, chat.account_id
		FROM chat_handle_join
		LEFT JOIN handle ON chat_handle_join.handle_id = handle.ROWID
    LEFT JOIN chat ON chat.ROWID = chat_id
		WHERE chat_id = ?
	`, chatId)
	if err != nil {
		return
	}
	defer rows.Close()

	found := rows.Next()
	if !found {
		err = fmt.Errorf("Could not find handle of chat %d", chatId)
		return
	}

	err = rows.Scan(&handleId, &serviceId)
	return
}
