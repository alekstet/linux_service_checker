package telegram

import (
	"fmt"
	"sync"

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

func (client *TelegramClient) Notify(service, data string, wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("from tg")
	return nil
}
