package api

import (
	"errors"

	"github.com/alekstet/linux_service_checker/notifier"
	"github.com/alekstet/linux_service_checker/notifier/slack"
	"github.com/alekstet/linux_service_checker/notifier/telegram"
)

var errNoNotifiers = errors.New("error no notifiers")

func (store *Store) registerNotifier(notifier notifier.Notifier) {
	store.notifiers = append(store.notifiers, notifier)
}

func (store *Store) RegisterNotifier() error {
	if store.config.NotifierPlatform.TelegramData.Token != "" {
		client := telegram.NewTelegramClient(store.config.NotifierPlatform.TelegramData.Token)
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
