package main

import (
	"bytes"
	"database/sql"
	"log"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
	_ "modernc.org/sqlite"
)

type entry struct {
	id     string
	number string
}

func main() {
	db, err := openDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	entries, err := getPhoneNumbers(db)
	if err != nil {
		log.Fatal(err)
	}

	normalizedEntries := []string{}
	for _, v := range entries {
		normalized := normalize(v.number)
		if slices.Contains(normalizedEntries, normalized) {
			delete(db, v)
			continue
		}

		normalizedEntries = append(normalizedEntries, normalized)
		update(db, v.id, normalized)
	}
}

func openDb() (*sql.DB, error) {
	dsn := "file:local.db"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	dropStmt := "DROP TABLE IF EXISTS phone_numbers"
	_, err := db.Exec(dropStmt)
	if err != nil {
		return err
	}

	createStmt := `CREATE TABLE IF NOT EXISTS phone_numbers(
        id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
        number TEXT NOT NULL
    );`

	_, err = db.Exec(createStmt)
	if err != nil {
		return err
	}

	phoneNumbers := []interface{}{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	insertStmt := `INSERT INTO phone_numbers(number)
    VALUES(?)`
	for i := range phoneNumbers {
		if i == len(phoneNumbers)-1 {
			break
		}
		insertStmt = strings.Join([]string{insertStmt, "(?)"}, ",")
	}

	_, err = db.Exec(insertStmt, phoneNumbers...)
	if err != nil {
		return err
	}

	return nil
}

func getPhoneNumbers(db *sql.DB) ([]entry, error) {
	stmt := "SELECT * FROM phone_numbers"

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []entry
	for rows.Next() {
		e := &entry{}
		err = rows.Scan(&e.id, &e.number)
		if err != nil {
			return nil, err
		}

		entries = append(entries, *e)
	}

	return entries, nil
}

func normalize(number string) string {
	var buf bytes.Buffer
	for _, v := range number {
		if unicode.IsDigit(v) {
			buf.WriteRune(v)
		}
	}

	return buf.String()
}

func update(db *sql.DB, id string, n string) error {
	stmt := `UPDATE phone_numbers
    SET number = ?
    WHERE id = ?`

	_, err := db.Exec(stmt, n, id)
	if err != nil {
		return err
	}

	return nil
}

func delete(db *sql.DB, e entry) error {
	stmt := `DELETE FROM phone_numbers
    WHERE id = ?`

	_, err := db.Exec(stmt, e.id)
	if err != nil {
		return err
	}

	return nil
}
