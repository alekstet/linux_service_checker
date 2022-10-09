package slack

import (
	"fmt"

	"github.com/alekstet/linux_service_checker/notifyer"
)

var _ notifyer.Notifyer = (*SlackClient)(nil)

type SlackClient struct {
	Token string
}

func (client *SlackClient) Notify(service, data string) error {
	fmt.Println("from slack")
	return nil
}

func NewSlackClient(token string) *SlackClient {
	return &SlackClient{
		Token: token,
	}
}
