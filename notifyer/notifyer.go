package notifyer

type Notifyer interface {
	Notify(service, data string) error
}
