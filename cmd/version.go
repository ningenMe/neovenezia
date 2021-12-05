package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ningenme/neovenezia/pkg/service"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current Neovenezia version",
	Long:  `Show the current Neovenezia version`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ExecVersion()
	},
}
