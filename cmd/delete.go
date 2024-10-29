package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var removeCmd = &cobra.Command{
	Use:     "delete [ID]",
	Short:   "Deleting task",
	Args:    CheckerArgs(1),
	Example: "todo delete 1",
	Run: func(cmd *cobra.Command, args []string) {
		remove(args[0])
	},
}

func remove(taskID string) {

	id, _ := strconv.Atoi(taskID)
	found := false

	for i := range listTask {
		if listTask[i].Id == id {
			found = true
			fmt.Printf("Task deleted successfully (ID: %d)", listTask[i].Id)
			listTask = append(listTask[:i], listTask[i+1:]...)
			break
		}
	}

	if !found {
		fmt.Printf("task %d not found", id)
	} else {
		saveData()
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
