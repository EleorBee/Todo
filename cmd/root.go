package cmd

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"todo/model"
)

func CheckerArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(listTask) < n {
			return errors.New("no tasks")
		} else if len(args) < n {
			return errors.New("too few arguments")
		} else if len(args) > n {
			return errors.New("too many arguments")
		}
		return nil
	}
}

var rootCmd = &cobra.Command{
	Use: "todo",
}

var listTask []model.Task

const pathData = "data/TodoData.json"

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func loadData() {

	if _, err := os.Stat(pathData); os.IsNotExist(err) {
		file, err := os.Create(pathData)

		if err != nil {

			errMK := os.Mkdir("data", os.ModePerm)

			if errMK != nil {
				log.Fatal(errMK)
			}

			file, _ := os.Create(pathData)

			errClose := file.Close()

			if errClose != nil {
				log.Fatal(errClose)
			}

		} else {

			errClose := file.Close()

			if errClose != nil {
				log.Fatal(errClose)
			}
		}

	} else {
		file, err := os.ReadFile(pathData)

		if err != nil {
			log.Fatal(err)
		}
		if len(file) > 1 {
			loadError := json.Unmarshal(file, &listTask)

			if loadError != nil {
				log.Fatal(loadError)
			}
		}
	}
}
func saveData() {

	var errorFile, errorJson error

	var value []byte

	value, errorJson = json.Marshal(&listTask)

	if errorJson != nil {
		log.Fatal(errorJson)
	}

	errorFile = os.WriteFile(pathData, value, os.ModePerm)

	if errorFile != nil {
		log.Fatal(errorFile)
	}

}

func init() {
	loadData()
}
