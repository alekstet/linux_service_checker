package notifier

import "sync"

type Notifier interface {
	Notify(service, exStatus, curStatus string, wg *sync.WaitGroup) error
}
