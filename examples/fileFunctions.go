package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"runtime"
	"strconv"
	"strings"
	"io/ioutil"
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

func getFileName(filePath string) string {
	pos := strings.LastIndex(filePath, "/")
	if pos != -1 {
		return filePath[(pos+1):]
	} else {
		return filePath
	}
}

func demoSendFiles(client *rtm.RTMServerClient, filePath string) {

	filename := getFileName(filePath)
	if len(filename) == 0 {
		fmt.Println("Invalid file path:", filePath)
		return
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file", filePath, "error:", err)
		return
	}

	fmt.Println("Read file, name:", filename, "data len:", len(data))

	//-- sync send P2P file
	mtime, err := client.SendFile(fromUid, toUid, data, filename)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[P2P File] %d send to %d in sync mode, return mtime: %d\n", fromUid, toUid, mtime)	
			} else {
				fmt.Printf("[P2P File] %d send to %d in sync mode, err: %v\n", fromUid, toUid, err)
			}
		})

	//-- async send P2P file
	_, err = client.SendFile(fromUid, toUid, data, filename, func(mtime int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[P2P File] %d send to %d in async mode, mtime:%d\n", fromUid, toUid, mtime)
					} else {
						fmt.Printf("[P2P File] %d send to %d in async mode, error code: %d, error info:%s\n",
							fromUid, toUid, errorCode, errInfo)
					}
			})
		})
	
	if err != nil {
		locker.print(func(){
				fmt.Printf("[P2P File] %d send to %d in async mode, err: %v\n", fromUid, toUid, err)
			})
	}

	//-- sync send multiple P2P file
	mtime, err = client.SendFiles(fromUid, toUids, data, filename)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Multiple P2P File] %d send to {%v} in sync mode, return mtime: %d\n", fromUid, toUids, mtime)	
			} else {
				fmt.Printf("[Multiple P2P File] %d send to {%v} in sync mode, err: %v\n", fromUid, toUids, err)
			}
		})

	//-- async send multiple P2P file
	_, err = client.SendFiles(fromUid, toUids, data, filename, func(mtime int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Multiple P2P File] %d send to {%v} in async mode, mtime:%d\n", fromUid, toUids, mtime)
					} else {
						fmt.Printf("[Multiple P2P File] %d send to {%v} in async mode, error code: %d, error info:%s\n",
							fromUid, toUids, errorCode, errInfo)
					}
			})
		})
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Multiple P2P File] %d send to {%v} in async mode, err: %v\n", fromUid, toUids, err)
			})
	}

	//-- sync send group file
	mtime, err = client.SendGroupFile(fromUid, groupId, data, filename)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Group File] %d send to group %d in sync mode, return mtime: %d\n", fromUid, groupId, mtime)	
			} else {
				fmt.Printf("[Group File] %d send to group %d in sync mode, err: %v\n", fromUid, groupId, err)
			}
		})

	//-- async send group file
	_, err = client.SendGroupFile(fromUid, groupId, data, filename, func(mtime int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Group File] %d send to group %d in async mode, mtime:%d\n", fromUid, groupId, mtime)
					} else {
						fmt.Printf("[Group File] %d send to group %d in async mode, error code: %d, error info:%s\n",
							fromUid, groupId, errorCode, errInfo)
					}
			})
		})
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Group File] %d send to group %d in async mode, err: %v\n", fromUid, groupId, err)
			})
	}

	//-- sync send room file
	mtime, err = client.SendRoomFile(fromUid, roomId, data, filename)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Room File] %d send to room %d in sync mode, return mtime: %d\n", fromUid, roomId, mtime)	
			} else {
				fmt.Printf("[Room File] %d send to room %d in sync mode, err: %v\n", fromUid, roomId, err)
			}
		})

	//-- async send room file
	_, err = client.SendRoomFile(fromUid, roomId, data, filename, func(mtime int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Room File] %d send to room %d in async mode, mtime:%d\n", fromUid, roomId, mtime)
					} else {
						fmt.Printf("[Room File] %d send to room %d in async mode, error code: %d, error info:%s\n",
							fromUid, roomId, errorCode, errInfo)
					}
			})
		})
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Room File] %d send to room %d in async mode, err: %v\n", fromUid, roomId, err)
			})
	}

	//-- sync send boardcast file
	mtime, err = client.SendBroadcastFile(adminUid, data, filename)
	locker.print(func(){
			if err == nil {
				fmt.Printf("[Boardcast File] %d send boardcast message in sync mode, return mtime: %d\n", adminUid, mtime)	
			} else {
				fmt.Printf("[Boardcast File] %d send boardcast message in sync mode, err: %v\n", adminUid, err)
			}
		})

	//-- async send boardcast file
	_, err = client.SendBroadcastFile(adminUid, data, filename, func(mtime int64, errorCode int, errInfo string){
		locker.print(func(){
				if errorCode == fpnn.FPNN_EC_OK {
						fmt.Printf("[Boardcast File] %d send boardcast message in async mode, mtime:%d\n", adminUid, mtime)
					} else {
						fmt.Printf("[Boardcast File] %d send boardcast message in async mode, error code: %d, error info:%s\n",
							adminUid, errorCode, errInfo)
					}
			})
		})
	if err != nil {
		locker.print(func(){
				fmt.Printf("[Boardcast File] %d send boardcast message in async mode, err: %v\n", adminUid, err)
			})
	}
}

func main() {

	if len(os.Args) != 5 {
		fmt.Println("Usage:", os.Args[0], "<endpoint>", "<pid>", "<secretKey>", "<filePath>")
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	pid, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Pid is invalid. Error:", err)
		return
	}
	client := rtm.NewRTMServerClient(int32(pid), os.Args[3], os.Args[1])


	demoSendFiles(client, os.Args[4])

	locker.print(func(){
			fmt.Println("Wait 3 second for async callbacks are printed.")
		})

	time.Sleep(3 * time.Second)		//-- Waiting for the async callback printed.
}