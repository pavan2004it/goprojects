package cmd

import (
	"github.com/spf13/cobra"
	"newCli/cmd/greet"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "newCli",
	Short: "Cli to interact with cobra",
	Long:  "Cli sample",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(greet.NewGreetCommand())
	rootCmd.Version = "0.0.1"
}
