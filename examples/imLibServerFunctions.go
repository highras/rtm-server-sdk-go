package main

import (
	"flag"
	"log"
	"runtime"
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
	uid        int64 = 1111
	gid        int64 = 1133
	uids             = []int64{135, 222, 666, 1207, 12345}
	nakes            = []string{"张三", "李四", "王五", "老刘", "老赵"}
	levels           = []string{"1", "2", "3", "2", "3"}
	customData       = []string{"i am 135", "i am 222", "i am 666", "i am 1207", "i am 12345"}
	gids             = []int64{666, 777, 888, 999}
)

//--------------[Demo]---------------------------------------
func setGroupInfos(client *rtm.RTMServerClient) {
	// 设置群组信息
	//param infos目前key支持：name,portraitUrl,profile,customData
	groupInfos := map[string]string{}
	groupInfos["name"] = "我是群主"
	groupInfos["customData"] = "extra group 1133"
	err1 := client.IMServer_SetGroupInfos(gid, groupInfos)
	locker.print(func() {
		if err1 == nil {
			log.Println("IMServer_SetGroupInfos in sync mode success")
		} else {
			log.Printf("IMServer_SetGroupInfos in sync mode failed , err: %v.\n", err1)
		}
	})
	/*
		// async method
		err1 = client.IMServer_SetGroupInfos(gid, groupInfos, func(errorCode int, errInfo string) {
			locker.print(func() {
				if errorCode == fpnn.FPNN_EC_OK {
					log.Println("IMServer_SetGroupInfos in async mode is ok.")
				} else {
					log.Printf("IMServer_SetGroupInfos in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
				}
			})
		})
		if err1 != nil {
			locker.print(func() {
				log.Printf("IMServer_SetGroupInfos in async mode error, err: %v.\n", err1)
			})
		}
	*/
}
func setInfos(client *rtm.RTMServerClient) {
	// sync method
	// 设置用户信息
	// param infos目前key支持：nick,portraitUrl,profile,level,customData
	for key, uid_i := range uids {
		userInfos := map[string]string{}
		userInfos["nick"] = nakes[key]
		userInfos["level"] = levels[key]
		userInfos["customData"] = customData[key]
		err := client.IMServer_SetUserInfos(uid_i, userInfos)
		locker.print(func() {
			if err == nil {
				log.Println("IMServer_SetUserInfos in sync mode is ok.")
			} else {
				log.Printf("IMServer_SetUserInfos in sync mode error, err:%v\n", err)
			}
		})
	}

	/*
		// async method
		err = client.IMServer_SetUserInfos(uid, userInfos, func(errorCode int, errInfo string) {
			locker.print(func() {
				if errorCode == fpnn.FPNN_EC_OK {
					log.Println("IMServer_SetUserInfos in async mode is ok.")
				} else {
					log.Printf("IMServer_SetUserInfos in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
				}
			})
		})
		if err != nil {
			locker.print(func() {
				log.Printf("IMServer_SetUserInfos in async mode error, err: %v.\n", err)
			})
		}
	*/
	time.Sleep(400 * time.Millisecond)

}

func getInfos(client *rtm.RTMServerClient) {
	// sync method
	//获取被添加权限
	//param: type：0:user 1:group 2:room
	//return grant_type:  0：允许任何人添加，1：需要验证，2：拒绝任何人添加，3：需要密码
	grant_type, err := client.IMServer_GetApplyGrant(1, gid)
	locker.print(func() {
		if err == nil {
			log.Printf("IMServer_GetApplyGrant in sync mode is fine, gid: %d, grant_type: %d.\n", gid, grant_type)
		} else {
			log.Printf("IMServer_GetApplyGrant in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	_, err = client.IMServer_GetApplyGrant(1, gid, func(grant_type int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_GetApplyGrant in async mode is fine, gid: %d, grant_type: %d.\n", gid, grant_type)
			} else {
				log.Printf("IMServer_GetApplyGrant in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			log.Printf("IMServer_GetApplyGrant in async mode error, err: %v.\n", err)
		})
	}

	// sync method
	/*
		获取邀请人群/房间权限
		param: type：0:user 1:group 2:room
		return invite_type:  0：不允许群成员邀请，1：允许群成员邀请
		return invite_manage_type： 0：邀请同意后需要管理员审核，1：邀请同意后直接入群
	*/
	invite_type, invite_manage_type, err1 := client.IMServer_GetInviteGrant(1, gid)
	locker.print(func() {
		if err1 == nil {
			log.Printf("IMServer_GetInviteGrant in sync mode is fine, gid: %d, invite_type: %d, invite_manage_type: %d.\n", gid, invite_type, invite_manage_type)
		} else {
			log.Printf("IMServer_GetInviteGrant in sync mode error, errinfo: %v.\n", err1)
		}
	})

	// async method
	_, _, err1 = client.IMServer_GetInviteGrant(1, gid, func(invite_type int32, invite_manage_type int32, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_GetInviteGrant in async mode is fine, gid: %d, invite_type: %d, invite_manage_type: %d.\n", gid, invite_type, invite_manage_type)
			} else {
				log.Printf("IMServer_GetInviteGrant in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})

	if err1 != nil {
		locker.print(func() {
			log.Printf("IMServer_GetInviteGrant in async mode error, err: %v.\n", err1)
		})
	}

	// sync method
	// return infos： uid(toString) => key => value
	// 目前key支持：nick,portraitUrl,profile,level,customData
	infos, err2 := client.IMServer_GetUserInfos(uids)
	locker.print(func() {
		if err2 == nil {
			log.Printf("IMServer_GetUserInfos in sync mode is fine, uids: %+v, infos: %+v.\n", uids, infos)
		} else {
			log.Printf("IMServer_GetUserInfos in sync mode error, errinfo: %v.\n", err1)
		}
	})

	// async method
	_, err2 = client.IMServer_GetUserInfos(uids, func(infos map[string]map[string]string, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_GetUserInfos in async mode is fine, uids: %+v, infos: %+v.\n", uids, infos)
			} else {
				log.Printf("IMServer_GetUserInfos in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err2 != nil {
		locker.print(func() {
			log.Printf("IMServer_GetUserInfos in async mode error, err: %v.\n", err2)
		})
	}

	// sync method
	// return infos： gid(toString) => key => value
	// 目前key支持：name,portraitUrl,profile,customData
	ginfos, err3 := client.IMServer_GetGroupInfos(gids)
	locker.print(func() {
		if err3 == nil {
			log.Printf("IMServer_GetGroupInfos in sync mode is fine, gids: %+v, infos: %+v.\n", gids, ginfos)
		} else {
			log.Printf("IMServer_GetGroupInfos in sync mode error, errinfo: %v.\n", err1)
		}
	})

	// async method
	_, err3 = client.IMServer_GetGroupInfos(gids, func(infos map[string]map[string]string, errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_GetGroupInfos in async mode is fine, gids: %+v, infos: %+v.\n", gids, infos)
			} else {
				log.Printf("IMServer_GetGroupInfos in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err3 != nil {
		locker.print(func() {
			log.Printf("IMServer_GetGroupInfos in async mode error, err: %v.\n", err3)
		})
	}

	time.Sleep(400 * time.Millisecond)

}

func options(client *rtm.RTMServerClient) {
	//创建群组
	// sync
	/*
		param: owner_uid: 群主UID
		param: infos: 初始群信息，key包括name,portraitUrl,profile,customData
		param: permissions: 初始权限信息： key包括：type(0：允许任何人添加，1：需要验证，2：拒绝任何人添加，3：需要密码),   extra(当type=3时为密码)
		return:
	*/
	groupInfos := map[string]string{}
	groupInfos["name"] = "i am group 666"
	groupInfos["customData"] = "extra group aaa"

	err2 := client.IMServer_CreateGroup(gid, uid, groupInfos, map[string]string{"type": "0"})
	locker.print(func() {
		if err2 == nil {
			log.Printf("IMServer_CreateGroup in sync mode is fine, gid: %d, uid: %d, groupInfos: %v.\n", gid, uid, groupInfos)
		} else {
			log.Printf("IMServer_CreateGroup in sync mode error, errinfo: %v.\n", err2)
		}
	})

	// async method
	err2 = client.IMServer_CreateGroup(gid, uid, groupInfos, map[string]string{"type": "0"}, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_CreateGroup in async mode is fine, gid: %d, uid: %d, groupInfos: %v.\n", gid, uid, groupInfos)
			} else {
				log.Printf("IMServer_CreateGroup in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err2 != nil {
		locker.print(func() {
			log.Printf("IMServer_CreateGroup in async mode error, err: %v.\n", err2)
		})
	}

	//设置被添加权限
	// sync
	/*
		param: type：0:user 1:group 2:room
		param: xid:  uid/gid/rid
		param: grant_type： 0：允许任何人添加，1：需要验证，2：拒绝任何人添加，3：需要密码
		param: extra: 当grant_type为需要密码时必须，为添加密码
		return:
	*/
	grant_type := 1
	err := client.IMServer_SetApplyGrant(1, gid, int32(grant_type))
	locker.print(func() {
		if err == nil {
			log.Printf("IMServer_SetApplyGrant in sync mode is fine, gid: %d, grant_type: %d.\n", gid, grant_type)
		} else {
			log.Printf("IMServer_SetApplyGrant in sync mode error, errinfo: %v.\n", err)
		}
	})

	// async method
	grant_type = 3
	extra := "333"
	err = client.IMServer_SetApplyGrant(1, gid, int32(grant_type), "333", func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_SetApplyGrant in async mode is fine, gid: %d, grant_type: %d, extra: %s.\n", gid, grant_type, extra)
			} else {
				log.Printf("IMServer_SetApplyGrant in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err != nil {
		locker.print(func() {
			log.Printf("IMServer_SetApplyGrant in async mode error, err: %v.\n", err)
		})
	}

	//设置群/房间邀请权限
	// sync
	/*
		param: type：1:group 2:room
		param: xid:  gid或rid
		param: invite_type：0：不允许群成员邀请，1：允许群成员邀请
		param: invite_manage_type:  0：邀请同意后需要管理员审核，1：邀请同意后直接入群
		return:
	*/
	invite_type := 1
	invite_manage_type := 1
	err1 := client.IMServer_SetInviteGrant(1, gid, int32(invite_type), int32(invite_manage_type))
	locker.print(func() {
		if err1 == nil {
			log.Printf("IMServer_SetInviteGrant in sync mode is fine, gid: %d, invite_type: %d, invite_manage_type: %d.\n", gid, invite_type, invite_manage_type)
		} else {
			log.Printf("IMServer_SetInviteGrant in sync mode error, errinfo: %v.\n", err1)
		}
	})

	// async method
	err1 = client.IMServer_SetInviteGrant(1, gid, int32(invite_type), int32(invite_manage_type), func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_SetInviteGrant in async mode is fine, gid: %d, invite_type: %d, invite_manage_type: %d.\n", gid, invite_type, invite_manage_type)
			} else {
				log.Printf("IMServer_SetInviteGrant in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err1 != nil {
		locker.print(func() {
			log.Printf("IMServer_SetInviteGrant in async mode error, err: %v.\n", err1)
		})
	}
	time.Sleep(400 * time.Millisecond)

}

func cleanUpGroup(client *rtm.RTMServerClient) {
	// 解散群组
	err3 := client.IMServer_DismissGroup(gid)
	locker.print(func() {
		if err3 == nil {
			log.Printf("IMServer_DismissGroup in sync mode is fine, gid: %d.\n", gid)
		} else {
			log.Printf("IMServer_DismissGroup in sync mode error, errinfo: %v.\n", err3)
		}
	})

	// async method
	err3 = client.IMServer_DismissGroup(gid, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				log.Printf("IMServer_DismissGroup in async mode is fine, gid: %d.\n", gid)
			} else {
				log.Printf("IMServer_DismissGroup in async mode error, error code: %v, error info: %v.\n", errorCode, errInfo)
			}
		})
	})
	if err3 != nil {
		locker.print(func() {
			log.Printf("IMServer_CreateGroup in async mode error, err: %v.\n", err3)
		})
	}
}

var (
	endpoint  = flag.String("a", "161.189.171.91:13315", "server address")
	pid       = flag.Int("p", 11000002, "project id")
	secretKey = flag.String("s", "f5a45c68-2279-4de7-b00e-aa10287531a8", "project secretKey")
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	client := rtm.NewRTMServerClient(int32(*pid), *secretKey, *endpoint)
	client.SetKeepAlive(true)
	time.Sleep(1 * time.Millisecond)
	options(client)
	setGroupInfos(client)
	setInfos(client)
	getInfos(client)
	cleanUpGroup(client)
	time.Sleep(400 * time.Millisecond)

	select {}
}
