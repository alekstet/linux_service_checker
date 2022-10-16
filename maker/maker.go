package maker

type Maker interface {
	Collect() *ServicesInfo
	Make(service, action string) error
}
