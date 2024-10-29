package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var updateCmd = &cobra.Command{
	Use:     "update [ID] [DESCRIPTION]",
	Args:    CheckerArgs(2),
	Short:   "Updating task",
	Example: "update 0 placeholder",
	Run: func(cmd *cobra.Command, args []string) {
		update(args[0], strings.Join(args[1:], " "))
	},
}

func update(taskID string, description string) {

	id, _ := strconv.Atoi(taskID)

	found := false

	for i := range listTask {
		if listTask[i].Id == id {
			found = true
			listTask[i].Description = description
			listTask[i].UpdatedAt = time.Now()
			fmt.Printf("Task updated successfully (ID: %d)", listTask[i].Id)
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
	rootCmd.AddCommand(updateCmd)
}
