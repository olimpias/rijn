package login

import (
	"context"
	"fmt"
	"os/exec"
)

type LoginCmd struct {
}

func NewLoginCmd() *LoginCmd {
	return &LoginCmd{}
}

// Execute logins into your gcloud account and creates a certification in application-default path. It will pop up google accounts to login after that gonna store default application credentials
// Referance here: https://cloud.google.com/sdk/gcloud/reference/auth/application-default/login
func (cmd *LoginCmd) Execute(ctx context.Context) error {
	terminalCmd := exec.CommandContext(ctx, "gcloud", "auth", "application-default", "login")
	stdOutput, err := terminalCmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(string(stdOutput))
	return nil
}
