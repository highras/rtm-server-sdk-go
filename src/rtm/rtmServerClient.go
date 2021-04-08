package rtm

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	SDKVersion = "0.9.1"
)

const (
	APIVersion = "2.7.0"
)

/*  for compatible before v0.3.1(include) maybe in after version this interface will be deprecated,
please use new serverPush interface IRTMServerMonitor
*/
type RTMServerMonitor interface {
	P2PMessage(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	GroupMessage(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	RoomMessage(fromUid int64, roomIid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
	P2PChat(fromUid int64, toUid int64, mid int64, message string, attrs string, mtime int64)
	GroupChat(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64)
	RoomChat(fromUid int64, roomIid int64, mid int64, message string, attrs string, mtime int64)
	P2PCmd(fromUid int64, toUid int64, mid int64, message string, attrs string, mtime int64)
	GroupCmd(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64)
	RoomCmd(fromUid int64, roomIid int64, mid int64, message string, attrs string, mtime int64)
	P2PFile(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	GroupFile(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
	RoomFile(fromUid int64, roomId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
}

// new serverpush version first use please use this interface
type IRTMServerMonitor interface {
	P2PMessage(messageInfo *RTMMessage)
	GroupMessage(messageInfo *RTMMessage)
	RoomMessage(messageInfo *RTMMessage)
	P2PChat(messageInfo *RTMMessage)
	GroupChat(messageInfo *RTMMessage)
	RoomChat(messageInfo *RTMMessage)

	P2PCmd(messageInfo *RTMMessage)
	GroupCmd(messageInfo *RTMMessage)
	RoomCmd(messageInfo *RTMMessage)

	P2PFile(messageInfo *RTMMessage)
	GroupFile(messageInfo *RTMMessage)
	RoomFile(messageInfo *RTMMessage)

	Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
}

//------------------------------[ RTM Server Client ]---------------------------------------//

type RtmRegressiveState struct {
	CurrentFailedCount       int
	ConnectStartMilliseconds int64
}

type RTMClientConnectEventUserCallback func(connId uint64, endpoint string, connected bool, autoReconnect bool, connectState *RtmRegressiveState)
type RTMClientCloseEventUserCallback func(connId uint64, endpoint string, autoReconnect bool, connectState *RtmRegressiveState)

type RTMServerClient struct {
	client                    *fpnn.TCPClient
	processor                 *rtmServerQuestProcessor
	logger                    *log.Logger
	pid                       int32
	secretKey                 string
	regressiveState           *RtmRegressiveState
	isClose                   bool
	defaultRegressiveStrategy *RTMRegressiveConnectStrategy
	regressiveConnectStrategy *RTMRegressiveConnectStrategy
	mutex                     sync.Mutex
	listenCache               *rtmListenCache
}

const (
	rtmServerConnectTimeout = 30 * time.Second
	rtmServerQuestTimeout   = 30 * time.Second
)

func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient {

	client := &RTMServerClient{}

	client.client = fpnn.NewTCPClient(endpoint)
	client.processor = newRTMServerQuestProcessor()
	client.secretKey = secretKey
	client.pid = pid
	client.regressiveState = &RtmRegressiveState{0, 0}
	client.isClose = false
	client.regressiveConnectStrategy = DefaultRegressiveStrategy
	client.listenCache = newRtmListenCache()

	client.client.SetQuestProcessor(client.processor)
	client.SetLogger(log.New(os.Stdout, "[RTM Go SDK] ", log.LstdFlags|log.Lshortfile))
	client.SetConnectTimeOut(rtmServerConnectTimeout)
	client.SetQuestTimeOut(rtmServerQuestTimeout)
	client.SetOnConnectedCallback(nil)
	client.SetOnClosedCallback(nil)
	return client
}

//------------------------------[ RTM Server Client Config Interfaces ]---------------------------------------//
/*	for compatible before v0.3.1(include) maybe in after version this interface will be deprecated,
	please use new set serverpush interface SetServerPushMonitor
*/
func (client *RTMServerClient) SetMonitor(monitor RTMServerMonitor) {
	client.processor.monitor = monitor
}

// new set serverpush api version first use please use this
func (client *RTMServerClient) SetServerPushMonitor(monitor IRTMServerMonitor) {
	client.processor.newMonitor = monitor
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

func (client *RTMServerClient) SetOnConnectedCallback(onConnect RTMClientConnectEventUserCallback) {
	client.client.SetOnConnectedCallback(func(connId uint64, endpoint string, connected bool) {
		if onConnect != nil {
			var state = &RtmRegressiveState{}
			state.ConnectStartMilliseconds = client.regressiveState.ConnectStartMilliseconds
			state.CurrentFailedCount = client.regressiveState.CurrentFailedCount
			onConnect(connId, endpoint, connected, client.canReconnect(), state)
		}
		//
		if connected {
			client.regressiveState.CurrentFailedCount = 0
			client.sendListenCache()
			return
		}

		if !connected && client.canReconnect() {
			go client.regressiveReconnection()
		}
	})
}

func (client *RTMServerClient) SetOnClosedCallback(onClosed RTMClientCloseEventUserCallback) {
	client.client.SetOnClosedCallback(func(connId uint64, endpoint string) {
		if onClosed != nil {
			var state = &RtmRegressiveState{}
			state.ConnectStartMilliseconds = client.regressiveState.ConnectStartMilliseconds
			state.CurrentFailedCount = client.regressiveState.CurrentFailedCount
			onClosed(connId, endpoint, client.canReconnect(), state)
		}
		//
		if client.canReconnect() {
			go client.regressiveReconnection()
		}
	})
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
	client.isClose = false
	client.regressiveState.ConnectStartMilliseconds = time.Now().UnixNano() / 1e6
	return client.client.Connect()
}

func (client *RTMServerClient) Dial() bool {
	return client.client.Dial()
}

func (client *RTMServerClient) IsConnected() bool {
	return client.client.IsConnected()
}

func (client *RTMServerClient) SetDefaultRegressiveConnectStrategy(strategy *RTMRegressiveConnectStrategy) {
	if strategy == nil {
		return
	}
	if client.defaultRegressiveStrategy == nil {
		client.defaultRegressiveStrategy = strategy
		client.regressiveConnectStrategy = strategy
	}
}

func (client *RTMServerClient) Close() {
	client.isClose = true
	client.client.Close()
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
	salt := idGen.genMid()
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
func (client *RTMServerClient) convertToInt64(value interface{}) int64 {
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
		client.logger.Printf("[ERROR] convertToInt64 Type %v convert failed.", reflect.TypeOf(value))
		return 0
	}
}

func (client *RTMServerClient) convertToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case []byte:
		return string(value.([]byte))
	case []rune:
		return string(value.([]rune))
	default:
		client.logger.Printf("[ERROR] convertToString Type  %v convert failed.", reflect.TypeOf(value))
		return ""
	}
}

func (client *RTMServerClient) canReconnect() bool {
	if !client.client.GetAutoReconnect() {
		return false
	}

	if client.isClose {
		return false
	}

	return true
}

func (client *RTMServerClient) reconnect() {
	if client.IsConnected() {
		return
	}
	if client.canReconnect() {
		client.Connect()
	}
}

func (client *RTMServerClient) regressiveReconnection() {
	current := int64(time.Now().UnixNano() / 1e6)
	strategy := client.regressiveConnectStrategy
	internval := current - int64(client.regressiveState.ConnectStartMilliseconds)
	if internval > int64(strategy.connectFailedMaxIntervalMilliseconds) {
		client.regressiveState.CurrentFailedCount = 0
		client.reconnect()
		return
	}
	client.regressiveState.CurrentFailedCount++
	if client.regressiveState.CurrentFailedCount <= strategy.startConnectFailedCount {
		client.reconnect()
		return
	}

	idleSeconds := strategy.maxIntervalSeconds - strategy.firstIntervalSeconds
	perIdleMillisecond := idleSeconds * 1000 / strategy.linearRegressiveCount
	currIdleMilliseconds := (client.regressiveState.CurrentFailedCount-strategy.startConnectFailedCount)*perIdleMillisecond + strategy.firstIntervalSeconds*1000
	if currIdleMilliseconds > strategy.maxIntervalSeconds*1000 {
		currIdleMilliseconds = strategy.maxIntervalSeconds * 1000
	}
	time.Sleep(time.Duration(currIdleMilliseconds) * time.Millisecond)

	client.reconnect()
}

func (client *RTMServerClient) addRTMListenCache(uids []int64, groupIds []int64, roomIds []int64, events []string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if uids != nil && len(uids) > 0 {
		client.listenCache.addUids(uids)
	}
	if groupIds != nil && len(groupIds) > 0 {
		client.listenCache.addGroupIds(groupIds)
	}
	if roomIds != nil && len(roomIds) > 0 {
		client.listenCache.addRoomIds(roomIds)
	}
	if events != nil && len(events) > 0 {
		client.listenCache.addEvents(events)
	}
}

func (client *RTMServerClient) removeRTMListenCache(uids []int64, groupIds []int64, roomIds []int64, events []string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if uids != nil && len(uids) > 0 {
		client.listenCache.removeUids(uids)
	}
	if groupIds != nil && len(groupIds) > 0 {
		client.listenCache.removeGroupIds(groupIds)
	}
	if roomIds != nil && len(roomIds) > 0 {
		client.listenCache.removeRoomIds(roomIds)
	}
	if events != nil && len(events) > 0 {
		client.listenCache.removeEvents(events)
	}
}

func (client *RTMServerClient) setRTMListenCache(uids []int64, groupIds []int64, roomIds []int64, events []string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if uids != nil && len(uids) > 0 {
		client.listenCache.setUids(uids)
	}
	if groupIds != nil && len(groupIds) > 0 {
		client.listenCache.setGroupIds(groupIds)
	}
	if roomIds != nil && len(roomIds) > 0 {
		client.listenCache.setRoomIds(roomIds)
	}
	if events != nil && len(events) > 0 {
		client.listenCache.setEvents(events)
	}
}

func (client *RTMServerClient) setRTMListenStateCache(allP2p bool, allGroup bool, allRoom bool, allEvent bool) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.listenCache.setAllUid(allP2p)
	client.listenCache.setAllGroup(allGroup)
	client.listenCache.setAllRoom(allRoom)
	client.listenCache.setAllEvent(allEvent)
}

func (client *RTMServerClient) sendListenCache() {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if !client.listenCache.empty() {
		quest := client.genServerQuest("setlisten")

		if len(client.listenCache.listenUids) > 0 {
			quest.Param("uids", client.listenCache.listenUids)
		}
		if len(client.listenCache.listenRoomIds) > 0 {
			quest.Param("rids", client.listenCache.listenRoomIds)
		}
		if len(client.listenCache.listenGroupIds) > 0 {
			quest.Param("gids", client.listenCache.listenGroupIds)
		}
		if len(client.listenCache.listenEvents) > 0 {
			quest.Param("events", client.listenCache.listenEvents)
		}
		err := client.sendSilentQuest(quest, 0, func(errorCode int, errInfo string) {
			if errorCode != fpnn.FPNN_EC_OK {
				client.logger.Printf("[ERROR] connected send add listencache error, errorCode:= %d, errorInfo:= %s.\n", errorCode, errInfo)
			}
		})
		if err != nil {
			client.logger.Printf("[ERROR] connected send add listencache error in async mode, err: %v.\n", err)
		}
	}

	if !client.listenCache.isAllFalse() {
		quest := client.genServerQuest("setlisten")
		quest.Param("group", client.listenCache.allGroup)
		quest.Param("room", client.listenCache.allRoom)
		quest.Param("p2p", client.listenCache.allP2P)
		quest.Param("ev", client.listenCache.allEvent)
		err := client.sendSilentQuest(quest, 0, func(errorCode int, errInfo string) {
			if errorCode != fpnn.FPNN_EC_OK {
				client.logger.Printf("[ERROR] connected send set listencache error, errorCode:= %d, errorInfo:= %s.\n", errorCode, errInfo)
			}
		})
		if err != nil {
			client.logger.Printf("[ERROR] connected send set listencache error in async mode, err: %v.\n", err)
		}
	}
}

//------------------------------[ RTM Server Client Interfaces ]---------------------------------------//
func (client *RTMServerClient) sendQuest(quest *fpnn.Quest, timeout time.Duration, callback func(answer *fpnn.Answer, errorCode int)) (*fpnn.Answer, error) {

	if callback == nil {
		var answer *fpnn.Answer
		var err error
		if timeout == 0 {
			answer, err = client.client.SendQuest(quest)
		} else {
			answer, err = client.client.SendQuest(quest, timeout)
		}
		if err != nil {
			return answer, err
		} else if answer.IsException() {
			code, _ := answer.GetInt("code")
			if code == fpnn.FPNN_EC_CORE_CONNECTION_CLOSED || code == fpnn.FPNN_EC_CORE_INVALID_CONNECTION {
				if timeout == 0 {
					return client.client.SendQuest(quest)
				} else {
					return client.client.SendQuest(quest, timeout)
				}
			}
		}
		return answer, err

	} else {
		if timeout == 0 {
			return nil, client.client.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_CORE_CONNECTION_CLOSED || errorCode == fpnn.FPNN_EC_CORE_INVALID_CONNECTION {
					err := client.client.SendQuestWithLambda(quest, callback)
					if err != nil {
						client.logger.Println("[ERROR] send async quest failed, err: ", err)
					}
					return
				}
				callback(answer, errorCode)
			})
		} else {
			return nil, client.client.SendQuestWithLambda(quest, func(answer *fpnn.Answer, errorCode int) {
				if errorCode == fpnn.FPNN_EC_CORE_CONNECTION_CLOSED || errorCode == fpnn.FPNN_EC_CORE_INVALID_CONNECTION {
					err := client.client.SendQuestWithLambda(quest, callback, timeout)
					if err != nil {
						client.logger.Println("[ERROR] send async quest failed, err: ", err)
					}
					return
				}
				callback(answer, errorCode)
			}, timeout)
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

func (client *RTMServerClient) convertSliceToInt64Slice(slice []interface{}) []int64 {
	if slice == nil || len(slice) == 0 {
		return make([]int64, 0, 1)
	}

	rev := make([]int64, 0, len(slice))
	for _, elem := range slice {
		val := client.convertToInt64(elem)
		rev = append(rev, val)
	}
	return rev
}

func (client *RTMServerClient) convertSliceToInt32Slice(slice []interface{}) []int32 {
	if slice == nil || len(slice) == 0 {
		return make([]int32, 0, 1)
	}

	rev := make([]int32, 0, len(slice))
	for _, elem := range slice {
		val := int32(client.convertToInt64(elem))
		rev = append(rev, val)
	}
	return rev
}

func (client *RTMServerClient) sendSliceQuest(quest *fpnn.Quest, timeout time.Duration,
	sliceKey string, callback func(slice []int64, errorCode int, errInfo string)) ([]int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertSliceToInt64Slice(answer.WantSlice(sliceKey)), fpnn.FPNN_EC_OK, "")
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
		slice := client.convertSliceToInt64Slice(answer.WantSlice(sliceKey))
		return slice, nil
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendDoubleSliceQuest(quest *fpnn.Quest, timeout time.Duration,
	sliceKey string, secondSliceKey string, callback func(first []int64, second []int64, errorCode int, errInfo string)) ([]int64, []int64, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(client.convertSliceToInt64Slice(answer.WantSlice(sliceKey)), client.convertSliceToInt64Slice(answer.WantSlice(secondSliceKey)), fpnn.FPNN_EC_OK, "")
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
		slice1 := client.convertSliceToInt64Slice(answer.WantSlice(sliceKey))
		slice2 := client.convertSliceToInt64Slice(answer.WantSlice(secondSliceKey))
		return slice1, slice2, nil
	} else {
		return nil, nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
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

func (client *RTMServerClient) sendIntQuest(quest *fpnn.Quest, timeout time.Duration, stringKey string,
	callback func(count int32, errorCode int, errInfo string)) (int32, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantInt32(stringKey), fpnn.FPNN_EC_OK, "")
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
		return answer.WantInt32(stringKey), nil
	} else {
		return 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendDoubleIntQuest(quest *fpnn.Quest, timeout time.Duration, stringKey string, stringKey1 string,
	callback func(sender int32, count int32, errorCode int, errInfo string)) (int32, int32, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				callback(answer.WantInt32(stringKey), answer.WantInt32(stringKey1), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(0, 0, errorCode, "")
			} else {
				callback(0, 0, answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return 0, 0, err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return 0, 0, err
	} else if !answer.IsException() {
		return answer.WantInt32(stringKey), answer.WantInt32(stringKey1), nil
	} else {
		return 0, 0, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
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

func (client *RTMServerClient) convertSliceToStringSlice(slice []interface{}) []string {
	if slice == nil || len(slice) == 0 {
		return make([]string, 0, 1)
	}

	rev := make([]string, 0, len(slice))
	for _, elem := range slice {
		val := client.convertToString(elem)
		rev = append(rev, val)
	}
	return rev
}

func (client *RTMServerClient) convertStringSliceToInt32Slice(slice []string) []int32 {
	if slice == nil || len(slice) == 0 {
		return make([]int32, 0, 1)
	}

	rev := make([]int32, 0, len(slice))
	for _, elem := range slice {
		if v, err := strconv.Atoi(elem); err == nil {
			rev = append(rev, int32(v))
		} else {
			client.logger.Printf("[ERROR] str atoi failed, value: %v, type: %v.", elem, reflect.TypeOf(elem))
		}

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
				classification := client.convertSliceToStringSlice(slice)
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
		classification := client.convertSliceToStringSlice(slice)
		return text, classification, nil
	} else {
		return "", make([]string, 0, 1), fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendSpeech2Text(quest *fpnn.Quest, timeout time.Duration,
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

func (client *RTMServerClient) sendOtherCheck(quest *fpnn.Quest, timeout time.Duration,
	callback func(result int32, tags []int32, errorCode int, errInfo string)) (int32, []int32, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				sliceTag, _ := answer.GetSlice("tags")
				tags := client.convertSliceToStringSlice(sliceTag)
				callback((int32)(answer.WantInt("result")), client.convertStringSliceToInt32Slice(tags), fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(-1, make([]int32, 0, 1), errorCode, "")
			} else {
				callback(-1, make([]int32, 0, 1), answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return -1, make([]int32, 0, 1), err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return -1, make([]int32, 0, 1), err
	} else if !answer.IsException() {
		sliceTag, _ := answer.GetSlice("tags")
		tags := client.convertSliceToStringSlice(sliceTag)
		return (int32)(answer.WantInt("result")), client.convertStringSliceToInt32Slice(tags), nil
	} else {
		return -1, make([]int32, 0, 1), fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

func (client *RTMServerClient) sendTextCheck(quest *fpnn.Quest, timeout time.Duration,
	callback func(result int32, text string, tags []int32, wlist []string, errorCode int, errInfo string)) (int32, string, []int32, []string, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				text, _ := answer.GetString("text")
				sliceTag, _ := answer.GetSlice("tags")
				tags := client.convertSliceToStringSlice(sliceTag)
				sliceSensitive, _ := answer.GetSlice("wlist")
				wlist := client.convertSliceToStringSlice(sliceSensitive)
				callback((int32)(answer.WantInt("result")), text, client.convertStringSliceToInt32Slice(tags), wlist, fpnn.FPNN_EC_OK, "")
			} else if answer == nil {
				callback(-1, "", make([]int32, 0, 1), make([]string, 0, 1), errorCode, "")
			} else {
				callback(-1, "", make([]int32, 0, 1), make([]string, 0, 1), answer.WantInt("code"), answer.WantString("ex"))
			}
		}

		_, err := client.sendQuest(quest, timeout, callbackFunc)
		return -1, "", make([]int32, 0, 1), make([]string, 0, 1), err
	}

	answer, err := client.sendQuest(quest, timeout, nil)
	if err != nil {
		return -1, "", make([]int32, 0, 1), make([]string, 0, 1), err
	} else if !answer.IsException() {
		text, _ := answer.GetString("text")
		sliceTag, _ := answer.GetSlice("tags")
		tags := client.convertSliceToStringSlice(sliceTag)
		sliceSensitive, _ := answer.GetSlice("wlist")
		wlist := client.convertSliceToStringSlice(sliceSensitive)
		return (int32)(answer.WantInt("result")), text, client.convertStringSliceToInt32Slice(tags), wlist, nil
	} else {
		return -1, "", make([]int32, 0, 1), make([]string, 0, 1), fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
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

func (client *RTMServerClient) getMsgInfo(answer *fpnn.Answer) (res *HistoryMessageUnit, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ERROR] Process MsgInfo exception. Panic: %v.", r)
		}
	}()

	result := &HistoryMessageUnit{}
	result.CursorId = answer.WantInt64("id")
	result.MessageType = answer.WantInt8("mtype")
	msg := answer.WantString("msg")
	result.Attrs = answer.WantString("attrs")
	result.ModifiedTime = answer.WantInt64("mtime")
	result.Message = msg
	if result.MessageType >= defaultMtype_Image && result.MessageType <= defaultMtype_File {
		fileInfo := processFileInfo(msg, result.Attrs, result.MessageType, client.logger)
		result.FileInfo = fileInfo
		result.Attrs = fetchFileCustomAttrs(result.Attrs, client.logger)
	}

	return result, err
}

func (client *RTMServerClient) sendGetMsgInfoQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(result *HistoryMessageUnit, errorCode int, errInfo string)) (*HistoryMessageUnit, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				if result, err := client.getMsgInfo(answer); err == nil {
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
		if result, err := client.getMsgInfo(answer); err == nil {
			return result, nil
		} else {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
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

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.Kickout() function.")
		}
	}

	quest := client.genServerQuest("kickout")
	quest.Param("uid", uid)

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
			return "", errors.New("Invaild params when call RTMServerClient.GetToken() function.")
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
			return errors.New("Invaild params when call RTMServerClient.RemoveToken() function.")
		}
	}

	quest := client.genServerQuest("removetoken")
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}
