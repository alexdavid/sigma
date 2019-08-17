package sigma

import (
	"fmt"
	"os/exec"
)

func (c *realClient) SendMessage(chatID int, message string) error {
	handleID, serviceID, err := c.getHandleAndServiceID(chatID)
	if err != nil {
		return err
	}
	cmd := exec.Command("osascript", "-e", applescript, message, handleID, serviceID)
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
func (c *realClient) getHandleAndServiceID(chatID int) (handleID string, serviceID string, err error) {
	rows, err := c.runSQL(`
		SELECT handle.id, chat.account_id
		FROM chat_handle_join
		LEFT JOIN handle ON chat_handle_join.handle_id = handle.ROWID
    LEFT JOIN chat ON chat.ROWID = chat_id
		WHERE chat_id = ?
	`, chatID)
	if err != nil {
		return
	}
	defer rows.Close()

	found := rows.Next()
	if !found {
		err = fmt.Errorf("Could not find handle of chat %d", chatID)
		return
	}

	err = rows.Scan(&handleID, &serviceID)
	return
}
