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
	adminUid int64   = 111
	fromUid  int64   = 102456
	toUid    int64   = 102457
	toUids   []int64 = []int64{102458, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	groupId  int64   = 12345
	roomId   int64   = 9981
	mtype    int8    = 127
)

//---------------[ Demo ]--------------------//

func addGroupMembers(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.AddGroupMembers(groupId, toUids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("AddGroupMembers in sync mode is fine.\n")
		} else {
			fmt.Printf("AddGroupMembers in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.AddGroupMembers(groupId, toUids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("AddGroupMembers in async mode is fine.\n")
			} else {
				fmt.Printf("AddGroupMembers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("AddGroupMembers in async mode error, err: %v\n", err)
		})
	}
}

func deleteGroupMembers(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.DelGroupMembers(groupId, toUids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("DelGroupMembers in sync mode is fine.\n")
		} else {
			fmt.Printf("DelGroupMembers in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.DelGroupMembers(groupId, toUids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("DelGroupMembers in async mode is fine.\n")
			} else {
				fmt.Printf("DelGroupMembers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("DelGroupMembers in async mode error, err: %v\n", err)
		})
	}
}

func deleteGroup(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.DelGroup(groupId)
	locker.print(func() {
		if err == nil {
			fmt.Printf("DelGroup in sync mode is fine.\n")
		} else {
			fmt.Printf("DelGroup in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.DelGroup(groupId, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("DelGroup in async mode is fine.\n")
			} else {
				fmt.Printf("DelGroup in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("DelGroup in async mode error, err: %v\n", err)
		})
	}
}

func getGroupMembers(client *rtm.RTMServerClient) {

	//-- sync mode
	uids, err := client.GetGroupMembers(groupId)
	locker.print(func() {
		if err == nil {
			fmt.Printf("GetGroupMembers in sync mode is fine, members: %v\n", uids)
		} else {
			fmt.Printf("GetGroupMembers in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	_, err = client.GetGroupMembers(groupId, func(uids []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetGroupMembers in async mode is fine, members: %v\n", uids)
			} else {
				fmt.Printf("GetGroupMembers in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetGroupMembers in async mode error, err: %v\n", err)
		})
	}
}

func isGroupMember(client *rtm.RTMServerClient) {

	//-- sync mode
	member, err := client.IsGroupMember(groupId, toUids[1])
	locker.print(func() {
		if err == nil {
			fmt.Printf("IsGroupMember in sync mode is fine, member: %t\n", member)
		} else {
			fmt.Printf("IsGroupMember in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	_, err = client.IsGroupMember(groupId, toUids[1], func(member bool, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("IsGroupMember in async mode is fine, member: %t\n", member)
			} else {
				fmt.Printf("IsGroupMember in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("IsGroupMember in async mode error, err: %v\n", err)
		})
	}
}

func getUserGroups(client *rtm.RTMServerClient) {

	//-- sync mode
	groups, err := client.GetUserGroups(toUids[1])
	locker.print(func() {
		if err == nil {
			fmt.Printf("GetUserGroups in sync mode is fine, groups: %v\n", groups)
		} else {
			fmt.Printf("GetUserGroups in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	_, err = client.GetUserGroups(toUids[1], func(groups []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetUserGroups in async mode is fine, groups: %v\n", groups)
			} else {
				fmt.Printf("GetUserGroups in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("GetUserGroups in async mode error, err: %v\n", err)
		})
	}
}

func clearProjectGroup(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.ClearProjectGroup()
	locker.print(func() {
		if err == nil {
			fmt.Printf("ClearProjectGroup in sync mode is fine\n")
		} else {
			fmt.Printf("ClearProjectGroup in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.ClearProjectGroup(func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("ClearProjectGroup in async mode is fine\n")
			} else {
				fmt.Printf("ClearProjectGroup in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("ClearProjectGroup in async mode error, err: %v\n", err)
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

	addGroupMembers(client)
	isGroupMember(client)
	getGroupMembers(client)
	cnt, err := client.GetGroupCount(groupId)
	fmt.Println("count: ", cnt)
	getUserGroups(client)

	// deleteGroupMembers(client)
	// time.Sleep(500 * time.Millisecond)
	// getGroupMembers(client)

	// addGroupMembers(client)
	// time.Sleep(500 * time.Millisecond)
	// deleteGroup(client)
	// time.Sleep(500 * time.Millisecond)
	// getGroupMembers(client)

	// addGroupMembers(client)
	// time.Sleep(500 * time.Millisecond)
	// clearProjectGroup(client)
	// time.Sleep(5000 * time.Millisecond)
	// getGroupMembers(client)

	// err = client.EnterUniqueGroup(1000, 1000)
	// if err != nil {
	// 	fmt.Errorf("error enteruniquegroup.")
	// }
	// gids, err := client.GetUserGroups(1000)
	// if err == nil {
	// 	fmt.Println(gids)
	// }
	// err = client.EnterUniqueGroup(1001, 1000)
	// if err != nil {
	// 	fmt.Errorf("error enteruniquegroup.")
	// }
	// gids, err = client.GetUserGroups(1000)
	// if err == nil {
	// 	fmt.Println(gids)
	// }

	locker.print(func() {
		fmt.Println("Wait 1 second for async callbacks are printed.")
	})

	time.Sleep(time.Second) //-- Waiting for the async callback printed.
}
