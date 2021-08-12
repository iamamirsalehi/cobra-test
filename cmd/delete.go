package cmd

import (
	"github.com/iamamirsalehi/note-cobra/data"

	"github.com/spf13/cobra"
)
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the notes you created",
	Long: `Delete the notes you created`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteNote()
	},
}

func init() {
	noteCmd.AddCommand(deleteCmd)
}

func deleteNote(){
	noteId := promptContent{
		errorMsg: "Note ID",
		label: "Enter the ID of the note that you want to delete",
	}

	id := promptGetInput(noteId)

	data.DeleteNote(id)
}