package slack

import (
	"fmt"

	"github.com/alekstet/linux_service_checker/notifier"
)

var _ notifier.Notifier = (*SlackClient)(nil)

type SlackClient struct {
	Token string
}

func NewSlackClient(token string) *SlackClient {
	return &SlackClient{
		Token: token,
	}
}

func (client *SlackClient) Notify(service, data string) error {
	fmt.Println("from slack")
	return nil
}
