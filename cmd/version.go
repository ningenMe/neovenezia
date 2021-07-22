package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ningenme/neovenetia/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current Neovenetia version",
	Long: `Show the current Neovenetia version`,
	Run: func(cmd *cobra.Command, args []string) {
		version.Main()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
