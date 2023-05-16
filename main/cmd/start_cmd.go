package cmd

import "github.com/spf13/cobra"

func StartCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start seventies",
		Long:  `This command start seventies app with a basic directory structure and configuration files.`,
		RunE:  startFunc,
	}

	return cmd
}

func startFunc(cmd *cobra.Command, args []string) error {
	// TODO: Implement initialization logic
	return nil
}
