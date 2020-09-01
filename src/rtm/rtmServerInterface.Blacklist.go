package rtm

import (
	"errors"
	"time"
)

//-----------[Black Functions]-------------

/*
	param:
		rest: can be include following params:
			timeout time.Duration
			func(errorCode int, errorInfo string)

		If include func param, this function will enter into async mode, and return(error),
		else this function work in sync mode, and return(err error)
*/
func (client *RTMServerClient) AddBlacks(uid int64, blacks []int64, rest ...interface{}) error {
	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invalid params when call RTMServerClient.AddBlacks() function.")
		}
	}

	quest := client.genServerQuest("addblacks")
	quest.Param("uid", uid)
	quest.Param("blacks", blacks)
	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	param:
		rest: can be include following params:
			timeout time.Duration
			func(errorCode int, errorInfo string)

		If include func param, this function will enter into async mode, and return(error),
		else this function work in sync mode, and return(err error)
*/
func (client *RTMServerClient) DelBlacks(uid int64, blacks []int64, rest ...interface{}) error {
	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invalid params when call RTMServerClient.DelBlacks() function.")
		}
	}

	quest := client.genServerQuest("delblacks")
	quest.Param("uid", uid)
	quest.Param("blacks", blacks)
	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	param:
		rest: can be include following params:
			timeout time.Duration
			func(buids []int64, errorCode int, errorInfo string)

		If include func param, this function will enter into async mode, and return(error),
		else this function work in sync mode, and return(buids []int64, err error)
*/
func (client *RTMServerClient) GetBlacks(uid int64, rest ...interface{}) ([]int64, error) {
	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invalid params when call RTMServerClient.GetBlacks() function.")
		}
	}

	quest := client.genServerQuest("getblacks")
	quest.Param("uid", uid)
	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	param:
		rest: can be include following params:
			timeout time.Duration
			func(ok bool, errorCode int, errorInfo string)

		If include func param, this function will enter into async mode, and return(error),
		else this function work in sync mode, and return(ok bool, err error)
*/
func (client *RTMServerClient) IsBlack(uid, buid int64, rest ...interface{}) (bool, error) {
	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invalid params when call RTMServerClient.IsBlack() function.")
		}
	}

	quest := client.genServerQuest("isblack")
	quest.Param("uid", uid)
	quest.Param("buid", buid)
	return client.sendOkQuest(quest, timeout, callback)
}

/*
	param:
		rest: can be include following params:
			timeout time.Duration
			func(buids []int64, errorCode int, errorInfo string)

		If include func param, this function will enter into async mode, and return(error),
		else this function work in sync mode, and return(buids []int64, err error)
*/
func (client *RTMServerClient) IsBlacks(uid int64, buids []int64, rest ...interface{}) ([]int64, error) {
	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invalid params when call RTMServerClient.IsBlacks() function.")
		}
	}

	quest := client.genServerQuest("isblacks")
	quest.Param("uid", uid)
	quest.Param("buids", buids)
	return client.sendSliceQuest(quest, timeout, "buids", callback)
}
