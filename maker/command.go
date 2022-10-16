package maker

import "bytes"

func (store *MakerImpl) getCommandOutput(cmd string) (string, error) {
	var stdoutBuf bytes.Buffer

	session, err := store.client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	session.Stdout = &stdoutBuf
	err = session.Run(cmd)
	if err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}
