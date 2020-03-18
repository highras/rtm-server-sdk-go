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
	adminUid int64 = 111
	fromUid int64 = 102456
	toUid int64 = 102457
	toUids []int64 = []int64{102458, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	groupId int64 = 12345
	roomId int64 = 9981
	mtype int8 = 127
)

//---------------[ Demo ]--------------------//

func getToken(client *rtm.RTMServerClient) {

	//-- sync mode
	token, err := client.GetToken(fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetToken in sync mode is fine, token %s\n", token)
			} else {
				fmt.Printf("GetToken in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.GetToken(fromUid, func(token string, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetToken in async mode is fine, token %s\n", token)
					} else {
						fmt.Printf("GetToken in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetToken in async mode error, err: %v\n", err)
			})
	}
}

func getOnlineUsers(client *rtm.RTMServerClient) {

	//-- sync mode
	uids, err := client.GetOnlineUsers(toUids)
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetOnlineUsers in sync mode is fine, online users %v\n", uids)
			} else {
				fmt.Printf("GetOnlineUsers in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.GetOnlineUsers(toUids, func(uids []int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetOnlineUsers in async mode is fine, online users %v\n", uids)
					} else {
						fmt.Printf("GetOnlineUsers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetOnlineUsers in async mode error, err: %v\n", err)
			})
	}
}

func addDevice(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddDevice(fromUid, "app type", "device token")
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddDevice in sync mode is fine.\n")
			} else {
				fmt.Printf("AddDevice in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.AddDevice(fromUid, "app type", "device token", func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddDevice in async mode is fine.\n")
					} else {
						fmt.Printf("AddDevice in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddDevice in async mode error, err: %v\n", err)
			})
	}
}

func removeDevice(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.RemoveDevice(fromUid, "device token")
	locker.print(func(){
			if err == nil {
				fmt.Printf("RemoveDevice in sync mode is fine.\n")
			} else {
				fmt.Printf("RemoveDevice in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.RemoveDevice(fromUid, "device token", func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("RemoveDevice in async mode is fine.\n")
					} else {
						fmt.Printf("RemoveDevice in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("RemoveDevice in async mode error, err: %v\n", err)
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


	getToken(client)
	getOnlineUsers(client)
	addDevice(client)
	removeDevice(client)


	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}