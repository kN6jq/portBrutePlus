package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"

	"log"
	"os"
	"portBrutePlus/brute"
	"portBrutePlus/common"
	"time"
)

var (
	h      bool
	ips    string
	user   string
	pass   string
	thread int
)

func main() {
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.StringVar(&ips, "ips", "ips.txt", "地址文件")
	flag.StringVar(&user, "user", "user.txt", "用户名文件")
	flag.StringVar(&pass, "pass", "pass.txt", "密码文件")
	flag.IntVar(&thread, "thread", 100, "线程数量")
	flag.Parse()
	// 打印帮助信息
	if h {
		fmt.Fprintf(os.Stderr, `支持协议FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD.快速开始: ./PortBruteMac`)
		flag.PrintDefaults()
		return
	}
	// 判断用户是否输入了文件
	if flag.Lookup("ips") == nil || flag.Lookup("user") == nil || flag.Lookup("pass") == nil {
		color.Red("file not empty")
		return
	}
	startTime := time.Now()
	userDict, uerr := common.ReadDictFile(user)
	passDict, perr := common.ReadDictFile(pass)
	ipList := common.ReadBruteFile(ips)
	if uerr == nil && perr == nil {
		// 创建任务
		tasks := brute.GenerateTask(ipList, userDict, passDict)
		color.Cyan("Number of all task : %d", len(tasks))
		brute.RunTasks(tasks, thread)
	} else {
		log.Println("Read file error!!!")
	}
	color.Red("Run Time is : %s\n", time.Since(startTime))
}
