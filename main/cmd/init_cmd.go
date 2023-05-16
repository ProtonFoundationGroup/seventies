package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitCmd() *cobra.Command {

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize seventies",
		Long:  `This command initializes seventies app with a basic directory structure and configuration files.`,
		RunE:  initFunc,
	}

	// Define flags for each command
	initCmd.Flags().String("config", "", "config file (default is seventies.yaml)")

	// Bind flags to viper
	viper.BindPFlag("config", initCmd.Flags().Lookup("config"))

	return initCmd
}

func initFunc(cmd *cobra.Command, args []string) error {
	// TODO: Implement initialization logic
	return nil
}
