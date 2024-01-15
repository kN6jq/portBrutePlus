package plugins

import (
	"github.com/jlaffaye/ftp"
	"time"
)

func ScanFtp(ip, port, username, password string) (err error, result bool) {
	result = false
	c, err := ftp.Dial(ip+":"+port, ftp.DialWithTimeout(3*time.Second))
	if err == nil {
		err = c.Login(username, password)
		if err == nil {
			result = true
			c.Logout()
		}
	}
	return err, result
}
