package plugins

import (
	"database/sql"
	"fmt"
	_ "github.com/netxfly/mysql"
)

func ScanMysql(ip, port, username, password string) (error, bool) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/?timeout=%ds", username, password, ip+":"+port, 3)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err, false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err, false
	}

	return nil, true
}
