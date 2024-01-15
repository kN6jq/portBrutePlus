package plugins

// 声明一个函数类型
type ScanFunc func(ip, port, username, password string) (err error, result bool)

// 定义一个map用来存放插件
var ScanFuncMap map[string]ScanFunc

func init() {
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["FTP"] = ScanFtp
	ScanFuncMap["JAVADEBUG"] = ScanJavaDebug
	ScanFuncMap["MONGOD"] = ScanMongodb
	ScanFuncMap["MSSQL"] = ScanMssql
	ScanFuncMap["MYSQL"] = ScanMysql
	ScanFuncMap["ORACLE"] = ScanOracle
	ScanFuncMap["POSTGRESQL"] = ScanPostgres
	ScanFuncMap["REDIS"] = ScanRedis
	ScanFuncMap["RSYNC"] = ScanRsync
	ScanFuncMap["SMB"] = ScanSmb
	ScanFuncMap["SSH"] = ScanSSH
}
