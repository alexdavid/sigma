package sigma

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

type realClient struct {
	db *sql.DB
}

func NewClient() (Client, error) {
	db, err := sql.Open(
		"sqlite3",
		path.Join(os.Getenv("HOME"), "Library/Messages/chat.db?mode=ro"),
	)
	if err != nil {
		return nil, err
	}
	return &realClient{
		db: db,
	}, nil
}

func (c *realClient) Close() {
	c.db.Close()
}

func (c *realClient) runSQL(query string, args ...interface{}) (*sql.Rows, error) {
	if len(args) == 0 {
		rows, err := c.db.Query(query)
		return rows, err
	}

	stmt, err := c.db.Prepare(query)
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
