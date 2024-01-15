package plugins

import (
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

// ssh弱口令扫描
func ScanSSH(ip, port, username, password string) (err error, result bool) {
	result = false
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 3,
	}

	client, err := ssh.Dial("tcp", ip+":"+port, config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		defer session.Close()
		errRet := session.Run("echo xsec")
		if err == nil && errRet == nil {
			result = true
		}
	}
	return err, result
}
