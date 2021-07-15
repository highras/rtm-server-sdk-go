package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"github.com/highras/rtm-server-sdk-go/src/rtm"
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
	fromUid int64   = 111
	uid     int64   = 102457
	toUids  []int64 = []int64{102457, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	roomId  int64   = 999
)

//---------------[ Demo ]--------------------//

func rtcFunctions(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.InviteUserIntoRTCRoom(roomId, fromUid, toUids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("InviteUserIntoRTCRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("InviteUserIntoRTCRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.InviteUserIntoRTCRoom(roomId, fromUid, toUids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("InviteUserIntoRTCRoom in async mode is fine.\n")
			} else {
				fmt.Printf("InviteUserIntoRTCRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("InviteUserIntoRTCRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode force pull into room
	err = client.PullUserIntoRTCRoom(roomId, toUids, 1)
	locker.print(func() {
		if err == nil {
			fmt.Printf("PullUserIntoRTCRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("PullUserIntoRTCRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.PullUserIntoRTCRoom(roomId, toUids, 1, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("PullUserIntoRTCRoom in async mode is fine.\n")
			} else {
				fmt.Printf("PullUserIntoRTCRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("PullUserIntoRTCRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	uids, managers, owner, err1 := client.GetRTCRoomMembers(roomId)
	locker.print(func() {
		if err1 == nil {
			fmt.Printf("GetRTCRoomMembers in sync mode is fine, uids: %v, managers: %v, owner: %d\n", uids, managers, owner)
		} else {
			fmt.Printf("GetRTCRoomMembers in sync mode error, err: %v\n", err1)
		}
	})

	//-- async mode
	_, _, _, err = client.GetRTCRoomMembers(roomId, func(uids []int64, managers []int64, owner int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetRTCRoomMembers in async mode is fine, uids: %v, managers: %v, owner: %d\n", uids, managers, owner)
			} else {
				fmt.Printf("GetRTCRoomMembers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetRTCRoomMembers in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	count, err2 := client.GetRTCRoomMemberCount(roomId)
	locker.print(func() {
		if err2 == nil {
			fmt.Printf("GetRTCRoomMemberCount in sync mode is fine, count: %d\n", count)
		} else {
			fmt.Printf("GetRTCRoomMemberCount in sync mode error, err: %v\n", err2)
		}
	})

	//-- async mode
	_, err = client.GetRTCRoomMemberCount(roomId, func(count int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetRTCRoomMemberCount in async mode is fine, count: %d\n", count)
			} else {
				fmt.Printf("GetRTCRoomMemberCount in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetRTCRoomMemberCount in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	rids, err3 := client.GetRTCRoomList()
	locker.print(func() {
		if err3 == nil {
			fmt.Printf("GetRTCRoomList in sync mode is fine. rids: %v\n", rids)
		} else {
			fmt.Printf("GetRTCRoomList in sync mode error, err: %v\n", err3)
		}
	})

	//-- async mode
	_, err = client.GetRTCRoomList(func(rids []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetRTCRoomList in async mode is fine, rids: %v\n", rids)
			} else {
				fmt.Printf("GetRTCRoomList in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetRTCRoomList in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.SetRTCRoomMicStatus(roomId, true)
	locker.print(func() {
		if err == nil {
			fmt.Printf("SetRTCRoomMicStatus in sync mode is fine.\n")
		} else {
			fmt.Printf("SetRTCRoomMicStatus in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.SetRTCRoomMicStatus(roomId, true, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("SetRTCRoomMicStatus in async mode is fine.\n")
			} else {
				fmt.Printf("SetRTCRoomMicStatus in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("SetRTCRoomMicStatus in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.KickoutFromRTCRoom(roomId, uid, fromUid)
	locker.print(func() {
		if err == nil {
			fmt.Printf("KickoutFromRTCRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("KickoutFromRTCRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.KickoutFromRTCRoom(roomId, uid, fromUid, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("KickoutFromRTCRoom in async mode is fine.\n")
			} else {
				fmt.Printf("KickoutFromRTCRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("KickoutFromRTCRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.CloseRTCRoom(roomId)
	locker.print(func() {
		if err == nil {
			fmt.Printf("CloseRTCRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("CloseRTCRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.CloseRTCRoom(roomId, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("CloseRTCRoom in async mode is fine.\n")
			} else {
				fmt.Printf("CloseRTCRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("CloseRTCRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.AdminCommand(roomId, toUids, 0)
	locker.print(func() {
		if err == nil {
			fmt.Printf("AdminCommand in sync mode is fine.\n")
		} else {
			fmt.Printf("AdminCommand in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.AdminCommand(roomId, toUids, 0, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("AdminCommand in async mode is fine.\n")
			} else {
				fmt.Printf("AdminCommand in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("AdminCommand in async mode error, err: %v\n", err)
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
	
	rtcFunctions(client)
	time.Sleep(500 * time.Millisecond)

	locker.print(func() {
		fmt.Println("Wait 1 second for async callbacks are printed.")
	})

	time.Sleep(time.Second) //-- Waiting for the async callback printed.
}
