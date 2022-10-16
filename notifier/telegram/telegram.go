package telegram

import (
	"log"
	"sync"

	"github.com/alekstet/linux_service_checker/notifier"

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

func (client *TelegramClient) Notify(service, exStatus, curStatus string, wg *sync.WaitGroup) error {
	defer wg.Done()

	msg := tgbotapi.NewMessage(client.chatID, service+" change status from: "+exStatus+" to: "+curStatus)
	client.bot.Send(msg)

	return nil
}
