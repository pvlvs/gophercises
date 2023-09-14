package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db := openDb()
		stmt := `SELECT * FROM todos;`
		rows, err := db.Query(stmt)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		index := 1
		for rows.Next() {
			var todo *string
			err = rows.Scan(&todo)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%d. %s\n", index, *todo)
			index++
		}

		if index == 1 {
			fmt.Println("You have no open tasks")
		}
	},
}
