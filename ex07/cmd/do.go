package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		db := openDb()
		todo := strings.Join(args, " ")
		deleteStmt := `DELETE FROM todos 
        WHERE todo = ?;`
		_, err := db.Exec(deleteStmt, todo)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		insertStmt := `INSERT INTO finished (todo) 
        VALUES (?)`
		_, err = db.Exec(insertStmt, todo)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Marked '%s' as done\n", todo)
	},
}
