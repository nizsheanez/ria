package ssh_utils

import (
	"io"
	"os"
	"io/ioutil"
	"bufio"
	"code.google.com/p/go.crypto/ssh"
)

type SignerContainer struct {
	signers []ssh.Signer
}

func (t *SignerContainer) Key(i int) (key ssh.PublicKey, err error) {
	if i >= len(t.signers) {
		return
	}
	key = t.signers[i].PublicKey()
	return
}

func (t *SignerContainer) Sign(i int, rand io.Reader, data []byte) (sig *ssh.Signature, err error) {
	if i >= len(t.signers) {
		return
	}
	sig, err = t.signers[i].Sign(rand, data)
	return
}

func makeSigner(keyname string) (signer ssh.Signer, err error) {
	fp, err := os.Open(keyname)
	if err != nil {
		return
	}
	defer fp.Close()

	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
		return
	}

	key, err := ssh.ParseRawPrivateKey(buf)
	if err != nil {
		panic(err)
		return
	}
	signer, err = ssh.NewSignerFromKey(key)
	if err != nil {
		panic(err)
		return
	}
	return
}

func MakeKeyring() ssh.AuthMethod {
	signer, err := makeSigner(os.Getenv("HOME") + "/.ssh/alex.sharov")
	if err != nil {
		panic(err)
	}

	return ssh.PublicKeys(signer)
}

func ExecuteCmd(cmd, hostname string, config *ssh.ClientConfig, stdout chan string, stderr chan error) {
	conn, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		stderr <-err
		return
	}

	session, err := conn.NewSession()
	defer session.Close()

	if err != nil {
		stderr <-err
		return
	}

	pipe, _ := session.StdoutPipe()
	session.Start(cmd)

	buffer := bufio.NewReader(pipe)
	for {
		line, err := buffer.ReadString('\n')
		if len(line) > 0 {
			stdout <- line
		}
		if err != nil {
			if err == io.EOF {
				close(stdout)
				return
			} else {
				stderr <-err
				return
			}
		}
	}
}
