package maker

import (
	"bytes"
)

func (store *makerImpl) getCommandOutput(cmd string) (string, error) {
	var stdoutBuf bytes.Buffer

	session, err := store.client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	session.Stdout = &stdoutBuf
	session.Run(cmd)

	return stdoutBuf.String(), nil
}
