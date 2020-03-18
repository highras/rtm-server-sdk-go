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

func addGroupBan(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddGroupBan(groupId, fromUid, 300)
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddGroupBan in sync mode is fine.\n")
			} else {
				fmt.Printf("AddGroupBan in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.AddGroupBan(groupId, fromUid, 300, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddGroupBan in async mode is fine.\n")
					} else {
						fmt.Printf("AddGroupBan in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddGroupBan in async mode error, err: %v\n", err)
			})
	}
}

func removeGroupBan(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.RemoveGroupBan(groupId, fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("RemoveGroupBan in sync mode is fine.\n")
			} else {
				fmt.Printf("RemoveGroupBan in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.RemoveGroupBan(groupId, fromUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("RemoveGroupBan in async mode is fine.\n")
					} else {
						fmt.Printf("RemoveGroupBan in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("RemoveGroupBan in async mode error, err: %v\n", err)
			})
	}
}

func addRoomBan(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddRoomBan(roomId, fromUid, 300)
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddRoomBan in sync mode is fine.\n")
			} else {
				fmt.Printf("AddRoomBan in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.AddRoomBan(roomId, fromUid, 300, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddRoomBan in async mode is fine.\n")
					} else {
						fmt.Printf("AddRoomBan in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddRoomBan in async mode error, err: %v\n", err)
			})
	}
}

func removeRoomBan(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.RemoveRoomBan(roomId, fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("RemoveRoomBan in sync mode is fine.\n")
			} else {
				fmt.Printf("RemoveRoomBan in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.RemoveRoomBan(roomId, fromUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("RemoveRoomBan in async mode is fine.\n")
					} else {
						fmt.Printf("RemoveRoomBan in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("RemoveRoomBan in async mode error, err: %v\n", err)
			})
	}
}

func addProjectBlack(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddProjectBlack(fromUid, 300)
	locker.print(func(){
			if err == nil {
				fmt.Printf("AddProjectBlack in sync mode is fine.\n")
			} else {
				fmt.Printf("AddProjectBlack in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.AddProjectBlack(fromUid, 300, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("AddProjectBlack in async mode is fine.\n")
					} else {
						fmt.Printf("AddProjectBlack in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("AddProjectBlack in async mode error, err: %v\n", err)
			})
	}
}

func removeProjectBlack(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.RemoveProjectBlack(fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("RemoveProjectBlack in sync mode is fine.\n")
			} else {
				fmt.Printf("RemoveProjectBlack in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.RemoveProjectBlack(fromUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("RemoveProjectBlack in async mode is fine.\n")
					} else {
						fmt.Printf("RemoveProjectBlack in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("RemoveProjectBlack in async mode error, err: %v\n", err)
			})
	}
}

func isBanOfGroup(client *rtm.RTMServerClient) {

	//-- sync mode
	status, err := client.IsBanOfGroup(groupId, fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("IsBanOfGroup in sync mode is fine, status %t\n", status)
			} else {
				fmt.Printf("IsBanOfGroup in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.IsBanOfGroup(groupId, fromUid, func(status bool, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("IsBanOfGroup in async mode is fine, status %t\n", status)
					} else {
						fmt.Printf("IsBanOfGroup in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("IsBanOfGroup in async mode error, err: %v\n", err)
			})
	}
}

func isBanOfRoom(client *rtm.RTMServerClient) {

	//-- sync mode
	status, err := client.IsBanOfRoom(roomId, fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("IsBanOfRoom in sync mode is fine, status %t\n", status)
			} else {
				fmt.Printf("IsBanOfRoom in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.IsBanOfRoom(roomId, fromUid, func(status bool, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("IsBanOfRoom in async mode is fine, status %t\n", status)
					} else {
						fmt.Printf("IsBanOfRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("IsBanOfRoom in async mode error, err: %v\n", err)
			})
	}
}

func isProjectBlack(client *rtm.RTMServerClient) {

	//-- sync mode
	status, err := client.IsProjectBlack(fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("IsProjectBlack in sync mode is fine, status %t\n", status)
			} else {
				fmt.Printf("IsProjectBlack in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	_, err = client.IsProjectBlack(fromUid, func(status bool, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("IsProjectBlack in async mode is fine, status %t\n", status)
					} else {
						fmt.Printf("IsProjectBlack in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("IsProjectBlack in async mode error, err: %v\n", err)
			})
	}
}

func kickout(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.Kickout(fromUid)
	locker.print(func(){
			if err == nil {
				fmt.Printf("Kickout in sync mode is fine.\n")
			} else {
				fmt.Printf("Kickout in sync mode error, err: %v\n", err)
			}
		})

	//-- async mode
	err = client.Kickout(fromUid, func(errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("Kickout in async mode is fine.\n")
					} else {
						fmt.Printf("Kickout in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("Kickout in async mode error, err: %v\n", err)
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

	addGroupBan(client)
	time.Sleep(500 * time.Millisecond)
	isBanOfGroup(client)
	time.Sleep(500 * time.Millisecond)
	removeGroupBan(client)
	time.Sleep(500 * time.Millisecond)

	addRoomBan(client)
	time.Sleep(500 * time.Millisecond)
	isBanOfRoom(client)
	time.Sleep(500 * time.Millisecond)
	removeRoomBan(client)
	time.Sleep(500 * time.Millisecond)

	addProjectBlack(client)
	time.Sleep(500 * time.Millisecond)
	isProjectBlack(client)
	time.Sleep(500 * time.Millisecond)
	removeProjectBlack(client)
	time.Sleep(500 * time.Millisecond)
	
	kickout(client)

	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}