package common

import (
	"strings"
)

var (
	PortNames = map[int]string{
		21:    "FTP",
		22:    "SSH",
		161:   "SNMP",
		445:   "SMB",
		1433:  "MSSQL",
		3306:  "MYSQL",
		5432:  "POSTGRESQL",
		6379:  "REDIS",
		9200:  "ELASTICSEARCH",
		27017: "MONGODB",
	}
	SupportProtocols map[string]bool
)

// 用于确定端口对应哪个服务
func init() {
	SupportProtocols = make(map[string]bool)
	for _, portname := range PortNames {
		SupportProtocols[strings.ToUpper(portname)] = true
	}
}
