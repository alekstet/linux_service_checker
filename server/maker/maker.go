package maker

type Maker interface {
	Make(service, action string) error
}
