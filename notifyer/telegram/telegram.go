package telegram

import (
	"fmt"

	"github.com/alekstet/linux_service_checker/notifyer"
)

var _ notifyer.Notifyer = (*TelegramClient)(nil)

type TelegramClient struct {
	Token string
}

func (client *TelegramClient) Notify(service, data string) error {
	fmt.Println("from tg")
	return nil
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		Token: token,
	}
}
