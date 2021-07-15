package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"runtime"
	"strconv"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"github.com/highras/rtm-server-sdk-go/src/rtm"
)

//---------------[ Help tools for serializing concurrent printing. ]---------------------//
type PrintLocker struct {
	mutex	sync.Mutex
}

func (locker *PrintLocker) print(proc func()) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	proc()
}

var locker PrintLocker = PrintLocker{}

var (
	uid int64 = 102456
	demoKey string = "demo key"
	demoKey2 string = "demo key 2"
)

//---------------[ Demo ]--------------------//

func demoSetData(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.SetData(uid, demoKey, "123 456 789")
	locker.print(func(){
			if err == nil {
				fmt.Printf("SetData in sync mode is fine.\n")
			} else {
				fmt.Printf("SetData in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.SetData(uid, demoKey2, "abc def ghi", func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("SetData in async mode is fine.\n")
					} else {
						fmt.Printf("SetData in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("SetData in async mode error, err: %v\n", err)
			})
	}
}


func demoGetData(client *rtm.RTMServerClient) {

	//-- sync mode
	data, err := client.GetData(uid, demoKey)
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetData in sync mode is fine, data: %s\n", data)
			} else {
				fmt.Printf("GetData in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.GetData(uid, demoKey2, func(text string, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetData in async mode is fine, data: %s\n", data)
					} else {
						fmt.Printf("GetData in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetData in async mode error, err: %v\n", err)
			})
	}
}

func demoDelData(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.DelData(uid, demoKey)
	locker.print(func(){
			if err == nil {
				fmt.Printf("DelData in sync mode is fine.\n")
			} else {
				fmt.Printf("DelData in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.DelData(uid, demoKey2, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("DelData in async mode is fine.\n")
					} else {
						fmt.Printf("DelData in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("DelData in async mode error, err: %v\n", err)
			})
	}
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "<endpoint>", "<pid>", "<secretKey>")
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	pid, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Pid is invalid. Error:", err)
		return
	}
	client := rtm.NewRTMServerClient(int32(pid), os.Args[3], os.Args[1])

	client.SetKeepAlive(true)
	
	demoSetData(client)
	demoGetData(client)
	time.Sleep(time.Second)
	demoDelData(client)
	time.Sleep(time.Second)	
	demoGetData(client)


	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}