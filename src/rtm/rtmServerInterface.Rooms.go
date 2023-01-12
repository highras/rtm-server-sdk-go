package rtm

import (
	"errors"
	"fmt"
	"strconv"
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
func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddRoomMember() function.")
		}
	}

	quest := client.genServerQuest("addroommember")
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string, []int64)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddRoomMembers(roomId int64, uids []int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.AddRoomMembers() function.")
		}
	}

	quest := client.genServerQuest("addroommembers")
	quest.Param("rid", roomId)
	quest.Param("uids", uids)

	return client.sendSliceQuest(quest, timeout, "successedUids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DelRoomMember() function.")
		}
	}

	quest := client.genServerQuest("delroommember")
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
func (client *RTMServerClient) DelRoomMembers(roomId int64, uids []int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.DelRoomMembers() function.")
		}
	}

	quest := client.genServerQuest("delroommembers")
	quest.Param("rid", roomId)
	quest.Param("uids", uids)

	return client.sendSliceQuest(quest, timeout, "successedUids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddRoomBan() function.")
		}
	}

	quest := client.genServerQuest("addroomban")
	if roomId > 0 {
		quest.Param("rid", roomId)
	}
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
func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveRoomBan() function.")
		}
	}

	quest := client.genServerQuest("removeroomban")
	if roomId > 0 {
		quest.Param("rid", roomId)
	}
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
func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ...interface{}) (bool, error) {

	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invaild params when call RTMServerClient.IsBanOfRoom() function.")
		}
	}

	quest := client.genServerQuest("isbanofroom")
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
func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetRoomInfo() function.")
		}
	}

	quest := client.genServerQuest("setroominfo")
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
func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ...interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			return "", "", errors.New("Invaild params when call RTMServerClient.GetRoomInfo() function.")
		}
	}

	quest := client.genServerQuest("getroominfo")
	quest.Param("rid", roomId)

	return client.sendGetObjectInfoQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ([]int64, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetRoomMembers(roomId int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetRoomMembers() function.")
		}
	}

	quest := client.genServerQuest("getroommembers")
	quest.Param("rid", roomId)
	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

func (client *RTMServerClient) convertToInt64MapByInt32(value map[interface{}]interface{}) map[int64]int32 {

	result := make(map[int64]int32)

	for k, v := range value {
		key := client.convertToString(k)
		count := int32(client.convertToInt64(v))
		if i, err := strconv.ParseInt(key, 10, 64); err == nil {
			result[i] = count
		}
	}

	return result
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (count int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (map[int64]int32, error);
		else this function work in sync mode, and return (count map[int64]int32, err error)
*/
func (client *RTMServerClient) GetRoomCount(roomIds []int64, rest ...interface{}) (map[int64]int32, error) {

	var timeout time.Duration
	var callback func(map[int64]int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(map[int64]int32, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetRoomCount() function.")
		}
	}

	quest := client.genServerQuest("getroomcount")
	quest.Param("rids", roomIds)
	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertToInt64MapByInt32(answer.WantMap("cn")), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(nil, errorCode, "")
			} else {
				callback(nil, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return nil, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return nil, err
	} else if !answer.IsException() {
		return client.convertToInt64MapByInt32(answer.WantMap("cn")), nil
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddUserRooms(roomIds []int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddUserRooms() function.")
		}
	}

	quest := client.genServerQuest("adduserrooms")
	quest.Param("rids", roomIds)
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
func (client *RTMServerClient) DeleteUserRooms(roomIds []int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DeleteUserRooms() function.")
		}
	}

	quest := client.genServerQuest("deleteuserrooms")
	quest.Param("rids", roomIds)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}
