package cmd

import (
	"github.com/ningenme/neovenezia/pkg/setup"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Create a config file or reinitialize an existing one",
	Long: `Create a config file or reinitialize an existing one`,
	Run: func(cmd *cobra.Command, args []string) {
		setup.Exec()
	},
}