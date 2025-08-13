package main

import (
	"log"

	"github.com/khodemobin/golang_boilerplate/internal/app"

	"github.com/khodemobin/golang_boilerplate/cmd"
	"github.com/spf13/cobra"
)

// @title           Golang OTP Testcase
// @version         1.0

// @host      localhost:3000
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api
func main() {
	appVar := app.New()

	rootCmd := &cobra.Command{
		Use:                "app",
		DisableAutoGenTag:  true,
		DisableSuggestions: true,
		Run: func(c *cobra.Command, args []string) {
			cmd.Execute(appVar)
		},
	}
	rootCmd.AddCommand(cmd.ServeCommand(appVar))
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
