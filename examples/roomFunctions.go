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

func addRoomMember(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddRoomMember(roomId, toUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddRoomMember in sync mode is fine.\n")
			} else {
				fmt.Printf("AddRoomMember in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.AddRoomMember(roomId, toUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddRoomMember in async mode is fine.\n")
					} else {
						fmt.Printf("AddRoomMember in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddRoomMember in async mode error, err: %v\n", err)
			})
	}
}

func deleteRoomMember(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.DelRoomMember(roomId, toUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("DelRoomMember in sync mode is fine.\n")
			} else {
				fmt.Printf("DelRoomMember in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.DelRoomMember(roomId, toUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("DelRoomMember in async mode is fine.\n")
					} else {
						fmt.Printf("DelRoomMember in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("DelRoomMember in async mode error, err: %v\n", err)
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


	addRoomMember(client)
	time.Sleep(500 * time.Millisecond)
	deleteRoomMember(client)


	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}