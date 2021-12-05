package cmd

import (
	"github.com/ningenme/neovenezia/pkg/service"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file or reinitialize an existing one",
	Long:  `Create a config file or reinitialize an existing one`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ExecInit()
	},
}
