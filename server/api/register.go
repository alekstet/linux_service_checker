package api

import (
	"errors"

	"github.com/alekstet/linux_service_checker/server/notifier"
	"github.com/alekstet/linux_service_checker/server/notifier/slack"
	"github.com/alekstet/linux_service_checker/server/notifier/telegram"
)

var errNoNotifiers = errors.New("error no notifiers")

func (store *store) registerNotifier(notifier notifier.Notifier) {
	store.notifiers = append(store.notifiers, notifier)
}

func (store *store) RegisterNotifier() error {
	if store.config.NotifierPlatform.TelegramData.Token != "" {
		client := telegram.NewTelegramClient(store.config.NotifierPlatform.TelegramData.Token, store.config.NotifierPlatform.TelegramData.ChatID)
		store.registerNotifier(client)
	}

	if store.config.NotifierPlatform.SlackData.Token != "" {
		client := slack.NewSlackClient(store.config.NotifierPlatform.SlackData.Token)
		store.registerNotifier(client)
	}

	if store.config.NotifierPlatform.SlackData.Token == "" && store.config.NotifierPlatform.TelegramData.Token == "" {
		return errNoNotifiers
	}

	return nil
}
