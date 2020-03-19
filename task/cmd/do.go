package cmd

import (
	"fmt"
	"strconv"

	"github.com/iwita/go-exercises/task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("do called")
		//Usage task do 1 2 3 (task id)

		// Store ids in an integer slice
		var ids []int
		for _, arg := range args {
			// strconv package
			// Converts strings into other datatypes
			// ParseInt: More generic form, so as to parse various formats of integers such as Hex.
			// Atoi: ASCI to Integer conversion with a base of 10. (Uses ParseInt behind the scenes)
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the arguement:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ReadAll()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid Task Number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark %d as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf("Task %d was deleted successfully\n", id)
			}

		}
		//fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
