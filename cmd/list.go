package cmd

import (
	"github.com/spf13/cobra"
	"github.com/iamamirsalehi/note-cobra/data"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "returns all created notes",
	Long: `returns all created notes`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
