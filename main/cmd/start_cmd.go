package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ProtonFundationGroup/seventies/v2/log"
	"github.com/ProtonFundationGroup/seventies/v2/server"
	"github.com/spf13/cobra"
)

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

	var (
		err error
		srv *server.Server
	)
	go func() {
		srv, err = server.New()
		if err != nil {
			log.Errorf("create Server failed, err = %v", err)
			os.Exit(1)
		}
		srv.Start()
	}()

	// Listen for quit signal and exit
	signalChan := make(chan os.Signal, 3)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-signalChan
	log.PrintGreen("catch `SIG_TERM | SIG_KILL | SIGINT` signal and exit.")
	srv.Stop()

	return nil
}
