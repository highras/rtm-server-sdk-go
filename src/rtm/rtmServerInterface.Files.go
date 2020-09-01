package rtm

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

//-----------[ Files functions ]-------------------//

func (client *RTMServerClient) sendFileTokenQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(token string, endpoint string, errorCode int, errInfo string)) (string, string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantString("token"), answer.WantString("endpoint"), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback("", "", errorCode, "")
			} else {
				callback("", "", answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return "", "", err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return "", "", err
	} else if !answer.IsException() {
		return answer.WantString("token"), answer.WantString("endpoint"), nil
	} else {
		return "", "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

type fileTokenInfo struct {
	toUid   int64
	toUids  []int64
	groupId int64
	roomId  int64
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (token string, endpoint string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) fileToken(fromUid int64, cmd string, info *fileTokenInfo, rest ...interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			return "", "", errors.New("Invaild params when call RTMServerClient.fileToken() function.")
		}
	}

	quest := client.genServerQuest("filetoken")
	quest.Param("from", fromUid)
	quest.Param("cmd", cmd)

	switch cmd {
	case "sendfile":
		quest.Param("to", info.toUid)
	case "sendfiles":
		quest.Param("tos", info.toUids)
	case "sendroomfile":
		quest.Param("rid", info.roomId)
	case "sendgroupfile":
		quest.Param("gid", info.groupId)
	case "broadcastfile":
	default:
		return "", "", errors.New("Invaild 'cmd' value for RTMServerClient.fileToken() function.")
	}

	return client.sendFileTokenQuest(quest, timeout, callback)
}

func calculateFileSign(token string, fileContent []byte) string {

	//-- tolower(md5(tolower(toString(md5(filecontent))) + ":" + token))

	fileMd5 := md5.Sum(fileContent)
	fileMd5Str := fmt.Sprintf("%x", fileMd5)

	ctx := md5.New()
	io.WriteString(ctx, fileMd5Str)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, token)

	sign := fmt.Sprintf("%x", ctx.Sum(nil))

	return sign
}

func genFileAttrs(token string, fileContent []byte, filename string, extension string) (string, error) {

	sgin := calculateFileSign(token, fileContent)
	attrsMap := make(map[string]string)

	attrsMap["sign"] = sgin

	if len(filename) > 0 {
		if len(extension) > 0 {
			attrsMap["filename"] = filename
		} else {
			pos := strings.LastIndex(filename, ".")
			if pos > 0 && pos < len(filename)-1 {
				attrsMap["filename"] = filename[:pos]
				attrsMap["exit"] = filename[(pos + 1):]
			} else {
				attrsMap["filename"] = filename
			}
		}

	}
	if len(extension) > 0 {
		attrsMap["ext"] = extension
	}

	bytes, err := json.Marshal(attrsMap)

	return string(bytes), err
}

/*
	Params:
		rest: can be include following params:
			mtype int8
			extension string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ...interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case int8:
			mtype = value
		case int:
			mtype = int8(value)
		case string:
			extension = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendFile() function.")
		}
	}

	info := &fileTokenInfo{}
	info.toUid = toUid

	if callback == nil {
		token, endpoint, err := client.fileToken(fromUid, "sendfile", info, timeout)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		attrs, err := genFileAttrs(token, fileContent, filename, extension)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when building file attrs, error: %v", err)
		}

		quest := fpnn.NewQuest("sendfile")
		quest.Param("pid", client.pid)
		quest.Param("token", token)
		quest.Param("mtype", mtype)

		quest.Param("from", fromUid)
		quest.Param("to", toUid)
		quest.Param("mid", client.idGen.genMid())
		quest.Param("file", fileContent)
		quest.Param("attrs", attrs)

		fileClient := fpnn.NewTCPClient(endpoint)
		answer, err := fileClient.SendQuest(quest)

		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when send P2P file, error: %v", err)
		} else if !answer.IsException() {
			return answer.WantInt64("mtime"), nil
		} else {
			return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
		}

	} else {

		realCallback := func(token string, endpoint string, errorCode int, errInfo string) {

			attrs, err := genFileAttrs(token, fileContent, filename, extension)
			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when building file attrs, error: %v", err))
			}

			quest := fpnn.NewQuest("sendfile")
			quest.Param("pid", client.pid)
			quest.Param("token", token)
			quest.Param("mtype", mtype)

			quest.Param("from", fromUid)
			quest.Param("to", toUid)
			quest.Param("mid", client.idGen.genMid())
			quest.Param("file", fileContent)
			quest.Param("attrs", attrs)

			fileClient := fpnn.NewTCPClient(endpoint)
			err = fileClient.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_OK {
					callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
				} else if answer != nil {
					callback(0, errorCode, answer.WantString("ex"))
				} else {
					callback(0, errorCode, "")
				}
				go fileClient.Close()
			}, timeout)

			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when send P2P file, error: %v", err))
			}
		}

		_, _, err := client.fileToken(fromUid, "sendfile", info, timeout, realCallback)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		return 0, nil
	}
}

/*
	Params:
		rest: can be include following params:
			mtype int8
			extension string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ...interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case int8:
			mtype = value
		case int:
			mtype = int8(value)
		case string:
			extension = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendFiles() function.")
		}
	}

	info := &fileTokenInfo{}
	info.toUids = toUids

	if callback == nil {
		token, endpoint, err := client.fileToken(fromUid, "sendfiles", info, timeout)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		attrs, err := genFileAttrs(token, fileContent, filename, extension)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when building file attrs, error: %v", err)
		}

		quest := fpnn.NewQuest("sendfiles")
		quest.Param("pid", client.pid)
		quest.Param("token", token)
		quest.Param("mtype", mtype)

		quest.Param("from", fromUid)
		quest.Param("tos", toUids)
		quest.Param("mid", client.idGen.genMid())
		quest.Param("file", fileContent)
		quest.Param("attrs", attrs)

		fileClient := fpnn.NewTCPClient(endpoint)
		answer, err := fileClient.SendQuest(quest)

		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when send files, error: %v", err)
		} else if !answer.IsException() {
			return answer.WantInt64("mtime"), nil
		} else {
			return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
		}

	} else {

		realCallback := func(token string, endpoint string, errorCode int, errInfo string) {

			attrs, err := genFileAttrs(token, fileContent, filename, extension)
			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when building file attrs, error: %v", err))
			}

			quest := fpnn.NewQuest("sendfiles")
			quest.Param("pid", client.pid)
			quest.Param("token", token)
			quest.Param("mtype", mtype)

			quest.Param("from", fromUid)
			quest.Param("tos", toUids)
			quest.Param("mid", client.idGen.genMid())
			quest.Param("file", fileContent)
			quest.Param("attrs", attrs)

			fileClient := fpnn.NewTCPClient(endpoint)
			err = fileClient.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_OK {
					callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
				} else if answer != nil {
					callback(0, errorCode, answer.WantString("ex"))
				} else {
					callback(0, errorCode, "")
				}
				go fileClient.Close()
			}, timeout)

			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when send files, error: %v", err))
			}
		}

		_, _, err := client.fileToken(fromUid, "sendfiles", info, timeout, realCallback)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		return 0, nil
	}
}

/*
	Params:
		rest: can be include following params:
			mtype int8
			extension string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ...interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case int8:
			mtype = value
		case int:
			mtype = int8(value)
		case string:
			extension = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendGroupFile() function.")
		}
	}

	info := &fileTokenInfo{}
	info.groupId = groupId

	if callback == nil {
		token, endpoint, err := client.fileToken(fromUid, "sendgroupfile", info, timeout)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		attrs, err := genFileAttrs(token, fileContent, filename, extension)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when building file attrs, error: %v", err)
		}

		quest := fpnn.NewQuest("sendgroupfile")
		quest.Param("pid", client.pid)
		quest.Param("token", token)
		quest.Param("mtype", mtype)

		quest.Param("from", fromUid)
		quest.Param("gid", groupId)
		quest.Param("mid", client.idGen.genMid())
		quest.Param("file", fileContent)
		quest.Param("attrs", attrs)

		fileClient := fpnn.NewTCPClient(endpoint)
		answer, err := fileClient.SendQuest(quest)

		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when send group file, error: %v", err)
		} else if !answer.IsException() {
			return answer.WantInt64("mtime"), nil
		} else {
			return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
		}

	} else {

		realCallback := func(token string, endpoint string, errorCode int, errInfo string) {

			attrs, err := genFileAttrs(token, fileContent, filename, extension)
			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when building file attrs, error: %v", err))
			}

			quest := fpnn.NewQuest("sendgroupfile")
			quest.Param("pid", client.pid)
			quest.Param("token", token)
			quest.Param("mtype", mtype)

			quest.Param("from", fromUid)
			quest.Param("gid", groupId)
			quest.Param("mid", client.idGen.genMid())
			quest.Param("file", fileContent)
			quest.Param("attrs", attrs)

			fileClient := fpnn.NewTCPClient(endpoint)
			err = fileClient.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_OK {
					callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
				} else if answer != nil {
					callback(0, errorCode, answer.WantString("ex"))
				} else {
					callback(0, errorCode, "")
				}
				go fileClient.Close()
			}, timeout)

			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when send group file, error: %v", err))
			}
		}

		_, _, err := client.fileToken(fromUid, "sendgroupfile", info, timeout, realCallback)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		return 0, nil
	}
}

/*
	Params:
		rest: can be include following params:
			mtype int8
			extension string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ...interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case int8:
			mtype = value
		case int:
			mtype = int8(value)
		case string:
			extension = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendRoomFile() function.")
		}
	}

	info := &fileTokenInfo{}
	info.roomId = roomId

	if callback == nil {
		token, endpoint, err := client.fileToken(fromUid, "sendroomfile", info, timeout)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		attrs, err := genFileAttrs(token, fileContent, filename, extension)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when building file attrs, error: %v", err)
		}

		quest := fpnn.NewQuest("sendroomfile")
		quest.Param("pid", client.pid)
		quest.Param("token", token)
		quest.Param("mtype", mtype)

		quest.Param("from", fromUid)
		quest.Param("rid", roomId)
		quest.Param("mid", client.idGen.genMid())
		quest.Param("file", fileContent)
		quest.Param("attrs", attrs)

		fileClient := fpnn.NewTCPClient(endpoint)
		answer, err := fileClient.SendQuest(quest)

		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when send room file, error: %v", err)
		} else if !answer.IsException() {
			return answer.WantInt64("mtime"), nil
		} else {
			return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
		}

	} else {

		realCallback := func(token string, endpoint string, errorCode int, errInfo string) {

			attrs, err := genFileAttrs(token, fileContent, filename, extension)
			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when building file attrs, error: %v", err))
			}

			quest := fpnn.NewQuest("sendroomfile")
			quest.Param("pid", client.pid)
			quest.Param("token", token)
			quest.Param("mtype", mtype)

			quest.Param("from", fromUid)
			quest.Param("rid", roomId)
			quest.Param("mid", client.idGen.genMid())
			quest.Param("file", fileContent)
			quest.Param("attrs", attrs)

			fileClient := fpnn.NewTCPClient(endpoint)
			err = fileClient.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_OK {
					callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
				} else if answer != nil {
					callback(0, errorCode, answer.WantString("ex"))
				} else {
					callback(0, errorCode, "")
				}
				go fileClient.Close()
			}, timeout)

			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when send room file, error: %v", err))
			}
		}

		_, _, err := client.fileToken(fromUid, "sendroomfile", info, timeout, realCallback)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		return 0, nil
	}
}

/*
	Params:
		rest: can be include following params:
			mtype int8
			extension string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ...interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case int8:
			mtype = value
		case int:
			mtype = int8(value)
		case string:
			extension = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendBroadcastFile() function.")
		}
	}

	if callback == nil {
		token, endpoint, err := client.fileToken(fromUid, "broadcastfile", nil, timeout)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		attrs, err := genFileAttrs(token, fileContent, filename, extension)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when building file attrs, error: %v", err)
		}

		quest := fpnn.NewQuest("broadcastfile")
		quest.Param("pid", client.pid)
		quest.Param("token", token)
		quest.Param("mtype", mtype)

		quest.Param("from", fromUid)
		quest.Param("mid", client.idGen.genMid())
		quest.Param("file", fileContent)
		quest.Param("attrs", attrs)

		fileClient := fpnn.NewTCPClient(endpoint)
		answer, err := fileClient.SendQuest(quest)

		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when send braodcast file, error: %v", err)
		} else if !answer.IsException() {
			return answer.WantInt64("mtime"), nil
		} else {
			return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
		}

	} else {

		realCallback := func(token string, endpoint string, errorCode int, errInfo string) {

			attrs, err := genFileAttrs(token, fileContent, filename, extension)
			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when building file attrs, error: %v", err))
			}

			quest := fpnn.NewQuest("broadcastfile")
			quest.Param("pid", client.pid)
			quest.Param("token", token)
			quest.Param("mtype", mtype)

			quest.Param("from", fromUid)
			quest.Param("mid", client.idGen.genMid())
			quest.Param("file", fileContent)
			quest.Param("attrs", attrs)

			fileClient := fpnn.NewTCPClient(endpoint)
			err = fileClient.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_OK {
					callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
				} else if answer != nil {
					callback(0, errorCode, answer.WantString("ex"))
				} else {
					callback(0, errorCode, "")
				}
				go fileClient.Close()
			}, timeout)

			if err != nil {
				callback(0, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("[Exception] Exception when send braodcast file, error: %v", err))
			}
		}

		_, _, err := client.fileToken(fromUid, "broadcastfile", nil, timeout, realCallback)
		if err != nil {
			return 0, fmt.Errorf("[Exception] Exception when get file token, error: %v", err)
		}

		return 0, nil
	}
}
