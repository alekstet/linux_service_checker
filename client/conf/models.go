package conf

type Config struct {
	MonitoringServer *MonitoringServer `yaml:"monitoring_server"`
	ExecutionServer  *ExecutionServer  `yaml:"execution_server"`
	CollectorServer  *ExecutionServer  `yaml:"collector_server"`
	NotifierPlatform *NotifierPlatform `yaml:"notifier_platform"`
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

type NotifierPlatform struct {
	TelegramData *TelegramData `yaml:"telegram"`
	SlackData    *SlackData    `yaml:"slack"`
}

type TelegramData struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

type SlackData struct {
	Token string `yaml:"token"`
}
