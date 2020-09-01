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

//--------------------[Help tools for serializing concurrent print.]--------------
type printLocker struct {
	mutex sync.Mutex
}

func (locker *printLocker) print(proc func()) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	proc()
}

var locker = printLocker{}

var (
	uid   int64 = 1111
	buid  int64 = 2222
	buids       = []int64{3333, 4444, 5555, 6666}
)

//--------------[Demo]---------------------------------------

func addBlacks(client *rtm.RTMServerClient) {
	// sync method
	err := client.AddBlacks(uid, buids)
	locker.print(func() {
		if err == nil {
			fmt.Println("AddBlacks in sync mode is ok.")
		} else {
			fmt.Printf("AddBlacks in sync mode error, err:%v\n", err)
		}
	})

	// async method
	err = client.AddBlacks(uid, buids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Println("AddBlacks in async mode is ok.")
			} else {
				fmt.Printf("AddBlacks in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("AddBlacks in async mode error, err: %v.\n", err)
		})
	}

	// test send chat
	mtime, err1 := client.SendChat(uid, buid, "test sync chat message")
	locker.print(func() {
		if err1 == nil {
			fmt.Printf("[P2P Chat] %v send to %v in sync mode, return mtime: %v.\n", uid, buid, mtime)
		} else {
			fmt.Printf("[P2P Chat] %v send to %v in sync mode, err: %v.\n", uid, buid, err1)
		}
	})

	mtime, err1 = client.SendChat(buid, uid, "test sync chat message")
	locker.print(func() {
		if err1 == nil {
			fmt.Printf("[P2P Chat] %v send to %v in sync mode, return mtime: %v.\n", buid, uid, mtime)
		} else {
			fmt.Printf("[P2P Chat] %v send to %v in sync mode, err: %v.\n", buid, uid, err1)
		}
	})
	time.Sleep(400 * time.Millisecond)

}

func getBlacks(client *rtm.RTMServerClient) {
	// sync method
	ids, err := client.GetBlacks(uid)
	locker.print(func() {
		if err == nil {
			fmt.Printf("GetBlacks in sync mode is fine, uids: %v.\n", ids)
		} else {
			fmt.Printf("GetBlacks in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	_, err = client.GetBlacks(uid, func(ids []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("GetBlacks in async mode is fine, uids: %v.\n", ids)
			} else {
				fmt.Printf("GetBlacks in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("GetBlacks in async mode error, err: %v.\n", err)
		})
	}
}

func isBlacks(client *rtm.RTMServerClient) {
	// sync method
	ids, err := client.IsBlacks(uid, buids)
	locker.print(func() {
		if err == nil {
			fmt.Printf("IsBlacks in sync mode is fine, buids: %v.\n", ids)
		} else {
			fmt.Printf("IsBlacks in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	_, err = client.IsBlacks(uid, buids, func(ids []int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("IsBlacks in async mode is fine, buids: %v.\n", ids)
			} else {
				fmt.Printf("IsBlacks in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("IsBlacks in async mode error, err: %v.\n", err)
		})
	}

}

func isBlack(client *rtm.RTMServerClient) {
	// sync method
	ok, err := client.IsBlack(uid, buid)
	locker.print(func() {
		if err == nil {
			fmt.Printf("IsBlack in sync mode is fine, ok: %t.\n", ok)
		} else {
			fmt.Printf("IsBlack in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	_, err = client.IsBlack(uid, buid, func(ok bool, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("IsBlack in async mode is fine, ok: %t.\n", ok)
			} else {
				fmt.Printf("IsBlack in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("IsBlack in async mode error, err: %v.\n", err)
		})
	}
}

func delBlacks(client *rtm.RTMServerClient) {
	// sync method
	err := client.DelBlacks(uid, buids)
	locker.print(func() {
		if err == nil {
			fmt.Println("DelBlacks in sync mode is fine.")
		} else {
			fmt.Printf("DelBlacks in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	err = client.DelBlacks(uid, buids, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Println("DelBlacks in async mode is fine.")
			} else {
				fmt.Printf("DelBlacks in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("DelBlacks in async mode error, err: %v.\n", err)
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
		fmt.Println("pid is invalid. Error:", err)
		return
	}

	client := rtm.NewRTMServerClient(int32(pid), os.Args[3], os.Args[1])

	buids = append(buids, buid)

	addBlacks(client)
	getBlacks(client)
	isBlack(client)
	isBlacks(client)
	time.Sleep(400 * time.Millisecond)
	delBlacks(client)
	time.Sleep(400 * time.Millisecond)
	getBlacks(client)
	locker.print(func() {
		fmt.Println("Wait 2 second for async callback all print")
	})
	time.Sleep(2 * time.Second)
}
