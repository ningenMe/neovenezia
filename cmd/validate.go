package cmd

import (
	"github.com/ningenme/neovenezia/pkg/service"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate package structure by neovenezia.yaml",
	Long: "Validate package structure by neovenezia.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		service.ExecValidate()
	},
}