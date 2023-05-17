package main

import (
	"fmt"
	"os"

	"github.com/ProtonFundationGroup/seventies/v2/main/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "seventies",
		Short: "Seventies app does awesome things",
	}

	initCmd := cmd.InitCmd()
	startCmd := cmd.StartCmd()

	// Add commands to root command
	rootCmd.AddCommand(initCmd, startCmd)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
