package cmd

import (
	"fmt"
	"os"

	"github.com/iwita/go-exercises/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadAll()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete yet! 🏖")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
