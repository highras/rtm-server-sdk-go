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

func demoSendChat(client *rtm.RTMServerClient) {

	//-- sync send P2P chat
	mtime, err := client.SendChat(fromUid, toUid, "test sync P2P chat")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[P2P Chat] %d send to %d in sync mode, return mtime: %d\n", fromUid, toUid, mtime)
		} else {
			fmt.Printf("[P2P Chat] %d send to %d in sync mode, err: %v\n", fromUid, toUid, err)
		}
	})

	//-- async send P2P chat
	_, err = client.SendChat(fromUid, toUid, "test async P2P chat", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[P2P Chat] %d send to %d in async mode, mtime:%d\n", fromUid, toUid, mtime)
			} else {
				fmt.Printf("[P2P Chat] %d send to %d in async mode, error code: %d, error info:%s\n",
					fromUid, toUid, errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[P2P Chat] %d send to %d in async mode, err: %v\n", fromUid, toUid, err)
		})
	}

	//-- sync send multiple P2P chat
	mtime, err = client.SendChats(fromUid, toUids, "test sync multiple P2P chat")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Multiple P2P Chat] %d send to {%v} in sync mode, return mtime: %d\n", fromUid, toUids, mtime)
		} else {
			fmt.Printf("[Multiple P2P Chat] %d send to {%v} in sync mode, err: %v\n", fromUid, toUids, err)
		}
	})

	//-- async send multiple P2P chat
	_, err = client.SendChats(fromUid, toUids, "test async multiple P2P chat", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Multiple P2P Chat] %d send to {%v} in async mode, mtime:%d\n", fromUid, toUids, mtime)
			} else {
				fmt.Printf("[Multiple P2P Chat] %d send to {%v} in async mode, error code: %d, error info:%s\n",
					fromUid, toUids, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Multiple P2P Chat] %d send to {%v} in async mode, err: %v\n", fromUid, toUids, err)
		})
	}

	//-- sync send group chat
	mtime, err = client.SendGroupChat(fromUid, groupId, "test sync group chat")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Group Chat] %d send to group %d in sync mode, return mtime: %d\n", fromUid, groupId, mtime)
		} else {
			fmt.Printf("[Group Chat] %d send to group %d in sync mode, err: %v\n", fromUid, groupId, err)
		}
	})

	//-- async send group chat
	_, err = client.SendGroupChat(fromUid, groupId, "test async group chat", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Group Chat] %d send to group %d in async mode, mtime:%d\n", fromUid, groupId, mtime)
			} else {
				fmt.Printf("[Group Chat] %d send to group %d in async mode, error code: %d, error info:%s\n",
					fromUid, groupId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Group Chat] %d send to group %d in async mode, err: %v\n", fromUid, groupId, err)
		})
	}

	//-- sync send room chat
	mtime, err = client.SendRoomChat(fromUid, roomId, "test sync room chat")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Room Chat] %d send to room %d in sync mode, return mtime: %d\n", fromUid, roomId, mtime)
		} else {
			fmt.Printf("[Room Chat] %d send to room %d in sync mode, err: %v\n", fromUid, roomId, err)
		}
	})

	//-- async send room chat
	_, err = client.SendRoomChat(fromUid, roomId, "test async room chat", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Room Chat] %d send to room %d in async mode, mtime:%d\n", fromUid, roomId, mtime)
			} else {
				fmt.Printf("[Room Chat] %d send to room %d in async mode, error code: %d, error info:%s\n",
					fromUid, roomId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Room Chat] %d send to room %d in async mode, err: %v\n", fromUid, roomId, err)
		})
	}

	//-- sync send boardcast chat
	mtime, err = client.SendBroadcastChat(adminUid, "test sync boardcast chat")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Boardcast Chat] %d send boardcast chat in sync mode, return mtime: %d\n", adminUid, mtime)
		} else {
			fmt.Printf("[Boardcast Chat] %d send boardcast chat in sync mode, err: %v\n", adminUid, err)
		}
	})

	//-- async send boardcast chat
	_, err = client.SendBroadcastChat(adminUid, "test async boardcast chat", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Boardcast Chat] %d send boardcast chat in async mode, mtime:%d\n", adminUid, mtime)
			} else {
				fmt.Printf("[Boardcast Chat] %d send boardcast chat in async mode, error code: %d, error info:%s\n",
					adminUid, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Boardcast Chat] %d send boardcast chat in async mode, err: %v\n", adminUid, err)
		})
	}
}

func demoSendCmd(client *rtm.RTMServerClient) {

	//-- sync send P2P audio
	mtime, err := client.SendCmd(fromUid, toUid, "test sync P2P cmd")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[P2P Cmd] %d send to %d in sync mode, return mtime: %d\n", fromUid, toUid, mtime)
		} else {
			fmt.Printf("[P2P Cmd] %d send to %d in sync mode, err: %v\n", fromUid, toUid, err)
		}
	})

	//-- async send P2P audio
	_, err = client.SendCmd(fromUid, toUid, "test async P2P cmd", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[P2P Cmd] %d send to %d in async mode, mtime:%d\n", fromUid, toUid, mtime)
			} else {
				fmt.Printf("[P2P Cmd] %d send to %d in async mode, error code: %d, error info:%s\n",
					fromUid, toUid, errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[P2P Cmd] %d send to %d in async mode, err: %v\n", fromUid, toUid, err)
		})
	}

	//-- sync send multiple P2P chat
	mtime, err = client.SendCmds(fromUid, toUids, "test sync multiple P2P cmd")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Multiple P2P Cmd] %d send to {%v} in sync mode, return mtime: %d\n", fromUid, toUids, mtime)
		} else {
			fmt.Printf("[Multiple P2P Cmd] %d send to {%v} in sync mode, err: %v\n", fromUid, toUids, err)
		}
	})

	//-- async send multiple P2P chat
	_, err = client.SendCmds(fromUid, toUids, "test async multiple P2P cmd", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Multiple P2P Cmd] %d send to {%v} in async mode, mtime:%d\n", fromUid, toUids, mtime)
			} else {
				fmt.Printf("[Multiple P2P Cmd] %d send to {%v} in async mode, error code: %d, error info:%s\n",
					fromUid, toUids, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Multiple P2P Cmd] %d send to {%v} in async mode, err: %v\n", fromUid, toUids, err)
		})
	}

	//-- sync send group chat
	mtime, err = client.SendGroupCmd(fromUid, groupId, "test sync group cmd")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Group Cmd] %d send to group %d in sync mode, return mtime: %d\n", fromUid, groupId, mtime)
		} else {
			fmt.Printf("[Group Cmd] %d send to group %d in sync mode, err: %v\n", fromUid, groupId, err)
		}
	})

	//-- async send group chat
	_, err = client.SendGroupCmd(fromUid, groupId, "test async group cmd", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Group Cmd] %d send to group %d in async mode, mtime:%d\n", fromUid, groupId, mtime)
			} else {
				fmt.Printf("[Group Cmd] %d send to group %d in async mode, error code: %d, error info:%s\n",
					fromUid, groupId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Group Cmd] %d send to group %d in async mode, err: %v\n", fromUid, groupId, err)
		})
	}

	//-- sync send room chat
	mtime, err = client.SendRoomCmd(fromUid, roomId, "test sync room cmd")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Room Cmd] %d send to room %d in sync mode, return mtime: %d\n", fromUid, roomId, mtime)
		} else {
			fmt.Printf("[Room Cmd] %d send to room %d in sync mode, err: %v\n", fromUid, roomId, err)
		}
	})

	//-- async send room chat
	_, err = client.SendRoomCmd(fromUid, roomId, "test async room cmd", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Room Cmd] %d send to room %d in async mode, mtime:%d\n", fromUid, roomId, mtime)
			} else {
				fmt.Printf("[Room Cmd] %d send to room %d in async mode, error code: %d, error info:%s\n",
					fromUid, roomId, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Room Cmd] %d send to room %d in async mode, err: %v\n", fromUid, roomId, err)
		})
	}

	//-- sync send boardcast chat
	mtime, err = client.SendBroadcastCmd(adminUid, "test sync boardcast cmd")
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Boardcast Cmd] %d send boardcast chat in sync mode, return mtime: %d\n", adminUid, mtime)
		} else {
			fmt.Printf("[Boardcast Cmd] %d send boardcast chat in sync mode, err: %v\n", adminUid, err)
		}
	})

	//-- async send boardcast chat
	_, err = client.SendBroadcastCmd(adminUid, "test async boardcast cmd", func(mtime int64, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Boardcast Cmd] %d send boardcast chat in async mode, mtime:%d\n", adminUid, mtime)
			} else {
				fmt.Printf("[Boardcast Cmd] %d send boardcast chat in async mode, error code: %d, error info:%s\n",
					adminUid, errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			fmt.Printf("[Boardcast Cmd] %d send boardcast chat in async mode, err: %v\n", adminUid, err)
		})
	}
}

func deleteChat(client *rtm.RTMServerClient) {

	var mid int64 = 123456

	//-- sync mode
	err := client.DelP2PChat(mid, fromUid, toUid)
	locker.print(func() {
		if err == nil {
			fmt.Printf("DelP2PChat in sync mode is fine.\n")
		} else {
			fmt.Printf("DelP2PChat in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.DelP2PChat(mid, fromUid, groupId, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("DelP2PChat in async mode is fine.\n")
			} else {
				fmt.Printf("DelP2PChat in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("DelP2PChat in async mode error, err: %v\n", err)
		})
	}
}

func demoTranslate(client *rtm.RTMServerClient) {

	sourceText := "Hello, Kitty!"

	//-- sync translate
	result, err := client.Translate(sourceText, "", "zh-CN", "", "", true, 0)
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Translate] Translate %s, return: %s\n", sourceText, result.TargetText)
		} else {
			fmt.Printf("[Translate] Translate %s, err: %v\n", sourceText, err)
		}
	})

	//-- async translate
	_, err = client.Translate(sourceText, "", "zh-CN", "", "", true, 0, func(result *rtm.TranslateResult, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Translate] Translate %s, return: %s\n", sourceText, result.TargetText)
			} else {
				fmt.Printf("[Translate] Translate %s, error code: %d, error info:%s\n",
					sourceText, errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[Translate] Translate %s, err: %v\n", sourceText, err)
		})
	}
}

func demoSpeech2Text(client *rtm.RTMServerClient) {

	audio := "test aaaaa" // 格式可能错误，此处只是展示api使用，具体格式请咨询client sdk人员

	//-- sync Speech2Text
	text, lang, err := client.Speech2Text(audio, 2, rtm.EN, "", 0, 0)
	locker.print(func() {
		if err == nil {
			fmt.Printf("[Speech2Text] Speech2Text text: %s, lang: %s\n", text, lang)
		} else {
			fmt.Printf("[Speech2Text] Speech2Text err: %v\n", err)
		}
	})

	//-- async Speech2Text
	_, _, err = client.Speech2Text(audio, 2, rtm.EN, "", 0, 0, func(text string, lang string, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[Speech2Text] Speech2Text text: %s, lang: %s\n", text, lang)
			} else {
				fmt.Printf("[Speech2Text] Speech2Text error code: %d, error info:%s\n",
					errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[Speech2Text] Speech2Text err: %v\n", err)
		})
	}
}

func demoTextCheck(client *rtm.RTMServerClient) {

	message := "shit"

	//-- sync TextCheck
	result, text, tags, wlist, err := client.TextCheck(message, 0)
	locker.print(func() {
		if err == nil {
			fmt.Printf("[TextCheck] TextCheck result: %d, text: %s, tags: %v, wlist: %v\n", result, text, tags, wlist)
		} else {
			fmt.Printf("[TextCheck] TextCheck err: %v\n", err)
		}
	})

	//-- async TextCheck
	_, _, _, _, err = client.TextCheck(message, 0, func(result int32, text string, tags []int32, wlist []string, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("[TextCheck] TextCheck result: %d, text: %s, tags: %v, wlist: %v\n", result, text, tags, wlist)
			} else {
				fmt.Printf("[TextCheck] TextCheck error code: %d, error info:%s\n",
					errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("[TextCheck] TextCheck err: %v\n", err)
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
	
	//demoSendChat(client)
	//demoSpeech2Text(client)
	demoTextCheck(client)
	time.Sleep(time.Second * 150) //-- Waiting for the async callback printed.
	return
	demoSendCmd(client)

	deleteChat(client)

	demoTranslate(client)

	locker.print(func() {
		fmt.Println("Wait 1 second for async callbacks are printed.")
	})

	time.Sleep(time.Second * 150) //-- Waiting for the async callback printed.
}
