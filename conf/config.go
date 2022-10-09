package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServicesNames []string `yaml:"services_names"`
	SSHserverName string   `yaml:"ssh_server_name"`
	SSHserverPort string   `yaml:"ssh_server_port"`
	ServerPort    string   `yaml:"server_port"`
	*AuthMethod   `yaml:"auth_method"`
}

type AuthMethod struct {
	Type            string `yaml:"type"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	PathToPublicKey string `yaml:"path_to_public_key"`
}

func ReadConfig() (*Config, error) {
	file, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
