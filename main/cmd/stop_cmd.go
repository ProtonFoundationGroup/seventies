package cmd

import "github.com/spf13/cobra"

func StopCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop seventies",
		Long:  `This command stop seventies app.`,
		RunE:  stopFunc,
	}

	return cmd
}

func stopFunc(cmd *cobra.Command, args []string) error {
	// TODO: Implement initialization logic
	return nil
}
