package rtm

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	SDKVersion = "0.3.1"
)

const (
	APIVersion = "2.1.0"
)

type RTMServerMonitor interface {
	P2PMessage(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	GroupMessage(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	RoomMessage(fromUid int64, roomIid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
	P2PChat(fromUid int64, toUid int64, mid int64, message string, attrs string, mtime int64)
	GroupChat(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64)
	RoomChat(fromUid int64, roomIid int64, mid int64, message string, attrs string, mtime int64)
	P2PAudio(fromUid int64, toUid int64, mid int64, message []byte, attrs string, mtime int64)
	GroupAudio(fromUid int64, groupId int64, mid int64, message []byte, attrs string, mtime int64)
	RoomAudio(fromUid int64, roomIid int64, mid int64, message []byte, attrs string, mtime int64)
	P2PCmd(fromUid int64, toUid int64, mid int64, message string, attrs string, mtime int64)
	GroupCmd(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64)
	RoomCmd(fromUid int64, roomIid int64, mid int64, message string, attrs string, mtime int64)
	P2PFile(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	GroupFile(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	RoomFile(fromUid int64, roomId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
}

//------------------------------[ RTM Server Client ]---------------------------------------//

type midGenerator struct {
	mutex  sync.Mutex
	idBase int16
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

	return mid*1000 + int64(gen.idBase)
}

type RTMServerClient struct {
	client    *fpnn.TCPClient
	processor *rtmServerQuestProcessor
	logger    *log.Logger
	idGen     *midGenerator
	pid       int32
	secretKey string
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

func (client *RTMServerClient) SetAutoReconnect(autoReconnect bool) {
	client.client.SetAutoReconnect(autoReconnect)
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

func (client *RTMServerClient) Connect() bool {
	return client.client.Connect()
}

func (client *RTMServerClient) Dial() bool {
	return client.client.Dial()
}

func (client *RTMServerClient) IsConnected() bool {
	return client.client.IsConnected()
}

/*
	Params:
		rest: can be include following params:
			pemPath		string
			rawPemData	[]byte
			reinforce	bool
*/
func (client *RTMServerClient) EnableEncryptor(rest ...interface{}) (err error) {
	return client.client.EnableEncryptor(rest...)
}

/*
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
}*/

func (client *RTMServerClient) genServerQuest(cmd string) *fpnn.Quest {

	now := time.Now()
	salt := now.UnixNano()
	ts := int32(now.Unix())

	pidStr := strconv.FormatInt(int64(client.pid), 10)
	saltStr := strconv.FormatInt(salt, 10)
	tsStr := strconv.FormatInt(int64(ts), 10)

	ctx := md5.New()
	io.WriteString(ctx, pidStr)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, client.secretKey)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, saltStr)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, cmd)
	io.WriteString(ctx, ":")
	io.WriteString(ctx, tsStr)

	sign := fmt.Sprintf("%X", ctx.Sum(nil))

	quest := fpnn.NewQuest(cmd)
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("ts", ts)

	return quest
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
			return client.client.SendQuest(quest, timeout)
		}

	} else {
		if timeout == 0 {
			return nil, client.client.SendQuestWithLambda(quest, callback)
		} else {
			return nil, client.client.SendQuestWithLambda(quest, callback, timeout)
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

func (client *RTMServerClient) sendStringQuest(quest *fpnn.Quest, timeout time.Duration,
	stringKey string, callback func(text string, errorCode int, errInfo string)) (string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantString(stringKey), fpnn.FPNN_EC_OK, "")
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
		return answer.WantString(stringKey), nil
	} else {
		return "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendTranscribeQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(text string, lang string, errorCode int, errInfo string)) (string, string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantString("text"), answer.WantString("lang"), fpnn.FPNN_EC_OK, "")
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
		return answer.WantString("text"), answer.WantString("lang"), nil
	} else {
		return "", "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func convertSliceToStringSlice(slice []interface{}) []string {
	if slice == nil || len(slice) == 0 {
		return make([]string, 0, 1)
	}

	rev := make([]string, 0, len(slice))
	for _, elem := range slice {
		val := convertToString(elem)
		rev = append(rev, val)
	}
	return rev
}

func (client *RTMServerClient) sendProfanityQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(text string, classification []string, errorCode int, errInfo string)) (string, []string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				text := answer.WantString("text")
				slice, _ := answer.GetSlice("classification")
				classification := convertSliceToStringSlice(slice)
				callback(text, classification, fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback("", make([]string, 0, 1), errorCode, "")
			} else {
				callback("", make([]string, 0, 1), answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return "", make([]string, 0, 1), err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return "", make([]string, 0, 1), err
	} else if !answer.IsException() {
		text := answer.WantString("text")
		slice, _ := answer.GetSlice("classification")
		classification := convertSliceToStringSlice(slice)
		return text, classification, nil
	} else {
		return "", make([]string, 0, 1), fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendGetObjectInfoQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(publicInfo string, privateInfo string, errorCode int, errInfo string)) (string, string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantString("oinfo"), answer.WantString("pinfo"), fpnn.FPNN_EC_OK, "")
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
		return answer.WantString("oinfo"), answer.WantString("pinfo"), nil
	} else {
		return "", "", fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendGetMsgInfoQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(id int64, mtype int8, msg string, attr string, mtime int64, errorCode int, errInfo string)) (int64, int8, string, string, int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantInt64("id"), answer.WantInt8("mtype"), answer.WantString("msg"), answer.WantString("attr"), answer.WantInt64("mtime"), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(0, 0, "", "", 0, errorCode, "")
			} else {
				callback(0, 0, "", "", 0, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return 0, 0, "", "", 0, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return 0, 0, "", "", 0, err
	} else if !answer.IsException() {
		return answer.WantInt64("id"), answer.WantInt8("mtype"), answer.WantString("msg"), answer.WantString("attr"), answer.WantInt64("mtime"), nil
	} else {
		return 0, 0, "", "", 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

//-----------[ Manage functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			clientEndpoint string
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) Kickout(uid int64, rest ...interface{}) error {

	var clientEndpoint string
	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case string:
			clientEndpoint = value
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.Kickout() function.")
		}
	}

	quest := client.genServerQuest("kickout")
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
func (client *RTMServerClient) GetToken(uid int64, rest ...interface{}) (string, error) {

	var timeout time.Duration
	var callback func(string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.GetToken() function.")
		}
	}

	quest := client.genServerQuest("gettoken")
	quest.Param("uid", uid)

	return client.sendTokenQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) RemoveToken(uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.RemoveToken() function.")
		}
	}

	quest := client.genServerQuest("removetoken")
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
			panic("Invaild params when call RTMServerClient.AddDevice() function.")
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
			panic("Invaild params when call RTMServerClient.RemoveDevice() function.")
		}
	}

	quest := client.genServerQuest("removedevice")
	quest.Param("uid", uid)
	quest.Param("devicetoken", deviceToken)

	return client.sendSilentQuest(quest, timeout, callback)
}
