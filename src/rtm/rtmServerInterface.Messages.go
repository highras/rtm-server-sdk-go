package rtm

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

type ClearType int8

const (
	P2P ClearType = iota
	Room
	Group
	Broadcast
	All
)

func (client *RTMServerClient) sendMessageQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(mtime int64, errorCode int, errInfo string)) (int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(0, errorCode, "")
			} else {
				callback(0, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return 0, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return 0, err
	} else if !answer.IsException() {
		return answer.WantInt64("mtime"), nil
	} else {
		return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

//-----------[ Message functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			attrs string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, messageType int8, message string, rest ...interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			attrs = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendMessage() function.")
		}
	}

	quest := client.genServerQuest("sendmsg")
	quest.Param("mtype", messageType)

	quest.Param("from", fromUid)
	quest.Param("to", toUid)
	quest.Param("mid", idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

func (client *RTMServerClient) SendMessageByBinary(fromUid int64, toUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error) {
	return client.SendMessage(fromUid, toUid, messageType, string(message), rest...)
}

/*
Params:

	rest: can be include following params:
		attrs string
		timeout time.Duration
		func (mtime int64, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (0, error);
	else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, messageType int8, message string, rest ...interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			attrs = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendMessages() function.")
		}
	}

	quest := client.genServerQuest("sendmsgs")
	quest.Param("mtype", messageType)

	quest.Param("from", fromUid)
	quest.Param("tos", toUids)
	quest.Param("mid", idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

func (client *RTMServerClient) SendMessagesByBinary(fromUid int64, toUids []int64, messageType int8, message []byte, rest ...interface{}) (int64, error) {
	return client.SendMessages(fromUid, toUids, messageType, string(message), rest...)
}

/*
Params:

	rest: can be include following params:
		attrs string
		timeout time.Duration
		func (mtime int64, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (0, error);
	else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, messageType int8, message string, rest ...interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			attrs = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendGroupMessage() function.")
		}
	}

	quest := client.genServerQuest("sendgroupmsg")
	quest.Param("mtype", messageType)

	quest.Param("from", fromUid)
	quest.Param("gid", groupId)
	quest.Param("mid", idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

func (client *RTMServerClient) SendGroupMessageByBinary(fromUid int64, groupId int64, messageType int8, message []byte, rest ...interface{}) (int64, error) {
	return client.SendGroupMessage(fromUid, groupId, messageType, string(message), rest...)
}

/*
Params:

	rest: can be include following params:
		attrs string
		timeout time.Duration
		func (mtime int64, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (0, error);
	else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, messageType int8, message string, rest ...interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			attrs = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendRoomMessage() function.")
		}
	}

	quest := client.genServerQuest("sendroommsg")
	quest.Param("mtype", messageType)

	quest.Param("from", fromUid)
	quest.Param("rid", roomId)
	quest.Param("mid", idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

func (client *RTMServerClient) SendRoomMessageByBinary(fromUid int64, roomId int64, messageType int8, message []byte, rest ...interface{}) (int64, error) {
	return client.SendRoomMessage(fromUid, roomId, messageType, string(message), rest...)
}

/*
Params:

	rest: can be include following params:
		attrs string
		timeout time.Duration
		func (mtime int64, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (0, error);
	else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, messageType int8, message string, rest ...interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func(int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			attrs = value
		case time.Duration:
			timeout = value
		case func(int64, int, string):
			callback = value
		default:
			return 0, errors.New("Invaild params when call RTMServerClient.SendBroadcastMessage() function.")
		}
	}

	quest := client.genServerQuest("broadcastmsg")
	quest.Param("mtype", messageType)

	quest.Param("from", fromUid)
	quest.Param("mid", idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

func (client *RTMServerClient) SendBroadcastMessageByBinary(fromUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error) {
	return client.SendBroadcastMessage(fromUid, messageType, string(message), rest...)
}

//-----------[ History Messages functions ]-------------------//
// old version
// type HistoryMessageUnit struct {
// 	Id      int64
// 	FromUid int64
// 	MType   int8
// 	Mid     int64
// 	//Deleted bool
// 	Message string
// 	Attrs   string
// 	MTime   int64
// }

// type HistoryMessageResult struct {
// 	Num      int16
// 	LastId   int64
// 	Begin    int64
// 	End      int64
// 	Messages []*HistoryMessageUnit
// }

// new version

type RTMAudioFileInfo struct {
	IsRTMaudio bool   // 是否是rtm语音消息
	Codec      string // rtm语音消息时有此值编 码格式
	Srate      int32  // rtm语音消息时有此值 采样率
	Lang       string // 如果是rtm语音会有此值
	Duration   int32  // ms，如果是rtm语音会有此值
}

type FileMsgInfo struct {
	Url      string `json:"url"`
	FileSize int64  `json:"size"` // 字节大小
	Surl     string `json:"surl"` // 缩略图的地址，如果是图片类型会有此值
	RTMAudioFileInfo
}

type RTMMessage struct {
	FromUid     int64
	ToId        int64
	MessageType int8
	MessageId   int64
	//Deleted bool
	Message      string
	Attrs        string
	ModifiedTime int64
	FileInfo     *FileMsgInfo
}

type HistoryMessageUnit struct {
	CursorId int64
	RTMMessage
}

type HistoryMessageResult struct {
	Num          int16
	LastCursorId int64
	Begin        int64
	End          int64
	Messages     []*HistoryMessageUnit
}

func processFileInfo(msg string, attrs string, mtype int8, logger RTMLogger) *FileMsgInfo {
	fileInfo := &FileMsgInfo{}
	err1 := json.Unmarshal(([]byte)(msg), fileInfo)
	if err1 != nil {
		logger.Printf("parse json error for get file msg, file msg := %s, err := %v.\n", msg, err1)
		return fileInfo
	}
	if mtype == defaultMtype_Audio {
		data := make(map[string]interface{})
		if err2 := json.Unmarshal(([]byte)(attrs), &data); err2 != nil {
			logger.Printf("parse json error for get file msg, attrs := %s, err := %v.\n", attrs, err2)
			return fileInfo
		}
		value, ok := data["rtm"]
		if !ok {
			logger.Printf("parse json error for get file msg, attrs not have rtm key, attrs := %s.\n", attrs)
			return fileInfo
		}
		rtmdata, ok1 := value.(map[string]interface{})
		if !ok1 {
			logger.Printf("parse json error for get file msg, attrs rtm key-value invalid, attrs := %s.\n", attrs)
			return fileInfo
		}

		if typeValue, ok2 := rtmdata["type"]; ok2 {
			if typeString, ok3 := typeValue.(string); ok3 {
				if typeString == "audiomsg" {
					if lang, ok4 := rtmdata["lang"]; ok4 {
						if realLang, ok5 := lang.(string); ok5 {
							fileInfo.Lang = realLang
						}
					}
					if duration, ok6 := rtmdata["duration"]; ok6 {
						if realDuration, ok7 := duration.(int32); ok7 {
							fileInfo.Duration = realDuration
						}
					}
					if rate, ok7 := rtmdata["srate"]; ok7 {
						if realRate, ok8 := rate.(int32); ok8 {
							fileInfo.Srate = realRate
						}
					}
					if codec, ok9 := rtmdata["codec"]; ok9 {
						if realCodec, ok10 := codec.(string); ok10 {
							fileInfo.Codec = realCodec
						}
					}

					fileInfo.IsRTMaudio = true
				}
			}

		}
	}
	return fileInfo
}

func fetchFileCustomAttrs(attrs string, logger RTMLogger) string {
	realAttrs := make(map[string]interface{})
	if err := json.Unmarshal(([]byte)(attrs), &realAttrs); err != nil {
		logger.Printf("parse file custom attrs error, attrs := %s, err := %v.\n", attrs, err)
	} else {
		if value, ok := realAttrs["custom"]; ok {
			if cAttrs, ok1 := value.(map[string]interface{}); ok1 {
				if attrJson, err := json.Marshal(cAttrs); err == nil {
					return (string)(attrJson)
				}
			} else if customString, ok2 := value.(string); ok2 {
				return customString
			}
		}
		logger.Printf("parse file custom attrs error, attrs := %s.\n", attrs)
	}
	return ""
}

func (client *RTMServerClient) processHistoryAnswer(answer *fpnn.Answer, p2pInfo []int64) (res *HistoryMessageResult, err error) {

	result := &HistoryMessageResult{}
	result.Num = answer.WantInt16("num")
	result.LastCursorId = answer.WantInt64("lastid")
	result.Begin = answer.WantInt64("begin")
	result.End = answer.WantInt64("end")

	if result.Num == 0 {
		return result, err
	}

	result.Messages = make([]*HistoryMessageUnit, 0, result.Num)

	messages := answer.WantSlice("msgs")
	for _, unit := range messages {
		elems := unit.([]interface{})
		msgUnit := &HistoryMessageUnit{}

		msgUnit.CursorId = client.convertToInt64(elems[0])
		msgUnit.FromUid = client.convertToInt64(elems[1])
		msgUnit.MessageType = int8(client.convertToInt64(elems[2]))
		msgUnit.MessageId = client.convertToInt64(elems[3])
		msgUnit.Attrs = client.convertToString(elems[6])
		//msgUnit.Deleted = elems[4].(bool)
		if msgUnit.MessageType >= defaultMtype_Image && msgUnit.MessageType <= defaultMtype_File {
			msg := client.convertToString(elems[5])
			msgUnit.Message = msg
			fileInfo := processFileInfo(msg, msgUnit.Attrs, msgUnit.MessageType, client.logger)
			msgUnit.FileInfo = fileInfo
			msgUnit.Attrs = fetchFileCustomAttrs(msgUnit.Attrs, client.logger)

		} else {
			msgUnit.Message = client.convertToString(elems[5])
		}
		msgUnit.ModifiedTime = client.convertToInt64(elems[7])

		if p2pInfo != nil {
			if msgUnit.FromUid == 1 {
				msgUnit.FromUid = p2pInfo[0]
			} else if msgUnit.FromUid == 2 {
				msgUnit.FromUid = p2pInfo[1]
			} else {
				client.logger.Printf("[ERROR] Unknown P2P history message direction %d.\n", msgUnit.FromUid)
			}
		}
		result.Messages = append(result.Messages, msgUnit)
	}

	return result, err
}

func (client *RTMServerClient) sendHistoryMessageQuest(quest *fpnn.Quest, timeout time.Duration, p2pInfo []int64,
	callback func(result *HistoryMessageResult, errorCode int, errInfo string)) (*HistoryMessageResult, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				if result, err := client.processHistoryAnswer(answer, p2pInfo); err == nil {
					callback(result, fpnn.FPNN_EC_OK, "")
				} else {
					callback(result, fpnn.FPNN_EC_CORE_UNKNOWN_ERROR, fmt.Sprintf("%v", err))
				}

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
		return client.processHistoryAnswer(answer, p2pInfo)
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
Params:

	rest: can be include following params:
		mtypes []int8
		timeout time.Duration
		func (result *HistoryMessageResult, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (nil, error);
	else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16,
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	var mtypes []int8
	var timeout time.Duration
	var callback func(*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case []int8:
			mtypes = value
		case time.Duration:
			timeout = value
		case func(*HistoryMessageResult, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetGroupMessage() function.")
		}
	}

	quest := client.genServerQuest("getgroupmsg")
	quest.Param("gid", groupId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastCursorId)
	quest.Param("uid", uid)

	if mtypes != nil {
		quest.Param("mtypes", mtypes)
	}

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
Params:

	rest: can be include following params:
		mtypes []int8
		timeout time.Duration
		func (result *HistoryMessageResult, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (nil, error);
	else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16,
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	var mtypes []int8
	var timeout time.Duration
	var callback func(*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case []int8:
			mtypes = value
		case time.Duration:
			timeout = value
		case func(*HistoryMessageResult, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetRoomMessage() function.")
		}
	}

	quest := client.genServerQuest("getroommsg")
	quest.Param("rid", roomId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastCursorId)
	quest.Param("uid", uid)

	if mtypes != nil {
		quest.Param("mtypes", mtypes)
	}

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
Params:

	rest: can be include following params:
		mtypes []int8
		timeout time.Duration
		func (result *HistoryMessageResult, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (nil, error);
	else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16,
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	var mtypes []int8
	var timeout time.Duration
	var callback func(*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case []int8:
			mtypes = value
		case time.Duration:
			timeout = value
		case func(*HistoryMessageResult, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetBroadcastMessage() function.")
		}
	}

	quest := client.genServerQuest("getbroadcastmsg")

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastCursorId)
	quest.Param("uid", uid)

	if mtypes != nil {
		quest.Param("mtypes", mtypes)
	}

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
Params:

	rest: can be include following params:
		mtypes []int8
		timeout time.Duration
		func (result *HistoryMessageResult, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (nil, error);
	else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16,
	begin int64, end int64, lastCursorId int64, rest ...interface{}) (*HistoryMessageResult, error) {

	var mtypes []int8
	var timeout time.Duration
	var callback func(*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case []int8:
			mtypes = value
		case time.Duration:
			timeout = value
		case func(*HistoryMessageResult, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetP2PMessage() function.")
		}
	}

	quest := client.genServerQuest("getp2pmsg")
	quest.Param("uid", uid)
	quest.Param("ouid", peerUid)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastCursorId)

	if mtypes != nil {
		quest.Param("mtypes", mtypes)
	}

	return client.sendHistoryMessageQuest(quest, timeout, []int64{uid, peerUid}, callback)
}

//-----------[ Delete Messages functions ]-------------------//

type MessageType int8

const (
	_                           = iota
	MessageType_P2P MessageType = iota
	MessageType_Group
	MessageType_Room
	MessageType_Broadcast
)

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelMessage(messageId int64, fromUid int64, xid int64, messageType MessageType, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DelMessage() function.")
		}
	}

	quest := client.genServerQuest("delmsg")

	quest.Param("mid", messageId)
	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", messageType)

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
func (client *RTMServerClient) CleanMessage(fromUid int64, xid int64, messageType MessageType, begin int64, end int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.CleanMessage() function.")
		}
	}

	quest := client.genServerQuest("cleanmsg")

	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", messageType)
	quest.Param("begin", begin)
	quest.Param("end", end)

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
func (client *RTMServerClient) ClearProjectMessage(clearType ClearType, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.CleanMessage() function.")
		}
	}

	quest := client.genServerQuest("clearprojectmsg")

	quest.Param("type", clearType)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (result *HistoryMessageUnit, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) GetMessage(messageId int64, fromUid int64, xid int64, messageType MessageType, rest ...interface{}) (*HistoryMessageUnit, error) {

	var timeout time.Duration
	var callback func(*HistoryMessageUnit, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(*HistoryMessageUnit, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetMessage() function.")
		}
	}

	quest := client.genServerQuest("getmsg")

	quest.Param("mid", messageId)
	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", messageType)

	return client.sendGetMsgInfoQuest(quest, timeout, callback)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) CleanP2PMessage(fromUid int64, to int64, begin int64, end int64, rest ...interface{}) error {
	return client.CleanMessage(fromUid, to, MessageType_P2P, begin, end, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) CleanGroupMessage(gid int64, begin int64, end int64, rest ...interface{}) error {
	return client.CleanMessage(0, gid, MessageType_Group, begin, end, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) CleanRoomMessage(rid int64, begin int64, end int64, rest ...interface{}) error {
	return client.CleanMessage(0, rid, MessageType_Room, begin, end, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) CleanBroadcastMessage(fromUid int64, begin int64, end int64, rest ...interface{}) error {
	return client.CleanMessage(fromUid, 0, MessageType_Broadcast, begin, end, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelP2PMessage(messageId int64, fromUid int64, to int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, to, MessageType_P2P, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelGroupMessage(messageId int64, fromUid int64, gid int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, gid, MessageType_Group, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelRoomMessage(messageId int64, fromUid int64, rid int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, rid, MessageType_Room, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelBroadcastMessage(messageId int64, fromUid int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, 0, MessageType_Broadcast, rest...)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (sender int32, count int32, errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (sender int32, count int32, err error)
*/
func (client *RTMServerClient) GetMsgCount(msgType MessageType, xid int64, begin int64, end int64, mtype []int8, rest ...interface{}) (sender int32, count int32, err error) {
	var timeout time.Duration
	var callback func(int32, int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int32, int, string):
			callback = value
		default:
			return 0, 0, errors.New("Invaild params when call RTMServerClient.GetMsgCount() function.")
		}
	}

	quest := client.genServerQuest("getmsgnum")
	quest.Param("type", msgType)
	quest.Param("xid", xid)
	if mtype != nil && len(mtype) > 0 {
		quest.Param("mtypes", mtype)
	}
	if begin > 0 {
		quest.Param("begin", begin)
	}
	if end > 0 {
		quest.Param("end", end)
	}

	return client.sendDoubleIntQuest(quest, timeout, "sender", "num", callback)
}

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DeleteConversationMessages(fromUid int64, xid int64, messageType MessageType, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.CleanMessage() function.")
		}
	}

	quest := client.genServerQuest("delconversationmsgs")

	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", messageType)

	return client.sendSilentQuest(quest, timeout, callback)
}

//-----------[ Edit Messages functions ]-------------------//

/*
Params:

	rest: can be include following params:
		timeout time.Duration
		func (errorCode int, errInfo string)

	If include func param, this function will enter into async mode, and return (error);
	else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) EditMessage(messageId int64, fromUid int64, xid int64, messageType MessageType, msg string, attrs string, timeLimit int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.EditMessage() function.")
		}
	}

	quest := client.genServerQuest("editmsg")

	quest.Param("mid", messageId)
	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", messageType)
	quest.Param("msg", msg)
	quest.Param("attrs", attrs)
	quest.Param("timeLimit", timeLimit)

	return client.sendSilentQuest(quest, timeout, callback)
}
