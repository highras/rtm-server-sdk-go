package rtm

import (
	"errors"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

//-----------[ Data functions ]-------------------//

func (client *RTMServerClient) sendGetDataQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(text string, errorCode int, errInfo string)) (string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				value, _ := answer.GetString("val")
				callback(value, fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback("", errorCode, "")
			} else {
				callback("", answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return "", err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return "", err
	} else if !answer.IsException() {
		value, _ := answer.GetString("val")
		return value, nil
	} else {
		return "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (text string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", error);
		else this function work in sync mode, and return (text string, err error)
*/
func (client *RTMServerClient) GetData(uid int64, key string, rest ...interface{}) (string, error) {

	var timeout time.Duration
	var callback func(string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, int, string):
			callback = value
		default:
			return "", errors.New("Invaild params when call RTMServerClient.GetData() function.")
		}
	}

	quest := client.genServerQuest("dataget")
	quest.Param("uid", uid)
	quest.Param("key", key)

	return client.sendGetDataQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetData(uid int64, key string, value string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetData() function.")
		}
	}

	quest := client.genServerQuest("dataset")
	quest.Param("uid", uid)
	quest.Param("key", key)
	quest.Param("val", value)

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
func (client *RTMServerClient) DelData(uid int64, key string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DelData() function.")
		}
	}

	quest := client.genServerQuest("datadel")
	quest.Param("uid", uid)
	quest.Param("key", key)

	return client.sendSilentQuest(quest, timeout, callback)
}
