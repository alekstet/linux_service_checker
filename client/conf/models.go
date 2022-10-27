package conf

type Config struct {
	MonitoringServer *MonitoringServer `yaml:"monitoring_server"`
	ExecutionServer  *ExecutionServer  `yaml:"execution_server"`
	CollectorServer  *ExecutionServer  `yaml:"collector_server"`
	Database
}

type Database struct {
	ConnectionString string `yaml:"connection_string"`
}

type CollectorServer struct {
	ServerURL  string `yaml:"server_url"`
	ServerPort string `yaml:"server_port"`
}

type ExecutionServer struct {
	ServerURL  string `yaml:"server_url"`
	ServerPort string `yaml:"server_port"`
}

type MonitoringServer struct {
	ServicesNames []string `yaml:"services_names"`
}
