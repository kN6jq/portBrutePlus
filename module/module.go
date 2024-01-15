package module

// 定义一个服务所需要的结构体
type Service struct {
	Ip       string
	Port     int
	Protocol string
	UserName string
	PassWord string
}

// 定义一个扫描结果的结构体
type ScanResult struct {
	Service Service
	Result  bool
}

// 定义一个ip地址结构体
type IpAddr struct {
	Ip       string
	Port     int
	Protocol string
}
