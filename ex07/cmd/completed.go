package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Show all completed tasks for the current day",
	Run: func(cmd *cobra.Command, args []string) {
		db := openDb()
		stmt := `SELECT * FROM finished
        WHERE timestamp >= date('now', '-1 day')`
		rows, err := db.Query(stmt)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		index := 1
		for rows.Next() {
			var finished struct {
				todo      string
				timestamp string
			}

			err = rows.Scan(&finished.todo, &finished.timestamp)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%d. %s\n", index, finished.todo)
			index++
		}

		if index == 1 {
			fmt.Println("No completed tasks for today")
		}
	},
}
