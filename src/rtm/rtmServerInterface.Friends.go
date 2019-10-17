package rtm

import (
	"time"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

//-----------[ Friends functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("friends", firendUids)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("friends", firendUids)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsFriend() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isfriend")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("fuid", peerUid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (firendUids []int64, err error)
*/
func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("fuids", uids)

	return client.sendSliceQuest(quest, timeout, "fuids", callback)
}
