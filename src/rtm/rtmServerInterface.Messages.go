package rtm

import (
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
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
func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ...interface{}) (int64, error) {

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
			panic("Invaild params when call RTMServerClient.SendMessage() function.")
		}
	}

	quest := client.genServerQuest("sendmsg")
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("to", toUid)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
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
func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ...interface{}) (int64, error) {

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
			panic("Invaild params when call RTMServerClient.SendMessages() function.")
		}
	}

	quest := client.genServerQuest("sendmsgs")
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("tos", toUids)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
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
func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ...interface{}) (int64, error) {

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
			panic("Invaild params when call RTMServerClient.SendGroupMessage() function.")
		}
	}

	quest := client.genServerQuest("sendgroupmsg")
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("gid", groupId)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
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
func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ...interface{}) (int64, error) {

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
			panic("Invaild params when call RTMServerClient.SendRoomMessage() function.")
		}
	}

	quest := client.genServerQuest("sendroommsg")
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("rid", roomId)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
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
func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, mtype int8, message string, rest ...interface{}) (int64, error) {

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
			panic("Invaild params when call RTMServerClient.SendBroadcastMessage() function.")
		}
	}

	quest := client.genServerQuest("broadcastmsg")
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

//-----------[ History Messages functions ]-------------------//

type HistoryMessageUnit struct {
	Id      int64
	FromUid int64
	MType   int8
	Mid     int64
	//Deleted bool
	Message string
	Attrs   string
	MTime   int64
}

type HistoryMessageResult struct {
	Num      int16
	LastId   int64
	Begin    int64
	End      int64
	Messages []*HistoryMessageUnit
}

func (client *RTMServerClient) processHistoryAnswer(answer *fpnn.Answer, p2pInfo []int64) (res *HistoryMessageResult, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ERROR] Process history messages exception. Panic: %v.", r)
		}
	}()

	result := &HistoryMessageResult{}
	result.Num = answer.WantInt16("num")
	result.LastId = answer.WantInt64("lastid")
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

		msgUnit.Id = convertToInt64(elems[0])
		msgUnit.FromUid = convertToInt64(elems[1])
		msgUnit.MType = int8(convertToInt64(elems[2]))
		msgUnit.Mid = convertToInt64(elems[3])

		//msgUnit.Deleted = elems[4].(bool)
		msgUnit.Message = convertToString(elems[5])
		msgUnit.Attrs = convertToString(elems[6])
		msgUnit.MTime = convertToInt64(elems[7])

		if p2pInfo != nil {
			if msgUnit.FromUid == 1 {
				msgUnit.FromUid = p2pInfo[0]
			} else if msgUnit.FromUid == 2 {
				msgUnit.FromUid = p2pInfo[1]
			} else {
				client.logger.Printf("[ERROR] Unknown P2P history message direction %d", msgUnit.FromUid)
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
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

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
			panic("Invaild params when call RTMServerClient.GetGroupMessage() function.")
		}
	}

	quest := client.genServerQuest("getgroupmsg")
	quest.Param("gid", groupId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

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
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

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
			panic("Invaild params when call RTMServerClient.GetRoomMessage() function.")
		}
	}

	quest := client.genServerQuest("getroommsg")
	quest.Param("rid", roomId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

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
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

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
			panic("Invaild params when call RTMServerClient.GetBroadcastMessage() function.")
		}
	}

	quest := client.genServerQuest("getbroadcastmsg")

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

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
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

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
			panic("Invaild params when call RTMServerClient.GetP2PMessage() function.")
		}
	}

	quest := client.genServerQuest("getp2pmsg")
	quest.Param("uid", uid)
	quest.Param("ouid", peerUid)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

	if mtypes != nil {
		quest.Param("mtypes", mtypes)
	}

	return client.sendHistoryMessageQuest(quest, timeout, []int64{uid, peerUid}, callback)
}

//-----------[ Delete Messages functions ]-------------------//

type MessageType int

const (
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
func (client *RTMServerClient) DelMessage(mid int64, fromUid int64, xid int64, messageType MessageType, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.DelMessage() function.")
		}
	}

	var realType int8
	switch messageType {
	case MessageType_P2P:
		realType = 1
	case MessageType_Group:
		realType = 2
	case MessageType_Room:
		realType = 3
	case MessageType_Broadcast:
		realType = 4
	default:
		panic("Invaild messageType when call RTMServerClient.DelMessage() function.")
	}

	quest := client.genServerQuest("delmsg")

	quest.Param("mid", mid)
	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", realType)

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
func (client *RTMServerClient) DelP2PMessage(mid int64, fromUid int64, xid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, xid, MessageType_P2P, rest)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelGroupMessage(mid int64, fromUid int64, xid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, xid, MessageType_Group, rest)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelRoomMessage(mid int64, fromUid int64, xid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, xid, MessageType_Room, rest)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelBroadcastMessage(mid int64, fromUid int64, xid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, xid, MessageType_Broadcast, rest)
}
