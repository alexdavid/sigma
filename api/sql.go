package api

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func runSQL(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := sql.Open(
		"sqlite3",
		path.Join(os.Getenv("HOME"), "Library/Messages/chat.db?mode=ro"),
	)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if len(args) == 0 {
		rows, err := db.Query(query)
		return rows, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
