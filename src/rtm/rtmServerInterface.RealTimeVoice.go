package rtm

import (
	"errors"
	"time"
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
func (client *RTMServerClient) InviteUserIntoVoiceRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.InviteUserIntoVoiceRoom() function.")
		}
	}

	quest := client.genServerQuest("inviteUserIntoVoiceRoom")
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
func (client *RTMServerClient) CloseVoiceRoom(roomId int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.CloseVoiceRoom() function.")
		}
	}

	quest := client.genServerQuest("closeVoiceRoom")
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
func (client *RTMServerClient) KickoutFromVoiceRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.KickoutFromVoiceRoom() function.")
		}
	}

	quest := client.genServerQuest("kickoutFromVoiceRoom")
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
func (client *RTMServerClient) GetVoiceRoomList(rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetVoiceRoomList() function.")
		}
	}

	quest := client.genServerQuest("getVoiceRoomList")

	return client.sendSliceQuest(quest, timeout, "rids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, managers []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (uids []int64, managers []int64, err error)
*/
func (client *RTMServerClient) GetVoiceRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, error) {

	var timeout time.Duration
	var callback func([]int64, []int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, []int64, int, string):
			callback = value
		default:
			return nil, nil, errors.New("Invaild params when call RTMServerClient.GetVoiceRoomMembers() function.")
		}
	}

	quest := client.genServerQuest("getVoiceRoomMembers")
	quest.Param("rid", roomId)

	return client.sendDoubleSliceQuest(quest, timeout, "uids", "managers", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (count int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (count int32, err error)
*/
func (client *RTMServerClient) GetVoiceRoomMemberCount(roomId int64, rest ...interface{}) (int32, error) {

	var timeout time.Duration
	var callback func(int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.GetVoiceRoomMemberCount() function.")
		}
	}

	quest := client.genServerQuest("getVoiceRoomMemberCount")
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
func (client *RTMServerClient) SetVoiceRoomMicStatus(roomId int64, status bool, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetVoiceRoomMicStatus() function.")
		}
	}

	quest := client.genServerQuest("setVoiceRoomMicStatus")
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
func (client *RTMServerClient) PullUserIntoVoiceRoom(roomId int64, uids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.PullUserIntoVoiceRoom() function.")
		}
	}

	quest := client.genServerQuest("pullIntoVoiceRoom")
	quest.Param("rid", roomId)
	quest.Param("toUids", uids)

	return client.sendSilentQuest(quest, timeout, callback)
}
