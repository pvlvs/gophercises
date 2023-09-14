package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		db := openDb()
		todo := strings.Join(args, " ")
		stmt := `DELETE FROM todos 
        WHERE todo = ?;`
		_, err := db.Exec(stmt, todo)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Marked '%s' as done\n", todo)
	},
}
