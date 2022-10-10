package conf

type Config struct {
	MonitoringServer *MonitoringServer `yaml:"monitoring_server"`
	ExecutionServer  *ExecutionServer  `yaml:"execution_server"`
	NotifierPlatform *NotifierPlatform `yaml:"notifier_platform"`
}

type ExecutionServer struct {
	FrontendPath string `yaml:"frontend_path"`
	ServerPort   string `yaml:"server_port"`
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
	Token string `yaml:"token"`
}

type SlackData struct {
	Token string `yaml:"token"`
}
