package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"modernc.org/sqlite"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		db := openDb()
		todo := strings.Join(args, " ")

		insertStmt := `INSERT INTO todos (todo) 
        VALUES (?)`
		var sqliteErr *sqlite.Error
		_, err := db.Exec(insertStmt, todo)
		if errors.As(err, &sqliteErr) {
			// Todos are unique.
			// 2067 is the error code for unique constraint violation.
			if sqliteErr.Code() == 2067 {
				fmt.Println("Task already exists")
				return
			}

			fmt.Println(sqliteErr)
			return
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Added '%s'\n", todo)
	},
}
