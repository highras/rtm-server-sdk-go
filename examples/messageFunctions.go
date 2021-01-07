package main

import (
	"fmt"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"github.com/highras/rtm-server-sdk-go/src/rtm"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//---------------[ Help tools for serializing concurrent printing. ]---------------------//
type PrintLocker struct {
	mutex sync.Mutex
}

func (locker *PrintLocker) print(proc func()) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	proc()
}

var locker PrintLocker = PrintLocker{}

var (
	adminUid int64   = 111
	fromUid  int64   = 102456
	toUid    int64   = 102457
	toUids   []int64 = []int64{102458, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	groupId  int64   = 12345
	roomId   int64   = 9981
	mtype    int8    = 127
)

//---------------[ Demo ]--------------------//

func demoSendMessage(client *rtm.RTMServerClient) {

	//-- sync send P2P message
	mtime, err := client.SendMessage(fromUid, toUid, mtype, "test sync P2P message")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[P2P Message] %d send to %d in sync mode, return mtime: %d\n", fromUid, toUid, mtime)
		} else {
			fmt.Printf("[P2P Message] %d send to %d in sync mode, err: %v\n", fromUid, toUid, err)
		}
	})

	//-- async send P2P message
	_, err = client.SendMessage(fromUid, toUid, mtype, "test async P2P message", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[P2P Message] %d send to %d in async mode, mtime:%d\n", fromUid, toUid, mtime)
			} else {
				fmt.Printf("[P2P Message] %d send to %d in async mode, error code: %d, error info:%s\n",
					fromUid, toUid, errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[P2P Message] %d send to %d in async mode, err: %v\n", fromUid, toUid, err)
		})
	}

	//-- sync send multiple P2P message
	mtime, err = client.SendMessages(fromUid, toUids, mtype, "test sync multiple P2P message")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Multiple P2P Message] %d send to {%v} in sync mode, return mtime: %d\n", fromUid, toUids, mtime)
		} else {
			fmt.Printf("[Multiple P2P Message] %d send to {%v} in sync mode, err: %v\n", fromUid, toUids, err)
		}
	})

	//-- async send multiple P2P message
	_, err = client.SendMessages(fromUid, toUids, mtype, "test async multiple P2P message", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Multiple P2P Message] %d send to {%v} in async mode, mtime:%d\n", fromUid, toUids, mtime)
			} else {
				fmt.Printf("[Multiple P2P Message] %d send to {%v} in async mode, error code: %d, error info:%s\n",
					fromUid, toUids, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Multiple P2P Message] %d send to {%v} in async mode, err: %v\n", fromUid, toUids, err)
		})
	}

	_ = client.AddGroupMembers(groupId, []int64{fromUid})
	//-- sync send group message
	mtime, err = client.SendGroupMessage(fromUid, groupId, mtype, "test sync group message")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Group Message] %d send to group %d in sync mode, return mtime: %d\n", fromUid, groupId, mtime)
		} else {
			fmt.Printf("[Group Message] %d send to group %d in sync mode, err: %v\n", fromUid, groupId, err)
		}
	})

	//-- async send group message
	_, err = client.SendGroupMessage(fromUid, groupId, mtype, "test async group message", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Group Message] %d send to group %d in async mode, mtime:%d\n", fromUid, groupId, mtime)
			} else {
				fmt.Printf("[Group Message] %d send to group %d in async mode, error code: %d, error info:%s\n",
					fromUid, groupId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Group Message] %d send to group %d in async mode, err: %v\n", fromUid, groupId, err)
		})
	}

	//-- sync send room message
	mtime, err = client.SendRoomMessage(fromUid, roomId, mtype, "test sync room message")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Room Message] %d send to room %d in sync mode, return mtime: %d\n", fromUid, roomId, mtime)
		} else {
			fmt.Printf("[Room Message] %d send to room %d in sync mode, err: %v\n", fromUid, roomId, err)
		}
	})

	//-- async send room message
	_, err = client.SendRoomMessage(fromUid, roomId, mtype, "test async room message", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Room Message] %d send to room %d in async mode, mtime:%d\n", fromUid, roomId, mtime)
			} else {
				fmt.Printf("[Room Message] %d send to room %d in async mode, error code: %d, error info:%s\n",
					fromUid, roomId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Room Message] %d send to room %d in async mode, err: %v\n", fromUid, roomId, err)
		})
	}

	//-- sync send boardcast message
	mtime, err = client.SendBroadcastMessage(adminUid, mtype, "test sync boardcast message")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Boardcast Message] %d send boardcast message in sync mode, return mtime: %d\n", adminUid, mtime)
		} else {
			fmt.Printf("[Boardcast Message] %d send boardcast message in sync mode, err: %v\n", adminUid, err)
		}
	})

	//-- async send boardcast message
	_, err = client.SendBroadcastMessage(adminUid, mtype, "test async boardcast message", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Boardcast Message] %d send boardcast message in async mode, mtime:%d\n", adminUid, mtime)
			} else {
				fmt.Printf("[Boardcast Message] %d send boardcast message in async mode, error code: %d, error info:%s\n",
					adminUid, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Boardcast Message] %d send boardcast message in async mode, err: %v\n", adminUid, err)
		})
	}
}

func getMsgCount(client *rtm.RTMServerClient) {
	// sync
	sender, count, err := client.GetMsgCount(rtm.MessageType_Group, groupId, 0, 0, nil)
	locker.print(func() {
		if err == nil {
			fmt.Printf("[GetMsgCount Group] in sync mode, return sender: %d, count: %d\n", sender, count)
		} else {
			fmt.Printf("[GetMsgCount Group] in sync mode, err: %v\n", err)
		}
	})

	//-- async
	_, _, err = client.GetMsgCount(rtm.MessageType_Group, groupId, 0, 0, nil, func(sender int32, count int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[GetMsgCount Group] in async mode, return sender: %d, count: %d\n", sender, count)
			} else {
				fmt.Printf("[GetMsgCount Group] in async mode, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[GetMsgCount Group] in async mode, err: %v\n", err)
		})
	}

	// sync
	sender, count, err = client.GetMsgCount(rtm.MessageType_Room, roomId, 0, 0, nil)
	locker.print(func() {
		if err == nil {
			fmt.Printf("[GetMsgCount Room] in sync mode, return sender: %d, count: %d\n", sender, count)
		} else {
			fmt.Printf("[GetMsgCount Room] in sync mode, err: %v\n", err)
		}
	})

	//-- async
	_, _, err = client.GetMsgCount(rtm.MessageType_Room, roomId, 0, 0, nil, func(sender int32, count int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[GetMsgCount Room] in async mode, return sender: %d, count: %d\n", sender, count)
			} else {
				fmt.Printf("[GetMsgCount Room] in async mode, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[GetMsgCount Room] in async mode, err: %v\n", err)
		})
	}
}

func deleteMessage(client *rtm.RTMServerClient) {

	var mid int64 = 123456

	//-- sync mode
	err := client.DelMessage(mid, fromUid, toUid, rtm.MessageType_P2P)
	locker.print(func() {
		if err == nil {
			fmt.Printf("DelMessage in sync mode is fine.\n")
		} else {
			fmt.Printf("DelMessage in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.DelMessage(mid, fromUid, groupId, rtm.MessageType_Group, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("DelMessage in async mode is fine.\n")
			} else {
				fmt.Printf("DelMessage in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("DelMessage in async mode error, err: %v\n", err)
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

	demoSendMessage(client)
	//deleteMessage(client)
	time.Sleep(2 * time.Second)
	getMsgCount(client)
	locker.print(func() {
		fmt.Println("Wait 1 second for async callbacks are printed.")
	})

	time.Sleep(time.Second) //-- Waiting for the async callback printed.
}
