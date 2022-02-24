package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CmdExecuter interface {
	Execute(ctx context.Context) error
	StopExecution() error
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rijn",
	Short: "rijn",
	Long:  `Rijn is a CLI tool that helps you to move data from subscription to another topic. So far only implemented for pubsub.. `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
