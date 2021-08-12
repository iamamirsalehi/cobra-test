package cmd

import (
	"fmt"
	"github.com/iamamirsalehi/note-cobra/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a note that you created",
	Long:  `Update a note that you created`,
	Run: func(cmd *cobra.Command, args []string) {
		updateNote()
	},
}

func init() {
	noteCmd.AddCommand(updateCmd)
}

func promptGetSelectUpdate(pc promptContent) string {
	items := []string{"word", "definition", "category"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.label,
			Items: items,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func updateNote() {
	idNote := promptContent{
		errorMsg: "Enter the note id",
		label:    "What is the id of the note you want to update? ",
	}

	id := promptGetInput(idNote)

	columnName := promptContent{
		errorMsg: "Enter the field ",
		label:    "What the name of the field you want to update? ",
	}

	column := promptGetSelectUpdate(columnName)

	if column == "category" {
		newCategory := promptContent{
			errorMsg: "Select the category",
			label:    "What category you select for this note to update? ",
		}

		categoryNewValue := promptGetSelect(newCategory)

		data.UpdateNote(id, column, categoryNewValue)

		os.Exit(1)
	}

	newValue := promptContent{
		errorMsg: "Enter the new value",
		label:    fmt.Sprintf("What is the new value of note(%s)? ", id),
	}

	value := promptGetInput(newValue)

	data.UpdateNote(id, column, value)
}
