package ssh

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"log"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/kayrus/putty"
	"golang.org/x/crypto/ssh"
)

var errUnknownAuthMethod = errors.New("error unknown auth method")

func publicKeyFile(file string) ssh.AuthMethod {
	/* buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	} */

	var privateKey interface{}

	// read the key
	puttyKey, err := putty.NewFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = puttyKey.ParseRawPrivateKey(nil)
	if err != nil {
		log.Fatal(err)
	}

	value, _ := privateKey.(*rsa.PrivateKey)

	bb := x509.MarshalPKCS1PrivateKey(value)

	key, err := ssh.ParsePrivateKey(bb)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func getCredsConfig(config *conf.Config) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            config.MonitoringServer.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.MonitoringServer.Password),
		},
	}
}

func getCertificateConfig(config *conf.Config) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            config.MonitoringServer.Username,
		Auth: []ssh.AuthMethod{
			publicKeyFile(config.MonitoringServer.PathToPublicKey),
		},
	}
}

func getClientConfig(config *conf.Config) (*ssh.ClientConfig, error) {
	var sshConfig *ssh.ClientConfig
	switch config.MonitoringServer.Type {
	case "creds":
		sshConfig = getCredsConfig(config)
	case "certificate":
		fmt.Println("here")
		sshConfig = getCertificateConfig(config)
		fmt.Println(sshConfig)
	default:
		return nil, errUnknownAuthMethod
	}

	return sshConfig, nil
}

func GetClient(config *conf.Config) (*ssh.Client, error) {
	sshConfig, err := getClientConfig(config)
	if err != nil {
		return nil, err
	}

	sshClient, err := ssh.Dial("tcp", config.MonitoringServer.ServerURL+config.MonitoringServer.ServerPort, sshConfig)
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}
