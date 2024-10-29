package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"todo/model"
)

var sort model.StatusTask

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "Listing tasks",
	Args: func(cmd *cobra.Command, args []string) error {

		if len(listTask) < 1 {
			return errors.New("no tasks")
		} else if len(args) > 0 {
			switch args[0] {
			case "done":
				sort = model.DONE
			case "in-progress":
				sort = model.IN_PROGRESS
			case "todo":
				sort = model.TODO
			default:
				return errors.New("invalid status")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		listing()

	},
}

func listing() {

	if sort != "" {
		for _, v := range listTask {
			if v.Status == sort {
				fmt.Println("Id:", v.Id, "Description:", v.Description, "Status:", v.Status, "CreateAt:", v.CreatedAt, "UpdateAt:", v.UpdatedAt)
			}
		}
	} else {
		for _, v := range listTask {
			fmt.Println("Id:", v.Id, "Description:", v.Description, "Status:", v.Status, "CreateAt:", v.CreatedAt, "UpdateAt:", v.UpdatedAt)
		}
	}
}
func init() {
	rootCmd.AddCommand(listCmd)

}
