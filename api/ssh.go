package api

import (
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

var errUnknownAuthMethod = errors.New("error unknown auth method")

func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func (store *Store) getCredsConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            store.Config.MonitoringServer.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(store.Config.MonitoringServer.Password),
		},
	}
}

func (store *Store) getCertificateConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            store.Config.MonitoringServer.Username,
		Auth: []ssh.AuthMethod{
			publicKeyFile(store.Config.MonitoringServer.PathToPublicKey),
		},
	}
}

func (store *Store) TestSSHconnection() error {
	var sshConfig *ssh.ClientConfig
	switch store.Config.MonitoringServer.Type {
	case "creds":
		sshConfig = store.getCredsConfig()
	case "certificate":
		sshConfig = store.getCertificateConfig()
	default:
		return errUnknownAuthMethod
	}

	sshClient, err := ssh.Dial("tcp", store.Config.MonitoringServer.ServerURL+store.Config.MonitoringServer.ServerPort, sshConfig)
	if err != nil {
		return err
	}

	store.Client = sshClient

	return nil
}
