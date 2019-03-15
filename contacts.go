package sigma

import (
	"database/sql"
	"os"
	"path"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

func getContactMap() (map[string]string, error) {
	contacts := map[string]string{}

	db, err := sql.Open(
		"sqlite3",
		path.Join(os.Getenv("HOME"), "Library/Application Support/AddressBook/AddressBook-v22.abcddb"),
	)
	if err != nil {
		return contacts, err
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT ZFIRSTNAME, ZLASTNAME, ZORGANIZATION, ZFULLNUMBER
    FROM ZABCDPHONENUMBER
    LEFT JOIN ZABCDRECORD ON ZOWNER=ZABCDRECORD.Z_PK
  `)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()

	for rows.Next() {
		var firstName string
		var lastName string
		var company string
		var number string

		err = rows.Scan(&firstName, &lastName, &company, &number)
		if err != nil {
			return contacts, err
		}
		contacts[normalizeNumber(number)] = formatName(firstName, lastName, company, number)
	}

	return contacts, nil
}

func normalizeNumber(number string) string {
	r := regexp.MustCompile(`[^0-9A-Z]`)
	return r.ReplaceAllString(number, "")
}

func formatName(firstName, lastName, company, number string) string {
	if firstName != "" && lastName != "" {
		return firstName + " " + lastName
	}
	if firstName != "" {
		return firstName
	}
	if lastName != "" {
		return lastName
	}
	if company != "" {
		return company
	}
	return number
}
