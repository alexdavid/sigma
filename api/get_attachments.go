package api

import (
	"os"
	"path"
)

func GetAttachments(messageId int) ([]string, error) {
	rows, err := runSQL(`
		SELECT attachment.filename
		FROM attachment
		LEFT JOIN message_attachment_join ON attachment.ROWID = message_attachment_join.attachment_id
		WHERE message_attachment_join.message_id = ?
	`, messageId)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	fileNames := []string{}
	for rows.Next() {
		var filePath string
		err = rows.Scan(&filePath)
		if err != nil {
			return []string{}, err
		}

		// filepath is stored as `~/Library/...`. We normalize this since `~` is a bashism
		if filePath[0] == '~' {
			filePath = path.Join(os.Getenv("HOME"), filePath[1:])
		}
		if err != nil {
			return []string{}, err
		}
		fileNames = append(fileNames, filePath)
	}

	return fileNames, nil
}
