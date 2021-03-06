package cmd

import (
	"github.com/olimpias/rijn/clouds/gcd/pubsub"
	"github.com/olimpias/rijn/log"

	"github.com/spf13/cobra"
)

var (
	moveCmdExecuter CmdExecuter
	config          pubsub.Config
)

var pubsubCmd = &cobra.Command{
	Use:   "pubsub",
	Short: "",
	Long:  `Runs the following command "gcloud auth application-default login" login into the gcloud`,
	PreRun: func(cmd *cobra.Command, args []string) {
		moveCmd, err := pubsub.NewMoveCmd(cmd.Context(), config)
		if err != nil {
			log.Fatal(err.Error())
		}

		moveCmdExecuter = moveCmd
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := moveCmdExecuter.Execute(cmd.Context()); err != nil {
			log.Fatal(err.Error())
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		if err := moveCmdExecuter.StopExecution(); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	pubsubCmd.PersistentFlags().StringVarP(&config.TopicProjectID, "topicProjectId", "p", "", "topic project Id")
	pubsubCmd.PersistentFlags().StringVarP(&config.SubscriptionProjectID, "subscriptionProjectId", "c", "", "subscription project Id")
	pubsubCmd.PersistentFlags().StringVarP(&config.Topic, "topic", "t", "", "topic name")
	pubsubCmd.PersistentFlags().StringVarP(&config.Subscription, "subscription", "s", "", "subscription name")
	rootCmd.AddCommand(pubsubCmd)
}
