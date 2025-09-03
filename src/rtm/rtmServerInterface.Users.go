package rtm

import (
	"errors"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

/*
Params:

rest: can be include following params:

	timeout time.Duration
	func (uids []int64, errorCode int, errInfo string)

If include func param, this function will enter into async mode, and return (nil, error);
else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetOnlineUsers() function.")
		}
	}

	quest := client.genServerQuest("getonlineusers")
	quest.Param("uids", uids)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
Params:

rest: can be include following params:

	timeout time.Duration
	func (errorCode int, errInfo string)

If include func param, this function will enter into async mode, and return (error);
else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddProjectBlack() function.")
		}
	}

	quest := client.genServerQuest("addprojectblack")
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
func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveProjectBlack() function.")
		}
	}

	quest := client.genServerQuest("removeprojectblack")
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
func (client *RTMServerClient) IsProjectBlack(uid int64, rest ...interface{}) (bool, error) {

	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invaild params when call RTMServerClient.IsProjectBlack() function.")
		}
	}

	quest := client.genServerQuest("isprojectblack")
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
Params:

rest: can be include following params:

	timeout time.Duration
	func (errorCode int, errInfo string)

If include func param, this function will enter into async mode, and return (error);
else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddProjectBan(uid int64, bannedSeconds int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddProjectBan() function.")
		}
	}

	quest := client.genServerQuest("addprojectban")
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
func (client *RTMServerClient) RemoveProjectBan(uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveProjectBan() function.")
		}
	}

	quest := client.genServerQuest("removeprojectban")
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
func (client *RTMServerClient) IsProjectBan(uid int64, rest ...interface{}) (bool, error) {

	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invaild params when call RTMServerClient.IsProjectBan() function.")
		}
	}

	quest := client.genServerQuest("isprojectban")
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
func (client *RTMServerClient) SetUserInfo(uid int64, publicInfo *string, privateInfo *string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetUserInfo() function.")
		}
	}

	quest := client.genServerQuest("setuserinfo")
	quest.Param("uid", uid)

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
func (client *RTMServerClient) GetUserInfo(uid int64, rest ...interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			return "", "", errors.New("Invaild params when call RTMServerClient.GetUserInfo() function.")
		}
	}

	quest := client.genServerQuest("getuserinfo")
	quest.Param("uid", uid)

	return client.sendGetObjectInfoQuest(quest, timeout, callback)
}

func (client *RTMServerClient) convertUserPublicInfoMap(info map[interface{}]interface{}) map[string]string {

	result := make(map[string]string)

	for k, v := range info {
		key := client.convertToString(k)
		result[key] = client.convertToString(v)
	}

	return result
}

func (client *RTMServerClient) sendGetUserPublicInfoQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(info map[string]string, errorCode int, errInfo string)) (map[string]string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertUserPublicInfoMap(answer.WantMap("info")), fpnn.FPNN_EC_OK, "")
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
		return client.convertUserPublicInfoMap(answer.WantMap("info")), nil
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
Params:

rest: can be include following params:

	timeout time.Duration
	func (map[string]string, errorCode int, errInfo string)

If include func param, this function will enter into async mode, and return (nil, error);
else this function work in sync mode, and return (map[string]string, err error)
*/
func (client *RTMServerClient) GetUserPublicInfo(uids []int64, rest ...interface{}) (map[string]string, error) {

	var timeout time.Duration
	var callback func(map[string]string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(map[string]string, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetUserPublicInfo() function.")
		}
	}

	quest := client.genServerQuest("getuseropeninfo")
	quest.Param("uids", uids)

	return client.sendGetUserPublicInfoQuest(quest, timeout, callback)
}

/*
Params:

rest: can be include following params:

	timeout time.Duration
	func (errorCode int, errInfo string)

If include func param, this function will enter into async mode, and return (error);
else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) ClearUserMessage(uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.ClearUserMessage() function.")
		}
	}

	quest := client.genServerQuest("clearusermsg")
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
func (client *RTMServerClient) AddSessions(uid int64, otherUids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddSessions() function.")
		}
	}

	quest := client.genServerQuest("addsessions")
	quest.Param("uid", uid)
	quest.Param("uids", otherUids)

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
func (client *RTMServerClient) RemoveSessions(uid int64, otherUids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddSessions() function.")
		}
	}

	quest := client.genServerQuest("removesessions")
	quest.Param("uid", uid)
	quest.Param("uids", otherUids)

	return client.sendSilentQuest(quest, timeout, callback)
}
