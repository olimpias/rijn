package cmd

import (
	"fmt"

	"github.com/olimpias/rijn/clouds/gcd/login"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login-gcd",
	Short: "Logins into gcloud",
	Long:  `Runs the following command "gcloud auth application-default login" login into the gcloud`,
	Run: func(cmd *cobra.Command, args []string) {
		login := login.NewLoginCmd()
		if err := login.Execute(cmd.Context()); err != nil {
			panic(fmt.Sprintf("unable to login gcloud err:%s", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
