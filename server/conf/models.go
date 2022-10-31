package conf

import "time"

type Config struct {
	MonitoringServer *MonitoringServer `yaml:"monitoring_server"`
	ExecutionServer  *ExecutionServer  `yaml:"execution_server"`
	NotifierPlatform *NotifierPlatform `yaml:"notifier_platform"`
	Database         *Database         `yaml:"database"`
	Log              *Log              `yaml:"log"`
}

type Log struct {
	LogFile string `yaml:"log_file"`
}

type ExecutionServer struct {
	ServerPort string        `yaml:"server_port"`
	PollPeriod time.Duration `yaml:"poll_period"`
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
