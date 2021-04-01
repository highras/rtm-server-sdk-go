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

func voiceFunctions(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.InviteUserIntoVoiceRoom(roomId, fromUid, toUids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("InviteUserIntoVoiceRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("InviteUserIntoVoiceRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.InviteUserIntoVoiceRoom(roomId, fromUid, toUids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("InviteUserIntoVoiceRoom in async mode is fine.\n")
			} else {
				fmt.Printf("InviteUserIntoVoiceRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("InviteUserIntoVoiceRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode force pull into room
	err = client.PullUserIntoVoiceRoom(roomId, toUids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("PullUserIntoVoiceRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("PullUserIntoVoiceRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.PullUserIntoVoiceRoom(roomId, toUids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("PullUserIntoVoiceRoom in async mode is fine.\n")
			} else {
				fmt.Printf("PullUserIntoVoiceRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("PullUserIntoVoiceRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	uids, managers, err1 := client.GetVoiceRoomMembers(roomId)
	locker.print(func() {
		if err1 == nil {
			fmt.Printf("GetVoiceRoomMembers in sync mode is fine, uids: %v, managers: %v\n", uids, managers)
		} else {
			fmt.Printf("GetVoiceRoomMembers in sync mode error, err: %v\n", err1)
		}
	})

	//-- async mode
	_, _, err = client.GetVoiceRoomMembers(roomId, func(uids []int64, managers []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetVoiceRoomMembers in async mode is fine, uids: %v, managers: %v\n", uids, managers)
			} else {
				fmt.Printf("GetVoiceRoomMembers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetVoiceRoomMembers in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	count, err2 := client.GetVoiceRoomMemberCount(roomId)
	locker.print(func() {
		if err2 == nil {
			fmt.Printf("GetVoiceRoomMemberCount in sync mode is fine, count: %d\n", count)
		} else {
			fmt.Printf("GetVoiceRoomMemberCount in sync mode error, err: %v\n", err2)
		}
	})

	//-- async mode
	_, err = client.GetVoiceRoomMemberCount(roomId, func(count int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetVoiceRoomMemberCount in async mode is fine, count: %d\n", count)
			} else {
				fmt.Printf("GetVoiceRoomMemberCount in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetVoiceRoomMemberCount in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	rids, err3 := client.GetVoiceRoomList()
	locker.print(func() {
		if err3 == nil {
			fmt.Printf("GetVoiceRoomList in sync mode is fine. rids: %v\n", rids)
		} else {
			fmt.Printf("GetVoiceRoomList in sync mode error, err: %v\n", err3)
		}
	})

	//-- async mode
	_, err = client.GetVoiceRoomList(func(rids []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetVoiceRoomList in async mode is fine, rids: %v\n", rids)
			} else {
				fmt.Printf("GetVoiceRoomList in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetVoiceRoomList in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.SetVoiceRoomMicStatus(roomId, true)
	locker.print(func() {
		if err == nil {
			fmt.Printf("SetVoiceRoomMicStatus in sync mode is fine.\n")
		} else {
			fmt.Printf("SetVoiceRoomMicStatus in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.SetVoiceRoomMicStatus(roomId, true, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("SetVoiceRoomMicStatus in async mode is fine.\n")
			} else {
				fmt.Printf("SetVoiceRoomMicStatus in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("SetVoiceRoomMicStatus in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.KickoutFromVoiceRoom(roomId, uid, fromUid)
	locker.print(func() {
		if err == nil {
			fmt.Printf("KickoutFromVoiceRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("KickoutFromVoiceRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.KickoutFromVoiceRoom(roomId, uid, fromUid, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("KickoutFromVoiceRoom in async mode is fine.\n")
			} else {
				fmt.Printf("KickoutFromVoiceRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("KickoutFromVoiceRoom in async mode error, err: %v\n", err)
		})
	}

	//-- sync mode
	err = client.CloseVoiceRoom(roomId)
	locker.print(func() {
		if err == nil {
			fmt.Printf("CloseVoiceRoom in sync mode is fine.\n")
		} else {
			fmt.Printf("CloseVoiceRoom in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.CloseVoiceRoom(roomId, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("CloseVoiceRoom in async mode is fine.\n")
			} else {
				fmt.Printf("CloseVoiceRoom in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("CloseVoiceRoom in async mode error, err: %v\n", err)
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

	voiceFunctions(client)
	time.Sleep(500 * time.Millisecond)

	locker.print(func() {
		fmt.Println("Wait 1 second for async callbacks are printed.")
	})

	time.Sleep(time.Second) //-- Waiting for the async callback printed.
}
