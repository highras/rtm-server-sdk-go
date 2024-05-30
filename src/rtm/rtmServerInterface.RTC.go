package rtm

import (
	"errors"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

//-----------[ real-time-voice functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) InviteUserIntoRTCRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.InviteUserIntoRTCRoom() function.")
		}
	}

	quest := client.genServerQuest("inviteUserIntoRTCRoom")
	quest.Param("rid", roomId)
	quest.Param("toUids", uids)
	quest.Param("fromUid", fromUid)

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
func (client *RTMServerClient) CloseRTCRoom(roomId int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.CloseRTCRoom() function.")
		}
	}

	quest := client.genServerQuest("closeRTCRoom")
	quest.Param("rid", roomId)

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
func (client *RTMServerClient) KickoutFromRTCRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.KickoutFromRTCRoom() function.")
		}
	}

	quest := client.genServerQuest("kickoutFromRTCRoom")
	quest.Param("rid", roomId)
	quest.Param("fromUid", fromUid)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (rids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (rids []int64, err error)
*/
func (client *RTMServerClient) GetRTCRoomList(rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetRTCRoomList() function.")
		}
	}

	quest := client.genServerQuest("getRTCRoomList")

	return client.sendSliceQuest(quest, timeout, "rids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, managers []int64, ownerId int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (uids []int64, managers []int64, ownerId int64, err error)
*/
func (client *RTMServerClient) GetRTCRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, int64, error) {

	var timeout time.Duration
	var callback func([]int64, []int64, int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, []int64, int64, int, string):
			callback = value
		default:
			return nil, nil, 0, errors.New("Invaild params when call RTMServerClient.GetRTCRoomMembers() function.")
		}
	}

	quest := client.genServerQuest("getRTCRoomMembers")
	quest.Param("rid", roomId)

	return client.sendDoubleSliceAndIntQuest(quest, timeout, "uids", "administrators", "owner", callback)
}

func (client *RTMServerClient) sendDoubleSliceAndIntQuest(quest *fpnn.Quest, timeout time.Duration,
	sliceKey string, secondSliceKey string, keyInt string, callback func(first []int64, second []int64, uid int64, errorCode int, errInfo string)) ([]int64, []int64, int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertSliceToInt64Slice(answer.WantSlice(sliceKey)), client.convertSliceToInt64Slice(answer.WantSlice(secondSliceKey)), answer.WantInt64(keyInt), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(nil, nil, 0, errorCode, "")
			} else {
				callback(nil, nil, 0, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return nil, nil, 0, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return nil, nil, 0, err
	} else if !answer.IsException() {
		slice1 := client.convertSliceToInt64Slice(answer.WantSlice(sliceKey))
		slice2 := client.convertSliceToInt64Slice(answer.WantSlice(secondSliceKey))
		return slice1, slice2, answer.WantInt64(keyInt), nil
	} else {
		return nil, nil, 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (count int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (count int32, err error)
*/
func (client *RTMServerClient) GetRTCRoomMemberCount(roomId int64, rest ...interface{}) (int32, error) {

	var timeout time.Duration
	var callback func(int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.GetRTCRoomMemberCount() function.")
		}
	}

	quest := client.genServerQuest("getRTCRoomMemberCount")
	quest.Param("rid", roomId)

	return client.sendIntQuest(quest, timeout, "count", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetRTCRoomMicStatus(roomId int64, status bool, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetRTCRoomMicStatus() function.")
		}
	}

	quest := client.genServerQuest("setRTCRoomMicStatus")
	quest.Param("rid", roomId)
	quest.Param("status", status)

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
func (client *RTMServerClient) PullUserIntoRTCRoom(roomId int64, uids []int64, roomType int32, voiceRange int32, maxReceiveStreams int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.PullUserIntoRTCRoom() function.")
		}
	}

	quest := client.genServerQuest("pullIntoRTCRoom")
	quest.Param("rid", roomId)
	quest.Param("toUids", uids)
	quest.Param("type", roomType)
	quest.Param("voiceRange", voiceRange)
	quest.Param("maxReceiveStreams", maxReceiveStreams)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		command: 0 赋予管理员权限，1 剥夺管理员权限，2 禁止发送音频数据，3 允许发送音频数据，
				 4 禁止发送视频数据，5 允许发送视频数据，6 关闭他人麦克风，7 关闭他人摄像头
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AdminCommand(roomId int64, uids []int64, command int32, rest ...interface{}) error {
	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AdminCommand() function.")
		}
	}

	quest := client.genServerQuest("adminCommand")
	quest.Param("rid", roomId)
	quest.Param("uids", uids)
	quest.Param("command", command)

	return client.sendSilentQuest(quest, timeout, callback)
}
