package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "All you need to work with your notes",
	Long: `All you need to work with your notes`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
