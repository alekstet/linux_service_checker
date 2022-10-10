package notifier

import "sync"

type Notifier interface {
	Notify(service, data string, wg *sync.WaitGroup) error
}
