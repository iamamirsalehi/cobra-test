package cmd

import (
	"errors"
	"fmt"
	"github.com/iamamirsalehi/note-cobra/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create your notes",
	Long:  `Create your notes`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) < 0 {
			return errors.New(pc.errorMsg)
		}

		return nil
	}

	template := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold}}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: template,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input %v\n", result)

	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"PHP", "Go", "Javascript", "Java"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		errorMsg: "Please provide a word",
		label:    "What word would you like to make note of? ",
	}

	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		errorMsg: "Please provide a definition",
		label:    fmt.Sprintf("What is the definition of %s? ", word),
	}

	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		errorMsg: "Please provide a category",
		label: fmt.Sprintf("What cateogy does %s belongs to? ", word),
	}

	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
