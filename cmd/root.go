package cmd

import (
	"fmt"
	"github.com/ningenme/neovenezia/pkg/service"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "neovenezia",
	Short: "Neovenezia is a validator of package config",
	Long:  `Neovenezia is a cli application that manages package config, directory structures`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			service.ExecVersion()
			return
		}
		cobra.CheckErr(cmd.Help())
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.Flags().BoolP("version", "v", false, "An alias for the `version` subcommand.")
}

func initConfig() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
