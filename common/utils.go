package common

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"portBrutePlus/module"
	"strconv"
	"strings"
)

// 读取账号密码文件
func ReadDictFile(file string) (data []string, err error) {
	filename, err := os.Open(file)
	if err != nil {
		log.Fatalf("open file fail, %v", err)
	}
	// 关闭文件
	defer filename.Close()
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			data = append(data, line)
		}
	}
	return data, err
}

// 读取待爆破的文件
// 实现从爆破文件读取,如果包含协议则判断协议是否支持，如果不包含，则通过端口识别服务
func ReadBruteFile(file string) (ipList []module.IpAddr) {
	filename, err := os.Open(file)
	if err != nil {
		log.Fatalf("open brute file fail, %v", err)
	}
	// 关闭文件
	defer filename.Close()
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// 跳过后续的处理逻辑，直接进行下一次循环迭代
			continue
		}
		ipPortProtocol := strings.TrimSpace(line)
		ipp := strings.Split(ipPortProtocol, ":")
		ip := ipp[0]           // 这里是ip
		portProtocol := ipp[1] // 这里是port
		pp := strings.Split(portProtocol, "|")
		if len(pp) == 2 {
			// 如果包含自定义的协议
			port, _ := strconv.Atoi(pp[0])
			protocol := strings.ToUpper(pp[1])
			// 如果输入的协议已经定义过,则代表被支持
			if SupportProtocols[strings.ToUpper(protocol)] {
				addr := module.IpAddr{Ip: ip, Port: port, Protocol: protocol}
				ipList = append(ipList, addr)
			} else {
				// 代表输入的协议不存在
				log.Fatalln("Protocol not Support ")
			}
		} else {
			// 默认处理
			port, err := strconv.Atoi(portProtocol)
			if err == nil {
				protocol := PortNames[port]
				if SupportProtocols[protocol] {
					addr := module.IpAddr{Ip: ip, Port: port, Protocol: protocol}
					ipList = append(ipList, addr)
				}
			}
		}
	}
	return ipList
}

// 对字符串进行md5
func Hash(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func WriteToFile(wireteString, filename string) {

	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(wireteString)
	fd.Write(buf)

}
