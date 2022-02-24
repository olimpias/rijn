# Rijn

Rijn is a cli tool that allows you to move message from message broker's subscription to another message broker's topic. 

Currently its only support pubsub. It could be quite handy when you want to move from especially from dead letter subscription into actual topic.

Rijn comes from the name of the [Rhine](https://en.wikipedia.org/wiki/Rhine). Purpose of this tool to move the message from one place to another, like the actualy river that moves water from The Nederlands to Switzerland. 

## Installation

1. [Install](https://go.dev/doc/install) go (golang 1.17 or later)
2. `go get github.com/olimpias/rijn`
3. you can now run rijn (from your gopath bin/ directory)

## Command line arguments

### Login-Gcd

Allows you to auth login into gcd. It will pop up web browser and once you login, it will save your credentials into application default path

Use the following command to run it.

``
rijn login-gcd
``

### Pubsub

Allows you to move pubsub message from one subscription to another subscription in same projectID. Before using this command, its mandatory to login either using `rijn login-gcd` or default terminal `gcloud` command. The credential must be in application default path

| Flag              | Description, example                                     |
|-------------------|----------------------------------------------------------|
| `-pId project-id` | String value that is locates your topic and subscription |
| `-s subscription` | Source of the messages that you want to move from        |
| `-t topic`        | Destination of the messages that you want to move to     |

##### Example

Below examples consume messages from `projects/testing/subscriptions/source-subscription` and pushes them to `projects/testing/topics/destination-topic`

```
rijn pubsub -pId testing -s source-subscription -t destination-topic
```

### Future Plans
- [ ] Support different projectIds for subscription and topic
- [ ] Add more configuration option for pubsub
- [ ] Add `brew` support
- [ ] Add support for AWS SQS
- [ ] Support moving from one specific cloud broker into another cloud broker



