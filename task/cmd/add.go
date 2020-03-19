package cmd

import (
	"fmt"
	"strings"

	"github.com/iwita/go-exercises/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your tasks list.",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")

		task := strings.Join(args, " ") // Joins the strings, so "hello world" and "hello" "world" is the exact thing
		//fmt.Printf("Added \"%s\" to your task list\n", task)
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

// A function that can be run before main function in the main package
// You need this in order to setup
func init() {
	RootCmd.AddCommand(addCmd)
}
