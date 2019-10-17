package rtm

import (
	"time"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
)


//-----------[ Room functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddRoomMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addroommember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

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
func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelRoomMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delroommember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

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
func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddRoomBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addroomban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)
	quest.Param("btime", bannedSeconds)

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
func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveRoomBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removeroomban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsBanOfRoom() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isbanofroom")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		publicInfo:
			Public info.
			Nil pointer means ignore the params when invoking.

		privateInfo:
			Private info.
			Nil pointer means ignore the params when invoking.

		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SetRoomInfo() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("setroominfo")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("rid", roomId)

	if publicInfo != nil {
		quest.Param("oinfo", *publicInfo)
	}
	
	if privateInfo != nil {
		quest.Param("pinfo", *privateInfo)
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (publicInfo string, privateInfo string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", "", error);
		else this function work in sync mode, and return (publicInfo string, privateInfo string, err error)
*/
func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ... interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func (string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (string, string, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetRoomInfo() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getroominfo")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)

	return client.sendGetObjectInfoQuest(quest, timeout, callback)
}
