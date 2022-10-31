package telegram

import (
	"fmt"
	"log"
	"sync"

	"github.com/alekstet/linux_service_checker/server/notifier"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ notifier.Notifier = (*TelegramClient)(nil)

type TelegramClient struct {
	token  string
	chatID int64
	bot    *tgbotapi.BotAPI
}

func NewTelegramClient(token string, chatID int64) *TelegramClient {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	return &TelegramClient{
		token:  token,
		bot:    bot,
		chatID: chatID,
	}
}

func (client *TelegramClient) Notify(service, curStatus, exStatus string, wg *sync.WaitGroup) error {
	defer wg.Done()

	text := fmt.Sprintf("service '%s' changes status from: '%s' to '%s'", service, exStatus, curStatus)
	msg := tgbotapi.NewMessage(client.chatID, text)
	_, err := client.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
