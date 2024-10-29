package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"todo/model"
)

var todoCmd = &cobra.Command{
	Use:   "mark-in-todo",
	Args:  CheckerArgs(1),
	Short: "Mark task in todo",
	Run: func(cmd *cobra.Command, args []string) {
		mark(args[0], model.TODO)
	},
}

var progressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Args:  CheckerArgs(1),
	Short: "Mark task in progress",
	Run: func(cmd *cobra.Command, args []string) {
		mark(args[0], model.IN_PROGRESS)
	},
}

var doneCmd = &cobra.Command{
	Use:   "mark-in-done",
	Args:  CheckerArgs(1),
	Short: "Mark task in done",
	Run: func(cmd *cobra.Command, args []string) {
		mark(args[0], model.DONE)
	},
}

func mark(taskID string, status model.StatusTask) {
	id, _ := strconv.Atoi(taskID)

	for i := range listTask {

		if listTask[i].Id == id {
			listTask[i].Status = status
			fmt.Printf("Task mark successfully (ID: %d)", listTask[i].Id)
		}
	}

	saveData()
}

func init() {
	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(progressCmd)
	rootCmd.AddCommand(doneCmd)
}
