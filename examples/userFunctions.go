package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"runtime"
	"strconv"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"../src/rtm"
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
	uid2 int64 = 102457
)

//---------------[ Demo ]--------------------//

func demoSetUserInfo(client *rtm.RTMServerClient) {

	publicInfo := "user piublic info"
	privateInfo := "user private info"

	//-- sync mode
	err := client.SetUserInfo(uid, &publicInfo, &privateInfo)
	locker.print(func(){
			if err == nil {
				fmt.Printf("SetUserInfo in sync mode is fine.\n")
			} else {
				fmt.Printf("SetUserInfo in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.SetUserInfo(uid2, &publicInfo, &privateInfo, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("SetUserInfo in async mode is fine.\n")
					} else {
						fmt.Printf("SetUserInfo in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("SetUserInfo in async mode error, err: %v\n", err)
			})
	}
}

func demoGetUserInfo(client *rtm.RTMServerClient) {

	//-- sync mode
	publicInfo, privateInfo, err := client.GetUserInfo(uid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetUserInfo in sync mode is fine, public info: %s, private info: %s\n", publicInfo, privateInfo)
			} else {
				fmt.Printf("GetUserInfo in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, _, err = client.GetUserInfo(uid2, func(publicInfos string, privateInfos string, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetUserInfo in async mode is fine, public info: %s, private info: %s\n", publicInfos, privateInfos)
					} else {
						fmt.Printf("GetUserInfo in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetUserInfo in async mode error, err: %v\n", err)
			})
	}
}

func demoGetUserPublicInfo(client *rtm.RTMServerClient) {

	//-- sync mode
	infos, err := client.GetUserPublicInfo([]int64{uid, uid2})
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetUserPublicInfo in sync mode is fine\n")
				for k, v := range infos {
					fmt.Printf("  user %s info %s\n", k, v)
				}
			} else {
				fmt.Printf("GetUserPublicInfo in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.GetUserPublicInfo([]int64{uid, uid2}, func(infos map[string]string, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetUserPublicInfo in async mode is fine\n")
						for k, v := range infos {
							fmt.Printf("  user %s info %s\n", k, v)
						}
					} else {
						fmt.Printf("GetUserPublicInfo in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetUserPublicInfo in async mode error, err: %v\n", err)
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
	
	demoSetUserInfo(client)
	demoGetUserInfo(client)
	demoGetUserPublicInfo(client)

	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}