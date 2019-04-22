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
	adminUid int64 = 111
	fromUid int64 = 102456
	toUid int64 = 102457
	toUids []int64 = []int64{102458, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	groupId int64 = 12345
	roomId int64 = 9981
	mtype int8 = 127
)

//---------------[ Demo ]--------------------//

func addFriends(client *rtm.RTMServerClient) {

	//-- sync add friends
	err := client.AddFriends(fromUid, toUids)
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddFriends in sync mode is fine.\n")
			} else {
				fmt.Printf("AddFriends in sync mode error, err: %v\n", err)
			}
		})

	//-- async add friends
	err = client.AddFriends(fromUid, toUids, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddFriends in async mode is fine.\n")
					} else {
						fmt.Printf("AddFriends in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddFriends in async mode error, err: %v\n", err)
			})
	}
}

func deleteFriends(client *rtm.RTMServerClient) {

	//-- sync delete friends
	err := client.DelFriends(fromUid, toUids)
	locker.print(func(){
			if err == nil {
				fmt.Printf("DelFriends in sync mode is fine.\n")
			} else {
				fmt.Printf("DelFriends in sync mode error, err: %v\n", err)
			}
		})

	//-- async delete friends
	err = client.DelFriends(fromUid, toUids, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("DelFriends in async mode is fine.\n")
					} else {
						fmt.Printf("DelFriends in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("DelFriends in async mode error, err: %v\n", err)
			})
	}

}

func getFriends(client *rtm.RTMServerClient) {

	//-- sync get friends
	friends, err := client.GetFriends(fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("GetFriends in sync mode is fine, friends %v\n", friends)
			} else {
				fmt.Printf("GetFriends in sync mode error, err: %v\n", err)
			}
		})

	//-- async get friends
	_, err = client.GetFriends(fromUid, func(friends []int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("GetFriends in async mode is fine, friends %v\n", friends)
					} else {
						fmt.Printf("GetFriends in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("GetFriends in async mode error, err: %v\n", err)
			})
	}
}

func isFriend(client *rtm.RTMServerClient) {
	
	//-- sync is friend
	friend, err := client.IsFriend(fromUid, toUids[1])
	locker.print(func(){
			if err == nil {
				fmt.Printf("IsFriend in sync mode is fine, friend %t\n", friend)
			} else {
				fmt.Printf("IsFriend in sync mode error, err: %v\n", err)
			}
		})

	//-- async is friend
	_, err = client.IsFriend(fromUid, toUids[1], func(friend bool, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("IsFriend in async mode is fine, friend %t\n", friend)
					} else {
						fmt.Printf("IsFriend in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("IsFriend in async mode error, err: %v\n", err)
			})
	}
}

func isFriends(client *rtm.RTMServerClient) {
	
	//-- sync get friends
	friends, err := client.IsFriends(fromUid, toUids)
	locker.print(func(){
			if err == nil {
				fmt.Printf("IsFriends in sync mode is fine, friends %v\n", friends)
			} else {
				fmt.Printf("IsFriends in sync mode error, err: %v\n", err)
			}
		})

	//-- async get friends
	_, err = client.IsFriends(fromUid, toUids, func(friends []int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("IsFriends in async mode is fine, friends %v\n", friends)
					} else {
						fmt.Printf("IsFriends in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("IsFriends in async mode error, err: %v\n", err)
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


	addFriends(client)
	getFriends(client)
	isFriend(client)
	isFriends(client)
	time.Sleep(500 * time.Millisecond)
	deleteFriends(client)
	time.Sleep(500 * time.Millisecond)
	getFriends(client)
	isFriend(client)

	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}