package notifier

type Notifier interface {
	Notify(service, data string) error
}
