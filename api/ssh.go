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
		User:            store.Config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(store.Config.Password),
		},
	}
}

func (store *Store) getCertificateConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            store.Config.Username,
		Auth: []ssh.AuthMethod{
			publicKeyFile(store.Config.PathToPublicKey),
		},
	}
}

func (store *Store) TestSSHconnection() error {
	var sshConfig *ssh.ClientConfig
	switch store.Config.Type {
	case "creds":
		sshConfig = store.getCredsConfig()
	case "certificate":
		sshConfig = store.getCertificateConfig()
	default:
		return errUnknownAuthMethod
	}

	sshClient, err := ssh.Dial("tcp", store.Config.SSHserverName+store.Config.SSHserverPort, sshConfig)
	if err != nil {
		return err
	}

	store.Client = sshClient

	return nil
}
