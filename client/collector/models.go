package collector

type ServicesInfo map[string]ServiceInfo

type ServiceInfo struct {
	Name        string
	Description string
	Loaded      string
	Active      string
	Journal     string
}
