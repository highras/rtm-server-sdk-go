package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/highras/rtm-server-sdk-go/src/rtm"
)

//--------------------[Help tools for serializing concurrent print.]--------------
type printLocker struct {
	mutex sync.Mutex
}

func (locker *printLocker) print(proc func()) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	proc()
}

var locker = printLocker{}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "<endpoint>", "<pid>", "<secretKey>")
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	pid, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("pid is invalid. Error:", err)
		return
	}

	client := rtm.NewRTMServerClient(int32(pid), os.Args[3], os.Args[1])
	client.SetKeepAlive(true)
	client.SetOnConnectedCallback(func(connId uint64, endpoint string, connected bool, autoReconnect bool, connectState *rtm.RtmRegressiveState) {
		locker.print(func() {
			if connected {
				fmt.Printf("connect success connId:= %d, endpoint:= %s.\n", connId, endpoint)
			} else {
				info := fmt.Sprintf("RTM last connected time at %d, currentFailedCount = %d", connectState.ConnectSuccessMilliseconds, connectState.CurrentFailedCount)
				fmt.Printf("connect result connId:= %d, endpoint:= %s, connected:= %t, autoReconnect:= %t, reconnectinfo = %s.\n", connId, endpoint, connected, autoReconnect, info)
			}
		})
	})

	client.SetOnClosedCallback(func(connId uint64, endpoint string, autoReconnect bool, connectState *rtm.RtmRegressiveState) {
		locker.print(func() {
			if connectState != nil {
				info := fmt.Sprintf("RTM last connected time at %d, currentFailedCount = %d", connectState.ConnectSuccessMilliseconds, connectState.CurrentFailedCount)
				fmt.Printf("connect close connId:= %d, endpoint:= %s, autoReconnect:= %t, reconnectinfo:= %s.\n", connId, endpoint, autoReconnect, info)
			}
		})
	})

	if ok := client.Connect(); !ok {
		fmt.Println("Connect to", os.Args[1], "failed")
	}
	locker.print(func() {
		fmt.Println("Wait 500 second for async callback all print")
	})
	time.Sleep(500 * time.Second)
}
