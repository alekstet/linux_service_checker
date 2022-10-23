package conf

type Config struct {
	MonitoringServer *MonitoringServer `yaml:"monitoring_server"`
	ExecutionServer  *ExecutionServer  `yaml:"execution_server"`
	Database         *Database         `yaml:"database"`
}

type ExecutionServer struct {
	ServerPort string `yaml:"server_port"`
}

type Database struct {
	ConnectionString string `yaml:"connection_string"`
}

type MonitoringServer struct {
	ServicesNames []string `yaml:"services_names"`
	*AuthMethod   `yaml:"auth_method"`
	ServerURL     string `yaml:"server_url"`
	ServerPort    string `yaml:"server_port"`
}

type AuthMethod struct {
	Type            string `yaml:"type"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	PathToPublicKey string `yaml:"path_to_public_key"`
}
