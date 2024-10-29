package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"time"
	"todo/model"
)

var addCmd = &cobra.Command{
	Use:   "add [DESCRIPTION]",
	Short: "Adding a new task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no description")
		}
		return nil
	},
	Example: "add placeholder",
	Run: func(cmd *cobra.Command, args []string) {
		add(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
func add(description string) {

	var id = 1

	if len(listTask) > 0 {
		id = listTask[len(listTask)-1].Id + 1
	}

	task := model.Task{
		Id:          id,
		Description: description,
		Status:      model.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	listTask = append(listTask, task)
	fmt.Printf("Task added successfully (ID: %d)", task.Id)
	saveData()
}
