package cmd

import (
	"database/sql"
	"log"
	"os"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var (
	rootCmd = &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	db := openDb()
	createTodosTable(db)
    createFinishedTable(db)
}

func createTodosTable(db *sql.DB) {
	stmt := `CREATE TABLE IF NOT EXISTS todos (
        todo TEXT UNIQUE NOT NULL
    );`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func createFinishedTable(db *sql.DB) {
	stmt := `CREATE TABLE IF NOT EXISTS finished (
        todo TEXT NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
    );`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func openDb() *sql.DB {
	dsn := "file:local.db"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return db
}
