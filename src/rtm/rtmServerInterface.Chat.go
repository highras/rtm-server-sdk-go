package rtm

import (
	"errors"
	"fmt"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	defaultMtype_Chat  = 30
	defaultMtype_Cmd   = 32
	defaultMtype_Image = 40
	defaultMtype_Audio = 41
	defaultMtype_Video = 42
	defaultMtype_File  = 50
)

var (
	getChatMtypes = []int8{defaultMtype_Chat, defaultMtype_Cmd, defaultMtype_Image, defaultMtype_Audio, defaultMtype_Video, defaultMtype_File}
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
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			return nil, errors.New("Invaild params when call RTMServerClient.GetGroupChat() function.")
		}
	}

	return client.GetGroupMessage(groupId, desc, num, begin, end, lastCursorId, uid, append(rest, getChatMtypes)...)
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
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			return nil, errors.New("Invaild params when call RTMServerClient.GetRoomChat() function.")
		}
	}

	return client.GetRoomMessage(roomId, desc, num, begin, end, lastCursorId, uid, append(rest, getChatMtypes)...)
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
	begin int64, end int64, lastCursorId int64, uid int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			return nil, errors.New("Invaild params when call RTMServerClient.GetBroadcastChat() function.")
		}
	}

	return client.GetBroadcastMessage(desc, num, begin, end, lastCursorId, uid, append(rest, getChatMtypes)...)
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
	begin int64, end int64, lastCursorId int64, rest ...interface{}) (*HistoryMessageResult, error) {

	for _, value := range rest {
		switch value.(type) {
		case []int8:
			return nil, errors.New("Invaild params when call RTMServerClient.GetP2PChat() function.")
		}
	}

	return client.GetP2PMessage(uid, peerUid, desc, num, begin, end, lastCursorId, append(rest, getChatMtypes)...)
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
func (client *RTMServerClient) DelP2PChat(messageId int64, fromUid int64, to int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, to, MessageType_P2P, rest...)
}

//-----------[ Get Chat functions ]-------------------//
/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (result *HistoryMessageUnit, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (result *HistoryMessageUnit, err error)
*/
func (client *RTMServerClient) GetChat(messageId int64, fromUid int64, xid int64, messageType MessageType, rest ...interface{}) (*HistoryMessageUnit, error) {
	return client.GetMessage(messageId, fromUid, xid, messageType, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelGroupChat(messageId int64, fromUid int64, groupId int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, groupId, MessageType_Group, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelRoomChat(messageId int64, fromUid int64, roomId int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, roomId, MessageType_Room, rest...)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) DelBroadcastChat(messageId int64, fromUid int64, rest ...interface{}) error {
	return client.DelMessage(messageId, fromUid, 0, MessageType_Broadcast, rest...)
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
	textType string, profanity string, postProfanity bool, uid int64, rest ...interface{}) (result *TranslateResult, err error) {

	var timeout time.Duration
	var callback func(*TranslateResult, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(*TranslateResult, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.Translate() function.")
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

	if postProfanity {
		quest.Param("postProfanity", postProfanity)
	}

	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendTranslateQuest(quest, timeout, callback)
}

// new translate api
/*
	Params:
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
func (client *RTMServerClient) TranslateByLanguageCode(text string, sourceLanguage RTMTranslateLanguage, targetLanguage RTMTranslateLanguage,
	textType string, profanity string, postProfanity bool, uid int64, rest ...interface{}) (result *TranslateResult, err error) {

	if len(targetLanguage.String()) <= 0 {
		return nil, fmt.Errorf("server not support this targetlanguage lang is: %s.", targetLanguage.String())
	}
	return client.Translate(text, sourceLanguage.String(), targetLanguage.String(), textType, profanity, postProfanity, uid, rest...)
}

//-----------[ Profanity functions ]-------------------//

/*
	Explain: maybe in after version this interface will be deprecated，recommend use TextCheck interface replace
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (text string, classification []string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", error);
		else this function work in sync mode, and return (text string, classification []string, err error)
*/
func (client *RTMServerClient) Profanity(text string, classify bool, uid int64, rest ...interface{}) (string, []string, error) {

	var timeout time.Duration
	var callback func(string, []string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, []string, int, string):
			callback = value
		default:
			return "", nil, errors.New("Invaild params when call RTMServerClient.Profanity() function.")
		}
	}

	quest := client.genServerQuest("profanity")
	quest.Param("text", text)
	quest.Param("classify", classify)

	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendProfanityQuest(quest, timeout, callback)
}

func (client *RTMServerClient) Speech2Text(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (string, string, error) {
	var timeout time.Duration = 120 * time.Second
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			return "", "", errors.New("Invalid params when call RTMServerClient.Speech2Text() function.")
		}
	}

	quest := client.genServerQuest("speech2text")
	quest.Param("audio", audio)
	quest.Param("type", audioType)
	quest.Param("lang", lang.String())
	if len(codec) != 0 {
		quest.Param("codec", codec)
	}
	if srate > 0 {
		quest.Param("srate", srate)
	}
	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendSpeech2Text(quest, timeout, callback)
}

func (client *RTMServerClient) TextCheck(text string, uid int64, rest ...interface{}) (int32, string, []int32, []string, error) {
	var timeout time.Duration
	var callback func(int32, string, []int32, []string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, string, []int32, []string, int, string):
			callback = value
		default:
			return -1, "", make([]int32, 0, 1), make([]string, 0, 1), errors.New("Invalid params when call RTMServerClient.TextCheck() function.")
		}
	}

	quest := client.genServerQuest("tcheck")
	quest.Param("text", text)
	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendTextCheck(quest, timeout, callback)
}

func (client *RTMServerClient) ImageCheck(image string, imageType int32, uid int64, rest ...interface{}) (int32, []int32, error) {
	var timeout time.Duration = 120 * time.Second
	var callback func(int32, []int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, []int32, int, string):
			callback = value
		default:
			return -1, make([]int32, 0, 1), errors.New("Invalid params when call RTMServerClient.ImageCheck() function.")
		}
	}

	quest := client.genServerQuest("icheck")
	quest.Param("image", image)
	quest.Param("type", imageType)
	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendOtherCheck(quest, timeout, callback)

}

func (client *RTMServerClient) AudioCheck(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (int32, []int32, error) {
	var timeout time.Duration = 120 * time.Second
	var callback func(int32, []int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, []int32, int, string):
			callback = value
		default:
			return -1, make([]int32, 0, 1), errors.New("Invalid params when call RTMServerClient.AudioCheck() function.")
		}
	}

	quest := client.genServerQuest("acheck")
	quest.Param("audio", audio)
	quest.Param("type", audioType)
	quest.Param("lang", lang.String())

	if len(codec) != 0 {
		quest.Param("codec", codec)
	}

	if srate > 0 {
		quest.Param("srate", srate)
	}

	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendOtherCheck(quest, timeout, callback)

}

func (client *RTMServerClient) VideoCheck(video string, videoType int32, videoName string, uid int64, rest ...interface{}) (int32, []int32, error) {
	var timeout time.Duration = 120 * time.Second
	var callback func(int32, []int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, []int32, int, string):
			callback = value
		default:
			return -1, make([]int32, 0, 1), errors.New("Invalid params when call RTMServerClient.VideoCheck() function.")
		}
	}

	quest := client.genServerQuest("vcheck")
	quest.Param("video", video)
	quest.Param("type", videoType)
	quest.Param("videoName", videoName)
	if uid > 0 {
		quest.Param("uid", uid)
	}

	return client.sendOtherCheck(quest, timeout, callback)
}
