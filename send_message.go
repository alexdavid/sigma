package sigma

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

const sendMessageApplescript = `
on run {msgText, handleId, serviceId}
	tell application "Messages"
		send msgText to buddy handleId of service id serviceId
	end tell
end run
`

const sendMediaApplescript = `
on run {filePath, handleId, serviceId}
  set media to POSIX file filePath
  set baseName to do shell script "basename " & filePath
	tell application "Messages"
		send media to buddy handleId of service id serviceId
    repeat
      set fileTransfers to file transfers
      repeat with fileTransfer in fileTransfers
        if name of fileTransfer = baseName then
          return transfer status of fileTransfer
        end if
      end repeat
      delay 1
    end repeat
	end tell
end run
`

func (c *realClient) SendMessage(chatID int, message string) error {
	handleID, serviceID, err := c.getHandleAndServiceID(chatID)
	if err != nil {
		return err
	}
	cmd := exec.Command("osascript", "-e", sendMessageApplescript, message, handleID, serviceID)
	return cmd.Run()
}

func (c *realClient) SendMedia(chatID int, name string, data io.Reader) error {
	handleID, serviceID, err := c.getHandleAndServiceID(chatID)
	if err != nil {
		return err
	}

	tmpFile, err := ioutil.TempFile("", "*."+name)
	if err != nil {
		return err
	}
	filePath := tmpFile.Name()
	defer os.Remove(filePath)

	if _, err = io.Copy(tmpFile, data); err != nil {
		return err
	}

	cmd := exec.Command("osascript", "-e", sendMediaApplescript, filePath, handleID, serviceID)
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	if string(stdout) != "finished" {
		return fmt.Errorf("Expected transfer status `finished` but got %s", string(stdout))
	}
	return nil
}

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
