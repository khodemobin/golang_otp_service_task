package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/golang_otp_service_task/internal/app"

	"github.com/khodemobin/golang_otp_service_task/internal/server"
	"github.com/spf13/cobra"
)

func ServeCommand(app *app.App) *cobra.Command {
	cmdServe := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			Execute(app)
		},
	}
	return cmdServe
}

func Execute(app *app.App) {
	restServer := server.New(app)
	go func() {
		if err := restServer.Start(); err != nil {
			msg := fmt.Sprintf("error happen while serving: %v", err)
			app.Log.Error(errors.New(msg))
			log.Println(msg)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	app.Log.Info("Received an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		app.Log.Info("Rest server doesn't shutdown in 10 seconds")
	}
}
