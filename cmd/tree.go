package cmd

import (
	"github.com/ningenme/neovenezia/pkg/service"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show file and directory structure as a tree",
	Long:  "Show file and directory structure as a tree",
	Run: func(cmd *cobra.Command, args []string) {
		service.ExecTree(args)
	},
}
