package cmd

import (
	"context"
	"jobs_golang_template/app"
	"jobs_golang_template/internal/config"
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run JOBs",
	Run:   run,
	Args:  cobra.MaximumNArgs(2),
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	log.Println("run")
	//viper
	config, err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("failed to setup viper: %s", err.Error())
	}
	application := app.NewApplication(context.TODO(), config)
	application.Setup()
}
