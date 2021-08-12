package cmd

import (
	"github.com/spf13/cobra"
	"github.com/iamamirsalehi/note-cobra/data"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates needed table",
	Long: `Creates needed table`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateStudybuddyDatabase()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
