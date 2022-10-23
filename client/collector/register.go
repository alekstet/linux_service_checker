package collector

import (
	"errors"

	"github.com/alekstet/linux_service_checker/client/notifier"
	"github.com/alekstet/linux_service_checker/client/notifier/slack"
	"github.com/alekstet/linux_service_checker/client/notifier/telegram"
)

var errNoNotifiers = errors.New("error no notifiers")

func (impl *collectorImpl) registerNotifier(notifier notifier.Notifier) {
	impl.notifiers = append(impl.notifiers, notifier)
}

func (impl *collectorImpl) RegisterNotifier() error {
	if impl.config.NotifierPlatform.TelegramData.Token != "" {
		client := telegram.NewTelegramClient(impl.config.NotifierPlatform.TelegramData.Token, impl.config.NotifierPlatform.TelegramData.ChatID)
		impl.registerNotifier(client)
	}

	if impl.config.NotifierPlatform.SlackData.Token != "" {
		client := slack.NewSlackClient(impl.config.NotifierPlatform.SlackData.Token)
		impl.registerNotifier(client)
	}

	if impl.config.NotifierPlatform.SlackData.Token == "" && impl.config.NotifierPlatform.TelegramData.Token == "" {
		return errNoNotifiers
	}

	return nil
}
