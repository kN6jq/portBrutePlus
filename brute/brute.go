package brute

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"portBrutePlus/common"
	"portBrutePlus/module"
	"portBrutePlus/plugins"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	lock        sync.Mutex
	successHash map[string]bool
	bruteResult map[string]module.Service
)

func GenerateTask(ipList []module.IpAddr, userDict []string, passDict []string) (scanTasks []module.Service) {
	// 所有列表，添加任务
	for _, user := range userDict {
		for _, pass := range passDict {
			for _, addr := range ipList {
				scanTask := module.Service{Ip: addr.Ip, Port: addr.Port, Protocol: addr.Protocol, UserName: user, PassWord: pass}
				scanTasks = append(scanTasks, scanTask)
			}
		}
	}
	return
}

func RunTasks(tasks []module.Service, thread int) {
	wg := &sync.WaitGroup{}
	successHash = make(map[string]bool)
	bruteResult = make(map[string]module.Service)

	// 创建扫描线程*2的channel作为缓冲
	taskChan := make(chan module.Service, thread*2)

	for i := 0; i < thread; i++ {
		go runBrute(taskChan, wg)
	}
	sumTotal := pb.StartNew(len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
		sumTotal.Increment()
	}

	close(taskChan)
	sumTotal.Finish()
	wg.Wait()
	waitTimeout(wg, 3*time.Second)
	common.WriteToFile("\n全部扫描完成\n", "res.txt")
	color.Red("Scan complete. %d vulnerabilities found ! \n", len(bruteResult))
	for _, service := range bruteResult {
		color.Cyan("[+] %s %d %s %s \n", service.Ip, service.Port, service.UserName, service.PassWord)
	}
}

func runBrute(taskChans chan module.Service, wg *sync.WaitGroup) {
	for target := range taskChans {
		protocol := strings.ToUpper(target.Protocol)
		// 做一次hash标记
		var k string
		if protocol == "REDIS" || protocol == "FTP" || protocol == "SNMP" || protocol == "POSTGRESQL" || protocol == "SSH" {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.UserName)
		}
		hash := common.Hash(k)
		// 已经爆破成功
		if checkTashHash(hash) {
			wg.Done()
			continue // 跳过当前循环的剩余代码
		}
		err, res := plugins.ScanFuncMap[protocol](target.Ip, strconv.Itoa(target.Port), target.UserName, target.PassWord)
		if err == nil && res == true {
			saveRes(target, hash)
		}
		wg.Done()
	}
}

// 爆破成功
func saveRes(target module.Service, hash string) {
	setTaskHash(hash)
	_, ok := bruteResult[hash] // 结果里面是否有值,没有的话写结果并且添加到这map里面
	if !ok {
		lock.Lock()
		s := fmt.Sprintf("[+] %s %d %s %s  \n", target.Ip, target.Port, target.UserName, target.PassWord)
		common.WriteToFile(s, "res.txt")
		bruteResult[hash] = module.Service{Ip: target.Ip, Port: target.Port, UserName: target.UserName, PassWord: target.PassWord}
		lock.Unlock()
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}

// 多线程读取map
// 判断是否已经爆破成功
func checkTashHash(hash string) bool {
	ok := successHash[hash]
	return ok
}

// 多线程写map
// 用户标记爆破成功
func setTaskHash(hash string) {
	lock.Lock()
	successHash[hash] = true
	lock.Unlock()
}
