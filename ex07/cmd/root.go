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
	openDb()

	stmt := `CREATE TABLE IF NOT EXISTS todos (
        todo TEXT UNIQUE NOT NULL
    );`

	db := openDb()
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
