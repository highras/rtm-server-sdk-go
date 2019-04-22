package rtm

import (
	"io"
	"os"
	"fmt"
	"log"
	"sync"
	"time"
	"strconv"
	"strings"
	"crypto/md5"
	"encoding/json"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

type RTMServerMonitor interface {
	P2PMessage(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	GroupMessage(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	RoomMessage(fromUid int64, roomIid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
}

//------------------------------[ RTM Server Client ]---------------------------------------//

type midGenerator struct {
	mutex			sync.Mutex
	idBase			int16
}

func (gen *midGenerator) genMid() int64 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()

	now := time.Now()
	mid := now.UnixNano() / 1000000

	gen.idBase += 1
	if gen.idBase > 999 {
		gen.idBase = 0
	}
		
	return mid * 1000 + int64(gen.idBase)
}

type RTMServerClient struct {
	client				*fpnn.TCPClient
	processor			*rtmServerQuestProcessor
	logger				*log.Logger
	idGen				*midGenerator
	pid					int32
	secretKey			string
}

func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient {

	client := &RTMServerClient{}
	
	client.client = fpnn.NewTCPClient(endpoint)
	client.processor = newRTMServerQuestProcessor()
	client.idGen = &midGenerator{}
	client.secretKey = secretKey
	client.pid = pid

	client.client.SetQuestProcessor(client.processor)
	client.SetLogger(log.New(os.Stdout, "[RTM Go SDK] ", log.LstdFlags))

	return client
}

//------------------------------[ RTM Server Client Config Interfaces ]---------------------------------------//

func (client *RTMServerClient) SetMonitor(monitor RTMServerMonitor) {
	client.processor.monitor = monitor
}

func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration) {
	client.client.SetConnectTimeOut(timeout)
}

func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration) {
	client.client.SetQuestTimeOut(timeout)
}

func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64)) {
	client.client.SetOnConnectedCallback(onConnected)
}

func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64)) {
	client.client.SetOnClosedCallback(onClosed)
}

func (client *RTMServerClient) SetLogger(logger *log.Logger) {
	client.logger = logger
	client.processor.logger = logger
	client.client.SetLogger(logger)
}

func (client *RTMServerClient) Endpoint() string {
	return client.client.Endpoint()
}

func (client *RTMServerClient) makeSignAndSalt() (string, int64) {

	now := time.Now()
	salt := now.UnixNano()

	pidStr := strconv.FormatInt(int64(client.pid), 10)
	saltStr := strconv.FormatInt(salt, 10)


	ctx := md5.New()
	io.WriteString(ctx, pidStr)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, client.secretKey)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, saltStr)

	sign := fmt.Sprintf("%X", ctx.Sum(nil))

	return sign, salt
}

//------------------------------[ Utilities Functions ]---------------------------------------//
func convertToInt64(value interface{}) int64 {
	switch value.(type) {
	case int64:
		return int64(value.(int64))
	case int32:
		return int64(value.(int32))
	case int16:
		return int64(value.(int16))
	case int8:
		return int64(value.(int8))
	case int:
		return int64(value.(int))

	case uint64:
		return int64(value.(uint64))
	case uint32:
		return int64(value.(uint32))
	case uint16:
		return int64(value.(uint16))
	case uint8:
		return int64(value.(uint8))
	case uint:
		return int64(value.(uint))

	case float32:
		return int64(value.(float32))
	case float64:
		return int64(value.(float64))

	default:
		panic("Type convert failed.")
	}
}

func convertToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case []byte:
		return string(value.([]byte))
	case []rune:
		return string(value.([]rune))
	default:
		panic("Type convert failed.")
	}
}

//------------------------------[ RTM Server Client Interfaces ]---------------------------------------//
func (client *RTMServerClient) sendQuest(quest *fpnn.Quest, timeout time.Duration, callback func(answer *fpnn.Answer, errorCode int)) (*fpnn.Answer, error) {
	
	if callback == nil {
		if timeout == 0 {
			return client.client.SendQuest(quest)
		} else {
			return client.client.SendQuestWithTimeout(quest, timeout)
		}

	} else {
		if timeout == 0 {
			return nil, client.client.SendQuestWithLambda(quest, callback)
		} else {
			return nil, client.client.SendQuestWithLambdaTimeout(quest, callback, timeout)
		}
	}
}

func (client *RTMServerClient) sendOkQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(ok bool, errorCode int, errInfo string)) (bool, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantBool("ok"), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(false, errorCode, "")
			} else {
				callback(false, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return true, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return false, err
	} else if !answer.IsException() {
		return answer.WantBool("ok"), nil
	} else {
		return false, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func convertSliceToInt64Slice(slice []interface{}) []int64 {
	if slice == nil || len(slice) == 0 {
		return make([]int64, 0, 1)
	}

	rev := make([]int64, 0, len(slice))
	for _, elem := range slice {
		val := convertToInt64(elem)
		rev = append(rev, val)
	}
	return rev
}

func (client *RTMServerClient) sendSliceQuest(quest *fpnn.Quest, timeout time.Duration,
	sliceKey string, callback func(slice []int64, errorCode int, errInfo string)) ([]int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(convertSliceToInt64Slice(answer.WantSlice(sliceKey)), fpnn.FPNN_EC_OK, "")
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
		slice := convertSliceToInt64Slice(answer.WantSlice(sliceKey))
		return slice, nil
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendSilentQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(errorCode int, errInfo string)) error {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(errorCode, "")
			} else {
				callback(answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return err
	} else if !answer.IsException() {
		return nil
	} else {
		return fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

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
func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func (int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				attrs = value
			case time.Duration:
				timeout = value
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("sendmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ... interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func (int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				attrs = value
			case time.Duration:
				timeout = value
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendMessages() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("sendmsgs")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ... interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func (int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				attrs = value
			case time.Duration:
				timeout = value
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendGroupMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("sendgroupmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ... interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func (int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				attrs = value
			case time.Duration:
				timeout = value
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendRoomMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("sendroommsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) SendBoradcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error) {

	var attrs string
	var timeout time.Duration
	var callback func (int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				attrs = value
			case time.Duration:
				timeout = value
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendBoradcastMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("broadcastmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("mtype", mtype)

	quest.Param("from", fromUid)
	quest.Param("mid", client.idGen.genMid())
	quest.Param("msg", message)
	quest.Param("attrs", attrs)

	return client.sendMessageQuest(quest, timeout, callback)
}

//-----------[ History Messages functions ]-------------------//

type HistoryMessageUnit struct {
	Id			int64
	FromUid		int64
	MType		int8
	Mid			int64
	Deleted		bool
	Message		string
	Attrs		string
	MTime		int64
}

type HistoryMessageResult struct {
	Num			int16
	LastId		int64
	Begin		int64
	End			int64
	Messages	[]*HistoryMessageUnit
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

		msgUnit.Deleted = elems[4].(bool)
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
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error) {

	var timeout time.Duration
	var callback func (*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (*HistoryMessageResult, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetGroupMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getgroupmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error) {

	var timeout time.Duration
	var callback func (*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (*HistoryMessageResult, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetRoomMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getroommsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16,
	begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error) {

	var timeout time.Duration
	var callback func (*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (*HistoryMessageResult, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetBroadcastMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getbroadcastmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

	return client.sendHistoryMessageQuest(quest, timeout, nil, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error) {

	var timeout time.Duration
	var callback func (*HistoryMessageResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (*HistoryMessageResult, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetP2PMessage() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getp2pmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("ouid", peerUid)

	quest.Param("desc", desc)
	quest.Param("num", num)
	quest.Param("begin", begin)
	quest.Param("end", end)
	quest.Param("lastid", lastid)

	return client.sendHistoryMessageQuest(quest, timeout, []int64{uid, peerUid}, callback)
}

//-----------[ Friends functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("friends", firendUids)

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
func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("friends", firendUids)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsFriend() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isfriend")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("fuid", peerUid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (firendUids []int64, err error)
*/
func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsFriends() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isfriends")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)
	quest.Param("fuids", uids)

	return client.sendSliceQuest(quest, timeout, "fuids", callback)
}

//-----------[ Group functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uids", uids)

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
func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uids", uids)

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
func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelGroup() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delgroup")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsGroupMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isgroupmember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (groupIds []int64, err error)
*/
func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetUserGroups() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getusergroups")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendSliceQuest(quest, timeout, "gids", callback)
}

//-----------[ Room functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddRoomMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addroommember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelRoomMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delroommember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}

//-----------[ Manage functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddGroupBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addgroupban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
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
func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveGroupBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removegroupban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
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
func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddRoomBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addroomban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
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
func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveRoomBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removeroomban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddProjectBlack() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addprojectblack")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveProjectBlack() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removeprojectblack")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsBanOfGroup() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isbanofgroup")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsBanOfRoom() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isbanofroom")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("rid", roomId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsProjectBlack(uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsProjectBlack() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isprojectblack")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

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
func (client *RTMServerClient) DelMessage(mid int64, fromUid int64, xid int64, messageType MessageType, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
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

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delmsg")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("mid", mid)
	quest.Param("from", fromUid)
	quest.Param("xid", xid)
	quest.Param("type", realType)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			clientEndpoint string
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) Kickout(uid int64, rest ... interface{}) error {

	var clientEndpoint string
	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case string:
				clientEndpoint = value
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.Kickout() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("kickout")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("uid", uid)
	quest.Param("ce", clientEndpoint)

	return client.sendSilentQuest(quest, timeout, callback)
}

//-----------[ Utilities functions ]-------------------//

func (client *RTMServerClient) sendTokenQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(token string, errorCode int, errInfo string)) (string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantString("token"), fpnn.FPNN_EC_OK, "")
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
		return answer.WantString("token"), nil
	} else {
		return "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (token string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", error);
		else this function work in sync mode, and return (token string, err error)
*/
func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error) {

	var timeout time.Duration
	var callback func (string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (string, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetToken() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("gettoken")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendTokenQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetOnlineUsers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getonlineusers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
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
func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddListen() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addlisten")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	if groupIds != nil && len(groupIds) > 0 {
		quest.Param("gids", groupIds)	
	}
	if roomIds != nil && len(roomIds) > 0 {
		quest.Param("rids", roomIds)	
	}
	if p2p {
		quest.Param("p2p", true)	
	}
	if events != nil && len(events) > 0 {
		quest.Param("events", events)	
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
func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveListen() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removelisten")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	if groupIds != nil && len(groupIds) > 0 {
		quest.Param("gids", groupIds)	
	}
	if roomIds != nil && len(roomIds) > 0 {
		quest.Param("rids", roomIds)	
	}
	if p2p {
		quest.Param("p2p", true)	
	}
	if events != nil && len(events) > 0 {
		quest.Param("events", events)	
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			all bool
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error {

	var all *bool
	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case bool:
				all = new(bool)
				*all = value
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SetListen() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("setlisten")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("gids", groupIds)
	quest.Param("rids", roomIds)	
	quest.Param("p2p", p2p)	
	quest.Param("events", events)
	
	if all != nil {
		quest.Param("all", *all)
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
func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddDevice() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("adddevice")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

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
func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveDevice() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removedevice")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("uid", uid)
	quest.Param("devicetoken", deviceToken)

	return client.sendSilentQuest(quest, timeout, callback)
}

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
	toUid int64
	toUids []int64
	groupId int64
	roomId int64
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (token string, endpoint string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) fileToken(fromUid int64, cmd string, info *fileTokenInfo, rest ... interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func (string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (string, string, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.fileToken() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("filetoken")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

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
		panic("Invaild 'cmd' value for RTMServerClient.fileToken() function.")
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
			if pos > 0 && pos < len(filename) - 1 {
				attrsMap["filename"] = filename[:pos]
				attrsMap["exit"] = filename[(pos+1):]
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
func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func (int64, int, string)

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
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendFile() function.")
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

		realCallback := func (token string, endpoint string, errorCode int, errInfo string) {
			
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
			err = fileClient.SendQuestWithLambdaTimeout(quest, func(answer *fpnn.Answer, errorCode int) {
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
func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ... interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func (int64, int, string)

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
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendFiles() function.")
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

		realCallback := func (token string, endpoint string, errorCode int, errInfo string) {
			
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
			err = fileClient.SendQuestWithLambdaTimeout(quest, func(answer *fpnn.Answer, errorCode int) {
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
func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func (int64, int, string)

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
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendGroupFile() function.")
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

		realCallback := func (token string, endpoint string, errorCode int, errInfo string) {
			
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
			err = fileClient.SendQuestWithLambdaTimeout(quest, func(answer *fpnn.Answer, errorCode int) {
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
func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func (int64, int, string)

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
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendRoomFile() function.")
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

		realCallback := func (token string, endpoint string, errorCode int, errInfo string) {
			
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
			err = fileClient.SendQuestWithLambdaTimeout(quest, func(answer *fpnn.Answer, errorCode int) {
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
func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error) {

	var mtype int8 = 50
	var extension string
	var timeout time.Duration
	var callback func (int64, int, string)

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
			case func (int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SendBroadcastFile() function.")
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

		realCallback := func (token string, endpoint string, errorCode int, errInfo string) {
			
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
			err = fileClient.SendQuestWithLambdaTimeout(quest, func(answer *fpnn.Answer, errorCode int) {
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
