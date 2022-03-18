package pubsub

import (
	"context"
	"errors"
	"fmt"

	"github.com/olimpias/rijn/log"

	"cloud.google.com/go/pubsub"
)

type MoveCmd struct {
	subscriberClient *pubsub.Client
	publisherClient  *pubsub.Client
	Config
}

type Config struct {
	Topic                 string
	Subscription          string
	TopicProjectID        string
	SubscriptionProjectID string
}

func (c Config) Validate() error {
	if c.TopicProjectID == "" {
		return errors.New("topic projectID is empty")
	}

	if c.SubscriptionProjectID == "" {
		return errors.New("subscription projectID is empty")
	}

	if c.Subscription == "" {
		return errors.New("subscription is empty")
	}

	if c.Topic == "" {
		return errors.New("topic is empty")
	}

	return nil
}

func NewMoveCmd(ctx context.Context, config Config) (*MoveCmd, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("unable to validate config err:%w", err)
	}

	subscriberClient, err := pubsub.NewClient(ctx, config.SubscriptionProjectID)
	if err != nil {
		return nil, fmt.Errorf("unable to create subscriber client err:%w", err)
	}

	publisherClient, err := pubsub.NewClient(ctx, config.TopicProjectID)
	if err != nil {
		return nil, fmt.Errorf("unable to create publisher client err:%w", err)
	}

	return &MoveCmd{
		subscriberClient: subscriberClient,
		publisherClient:  publisherClient,
		Config:           config,
	}, nil
}

func (cmd *MoveCmd) Execute(ctx context.Context) error {
	publisherTopic := cmd.publisherClient.Topic(cmd.Config.Topic)
	if err := cmd.subscriberClient.Subscription(cmd.Subscription).Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		result := publisherTopic.Publish(ctx, message)
		if _, err := result.Get(ctx); err != nil {
			log.Fatal(fmt.Sprintf("unable to publish message to the topic err %s", err))
		}
		message.Ack()
	}); err != nil {
		return fmt.Errorf("unable to perform subscription err %w", err)
	}

	return nil
}

func (cmd *MoveCmd) StopExecution() error {
	if err := cmd.publisherClient.Close(); err != nil {
		return fmt.Errorf("unable to close publisher client err: %w", err)
	}
	if err := cmd.subscriberClient.Close(); err != nil {
		return fmt.Errorf("unable to close subscriber client err: %w", err)
	}

	return nil
}
