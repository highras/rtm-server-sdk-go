package rtm

import (
	"errors"
	"fmt"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"strconv"
	"time"
)

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddDevice() function.")
		}
	}

	quest := client.genServerQuest("adddevice")
	quest.Param("uid", uid)
	quest.Param("apptype", appType)
	quest.Param("devicetoken", deviceToken)

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
func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveDevice() function.")
		}
	}

	quest := client.genServerQuest("removedevice")
	quest.Param("uid", uid)
	quest.Param("devicetoken", deviceToken)

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
func (client *RTMServerClient) AddDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddDevicePushOption() function.")
		}
	}

	quest := client.genServerQuest("addoption")
	quest.Param("uid", uid)
	quest.Param("type", pushType)
	quest.Param("xid", xid)
	if mtype != nil && len(mtype) > 0 {
		quest.Param("mtypes", mtype)
	}

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
func (client *RTMServerClient) RemoveDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveDevicePushOption() function.")
		}
	}

	quest := client.genServerQuest("removeoption")
	quest.Param("uid", uid)
	quest.Param("type", pushType)
	quest.Param("xid", xid)
	if mtype != nil && len(mtype) > 0 {
		quest.Param("mtypes", mtype)
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

func (client *RTMServerClient) convertValueToInt8(value interface{}) int8 {
	switch tmp := value.(type) {
	case int64:
		return int8(tmp)
	case int32:
		return int8(tmp)
	case int16:
		return int8(tmp)
	case int8:
		return int8(tmp)
	case int:
		return int8(tmp)

	case uint64:
		return int8(tmp)
	case uint32:
		return int8(tmp)
	case uint16:
		return int8(tmp)
	case uint8:
		return int8(tmp)
	case uint:
		return int8(tmp)

	case float32:
		return int8(tmp)
	case float64:
		return int8(tmp)
	default:
		client.logger.Println("[ERROR] convertValueToInt8 Type convert failed.")
		return 0
	}
}

func (client *RTMServerClient) convertToInt64Map(value map[interface{}]interface{}) map[int64][]int8 {

	result := make(map[int64][]int8)

	for k, v := range value {
		key := client.convertToString(k)
		var result_type []int8
		if data, ok := v.([]interface{}); ok {
			for _, tmp := range data {
				tmp_r := client.convertValueToInt8(tmp)
				result_type = append(result_type, tmp_r)
			}
		}
		if i, err := strconv.ParseInt(key, 10, 64); err == nil {
			result[i] = result_type
		}
	}

	return result
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (p2pOption map[int64][]int8, groupOption map[int64][]int8, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) GetDevicePushOption(uid int64, rest ...interface{}) (map[int64][]int8, map[int64][]int8, error) {

	var timeout time.Duration
	var callback func(map[int64][]int8, map[int64][]int8, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(map[int64][]int8, map[int64][]int8, int, string):
			callback = value
		default:
			return nil, nil, errors.New("Invaild params when call RTMServerClient.GetDevicePushOption() function.")
		}
	}

	quest := client.genServerQuest("getoption")
	quest.Param("uid", uid)

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertToInt64Map(answer.WantMap("p2p")), client.convertToInt64Map(answer.WantMap("group")), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(nil, nil, errorCode, "")
			} else {
				callback(nil, nil, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return nil, nil, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return nil, nil, err
	} else if !answer.IsException() {
		p2pMap := client.convertToInt64Map(answer.WantMap("p2p"))
		groupMap := client.convertToInt64Map(answer.WantMap("group"))
		return p2pMap, groupMap, nil
	} else {
		return nil, nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}
