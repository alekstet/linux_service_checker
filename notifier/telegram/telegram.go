package telegram

import (
	"fmt"

	"github.com/alekstet/linux_service_checker/notifier"
)

var _ notifier.Notifier = (*TelegramClient)(nil)

type TelegramClient struct {
	Token string
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		Token: token,
	}
}

func (client *TelegramClient) Notify(service, data string) error {
	fmt.Println("from tg")
	return nil
}
