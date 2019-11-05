package rtm

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	defaultMtype_Chat  = 30
	defaultMtype_Audio = 31
	defaultMtype_Cmd   = 32
)

//-----------[ Chat functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			attrs string
			timeout time.Duration
			func (mtime int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (0, error);
		else this function work in sync mode, and return (mtime int64, err error)
*/
func (client *RTMServerClient) SendChat(fromUid int64, toUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessage(fromUid, toUid, defaultMtype_Chat, message, rest...)
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
func (client *RTMServerClient) SendAudio(fromUid int64, toUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessage(fromUid, toUid, defaultMtype_Audio, base64.StdEncoding.EncodeToString([]byte(message)), rest...)
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
func (client *RTMServerClient) SendCmd(fromUid int64, toUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessage(fromUid, toUid, defaultMtype_Cmd, message, rest...)
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
func (client *RTMServerClient) SendChats(fromUid int64, toUids []int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessages(fromUid, toUids, defaultMtype_Chat, message, rest...)
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
func (client *RTMServerClient) SendAudios(fromUid int64, toUids []int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessages(fromUid, toUids, defaultMtype_Audio, base64.StdEncoding.EncodeToString([]byte(message)), rest...)
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
func (client *RTMServerClient) SendCmds(fromUid int64, toUids []int64, message string, rest ...interface{}) (int64, error) {
	return client.SendMessages(fromUid, toUids, defaultMtype_Cmd, message, rest...)
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
func (client *RTMServerClient) SendGroupChat(fromUid int64, groupId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendGroupMessage(fromUid, groupId, defaultMtype_Chat, message, rest...)
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
func (client *RTMServerClient) SendGroupAudio(fromUid int64, groupId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendGroupMessage(fromUid, groupId, defaultMtype_Audio, base64.StdEncoding.EncodeToString([]byte(message)), rest...)
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
func (client *RTMServerClient) SendGroupCmd(fromUid int64, groupId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendGroupMessage(fromUid, groupId, defaultMtype_Cmd, message, rest...)
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
func (client *RTMServerClient) SendRoomChat(fromUid int64, roomId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendRoomMessage(fromUid, roomId, defaultMtype_Chat, message, rest...)
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
func (client *RTMServerClient) SendRoomAudio(fromUid int64, roomId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendRoomMessage(fromUid, roomId, defaultMtype_Audio, base64.StdEncoding.EncodeToString([]byte(message)), rest...)
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
func (client *RTMServerClient) SendRoomCmd(fromUid int64, roomId int64, message string, rest ...interface{}) (int64, error) {
	return client.SendRoomMessage(fromUid, roomId, defaultMtype_Cmd, message, rest...)
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
func (client *RTMServerClient) SendBroadcastChat(fromUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendBroadcastMessage(fromUid, defaultMtype_Chat, message, rest...)
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
func (client *RTMServerClient) SendBroadcastAudio(fromUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendBroadcastMessage(fromUid, defaultMtype_Audio, base64.StdEncoding.EncodeToString([]byte(message)), rest...)
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
func (client *RTMServerClient) SendBroadcastCmd(fromUid int64, message string, rest ...interface{}) (int64, error) {
	return client.SendBroadcastMessage(fromUid, defaultMtype_Cmd, message, rest...)
}

//-----------[ History Chats functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			panic("Invaild params when call RTMServerClient.GetGroupChat() function.")
		}
	}

	return client.GetGroupMessage(groupId, desc, num, begin, end, lastid, append(rest, []int8{defaultMtype_Chat, defaultMtype_Audio, defaultMtype_Cmd})...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			panic("Invaild params when call RTMServerClient.GetRoomChat() function.")
		}
	}

	return client.GetRoomMessage(roomId, desc, num, begin, end, lastid, append(rest, []int8{defaultMtype_Chat, defaultMtype_Audio, defaultMtype_Cmd})...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16,
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			panic("Invaild params when call RTMServerClient.GetBroadcastChat() function.")
		}
	}

	return client.GetBroadcastMessage(desc, num, begin, end, lastid, append(rest, []int8{defaultMtype_Chat, defaultMtype_Audio, defaultMtype_Cmd})...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *HistoryMessageResult, err error)
*/
func (client *RTMServerClient) GetP2PChat(uid int64, peerUid int64, desc bool, num int16,
	begin int64, end int64, lastid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			panic("Invaild params when call RTMServerClient.GetP2PChat() function.")
		}
	}

	return client.GetP2PMessage(uid, peerUid, desc, num, begin, end, lastid, append(rest, []int8{defaultMtype_Chat, defaultMtype_Audio, defaultMtype_Cmd})...)
}

//-----------[ Delete Chat functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelP2PChat(mid int64, fromUid int64, to int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, to, MessageType_P2P, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelGroupChat(mid int64, fromUid int64, gid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, gid, MessageType_Group, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelRoomChat(mid int64, fromUid int64, rid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, rid, MessageType_Room, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelBroadcastChat(mid int64, fromUid int64, rest ...interface{}) error {
	return client.DelMessage(mid, fromUid, 0, MessageType_Broadcast, rest...)
}

//-----------[ Translate functions ]-------------------//

type TranslateResult struct {
	SourceLanguage string
	TargetLanguage string
	SourceText     string
	TargetText     string
}

func (client *RTMServerClient) processTranslateAnswer(answer *fpnn.Answer) (res *TranslateResult, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ERROR] Process translate answer exception. Panic: %v.", r)
		}
	}()

	result := &TranslateResult{}
	result.SourceLanguage = answer.WantString("source")
	result.TargetLanguage = answer.WantString("target")
	result.SourceText = answer.WantString("sourceText")
	result.TargetText = answer.WantString("targetText")

	return result, err
}

func (client *RTMServerClient) sendTranslateQuest(quest *fpnn.Quest, timeout time.Duration,
	callback func(result *TranslateResult, errorCode int, errInfo string)) (*TranslateResult, error) {

	if callback != nil {
		callbackFunc := func(answer *fpnn.Answer, errorCode int) {
			if errorCode == fpnn.FPNN_EC_OK {
				if result, err := client.processTranslateAnswer(answer); err == nil {
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
		return client.processTranslateAnswer(answer)
	} else {
		return nil, fmt.Errorf("[Exception] code: %d, ex: %s", answer.WantInt("code"), answer.WantString("ex"))
	}
}

/*
	Params:
		sourceLanguage:
			A ISO 639-1 language code. Can be empty.

		targetLanguage:
			A ISO 639-1 language code. Cannot be empty.

		textType:
			"chat": will modify the '\t', '\n', ' ' for output;
			"mail": keep the '\t', '\n', ' ' as original.

			Can be empty which means "chat".

		profanity:
			"off": without sensitive words filtering;
			"stop": will return error when sensitive words found;
			"censor": replace all sensitive words as '*'.

			Can be empty as "off".

		rest: can be include following params:
			timeout time.Duration
			func (result *TranslateResult, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (result *TranslateResult, err error)
*/
func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string,
	textType string, profanity string, rest ...interface{}) (result *TranslateResult, err error) {

	var timeout time.Duration
	var callback func(*TranslateResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(*TranslateResult, int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.Translate() function.")
		}
	}

	quest := client.genServerQuest("translate")
	quest.Param("text", text)

	quest.Param("dst", targetLanguage)

	if len(sourceLanguage) != 0 {
		quest.Param("src", sourceLanguage)
	}

	if len(textType) != 0 {
		quest.Param("type", textType)
	}

	if len(profanity) != 0 {
		quest.Param("profanity", profanity)
	}

	return client.sendTranslateQuest(quest, timeout, callback)
}

//-----------[ Profanity functions ]-------------------//

/*
	Params:
		action:
			"stop": will return error when sensitive words found;
			"censor": replace all sensitive words as '*'.

		Can be empty as "censor".

		rest: can be include following params:
			timeout time.Duration
			func (text string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", error);
		else this function work in sync mode, and return (text string, err error)
*/
func (client *RTMServerClient) Profanity(text string, action string, rest ...interface{}) (string, error) {

	var timeout time.Duration
	var callback func(string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.Profanity() function.")
		}
	}

	quest := client.genServerQuest("profanity")
	quest.Param("text", text)
	quest.Param("action", action)

	return client.sendStringQuest(quest, timeout, "text", callback)
}

/*
	Params:
		action:
			"stop": will return error when sensitive words found;
			"censor": replace all sensitive words as '*'.

		Can be empty as "censor".

		rest: can be include following params:
			timeout time.Duration
			func (text string, lang string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", error);
		else this function work in sync mode, and return (text string, err error)
*/
func (client *RTMServerClient) Transcribe(audio string, action string, lang string, rest ...interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			panic("Invaild params when call RTMServerClient.Transcribe() function.")
		}
	}

	quest := client.genServerQuest("transcribe")
	quest.Param("audio", audio)
	quest.Param("action", action)
	quest.Param("lang", lang)

	return client.sendTranscribeQuest(quest, timeout, callback)
}
