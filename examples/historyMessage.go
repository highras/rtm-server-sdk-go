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
	mtype int8 = 60
)

//---------------[ Demo ]--------------------//

func printHistory(result *rtm.HistoryMessageResult) {
	fmt.Printf("History count %d, lastId %d, begin %d, end %d\n", result.Num, result.LastId, result.Begin, result.End)
		for _, v := range result.Messages {
			fmt.Printf(" -- id %d, sender %d, mtype %d, mid %d, delete %t, mtime: %d\n", v.Id, v.FromUid, v.MType, v.Mid, v.Deleted, v.MTime)
			fmt.Printf(" -- message: %s\n", v.Message)
			fmt.Printf(" -- attrs: %s\n", v.Attrs)
		}
}

func demoHistoryMessage(client *rtm.RTMServerClient) {

	//-- sync get P2P message
	result, err := client.GetP2PMessage(fromUid, toUid, true, 10, 0, 0, 0)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[P2P History Message] get history in %d with %d in sync mode\n", fromUid, toUid)
				printHistory(result)
			} else {
				fmt.Printf("[P2P History Message] get history in %d with %d in sync mode, err: %v\n", fromUid, toUid, err)
			}
		})

	//-- async get P2P message
	_, err = client.GetP2PMessage(fromUid, toUid, true, 10, 0, 0, 0, func(result *rtm.HistoryMessageResult, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[P2P History Message] get history in %d with %d in async mode\n", fromUid, toUid)
						printHistory(result)
					} else {
						fmt.Printf("[P2P History Message] get history in %d with %d in async mode, error code: %d, error info:%s\n",
							fromUid, toUid, errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("[P2P History Message] get history in %d with %d in async mode, err: %v\n", fromUid, toUid, err)
			})
	}

	//-- sync get group message
	result, err = client.GetGroupMessage(groupId, true, 10, 0, 0, 0)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Group History Message] get group %d history in sync mode\n", groupId)
				printHistory(result)
			} else {
				fmt.Printf("[Group History Message] get group %d history in sync mode, err: %v\n", groupId, err)
			}
		})

	//-- async get group message
	_, err = client.GetGroupMessage(groupId, true, 10, 0, 0, 0, func(result *rtm.HistoryMessageResult, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Group History Message] get group %d history in async mode\n", groupId)
						printHistory(result)
					} else {
						fmt.Printf("[Group History Message] get group %d history in async mode, error code: %d, error info:%s\n",
							groupId, errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Group History Message] get group %d history in async mode, err: %v\n", groupId, err)
			})
	}

	//-- sync get room message
	result, err = client.GetRoomMessage(roomId, true, 10, 0, 0, 0)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Room History Message] get room %d history in sync mode\n", roomId)
				printHistory(result)
			} else {
				fmt.Printf("[Room History Message] get room %d history in sync mode, err: %v\n", roomId, err)
			}
		})

	//-- async get room message
	_, err = client.GetRoomMessage(roomId, true, 10, 0, 0, 0, func(result *rtm.HistoryMessageResult, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Room History Message] get room %d history in async mode\n", roomId)
						printHistory(result)
					} else {
						fmt.Printf("[Room History Message] get room %d history in async mode, error code: %d, error info:%s\n",
							roomId, errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Room History Message] get room %d history in async mode, err: %v\n", roomId, err)
			})
	}

	//-- sync get broadcast message
	result, err = client.GetBroadcastMessage(true, 10, 0, 0, 0)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Broadcast History Message] get broadcast history in sync mode\n")
				printHistory(result)
			} else {
				fmt.Printf("[Broadcast History Message] get broadcast history in sync mode, err: %v\n", err)
			}
		})

	//-- async get broadcast message
	_, err = client.GetBroadcastMessage(true, 10, 0, 0, 0, func(result *rtm.HistoryMessageResult, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Broadcast History Message] get broadcast history in async mode\n")
						printHistory(result)
					} else {
						fmt.Printf("[Broadcast History Message] get broadcast history in async mode, error code: %d, error info:%s\n",
							errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Broadcast History Message] get broadcast history in async mode, err: %v\n", err)
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


	demoHistoryMessage(client)

	locker.print(func(){
			fmt.Println("Wait 1 second for async callbacks are printed.")
		})

	time.Sleep(time.Second)		//-- Waiting for the async callback printed.
}