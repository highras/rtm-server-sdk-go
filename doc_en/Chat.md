# RTM Server-End Go SDK Chat API Docs

# Index

[TOC]

### -----------------------[ Send chat message interface ]-----------------------------

### func (client *RTMServerClient) SendChat(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendChat(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

Send P2P chat messages

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendCmd(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendCmd(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

Send P2P control commands.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendChats(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendChats(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

Send multi-person P2P chat messages.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendCmds(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendCmds(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

Send multi-person P2P control commands.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendGroupChat(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupChat(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

Send group chat messages.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendGroupCmd(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupCmd(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

Send group control commands.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendRoomChat(fromUid int64, roomId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomChat(fromUid int64, roomId int64, message string, rest ... interface{}) (int64, error)

Send room chat messages.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendBroadcastChat(fromUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastChat(fromUid int64, message string, rest ... interface{}) (int64, error)

Send broadcast chat messages.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendBroadcastCmd(fromUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastCmd(fromUid int64, message string, rest ... interface{}) (int64, error)

Send broadcast control commands.

The acceptable parameters are：

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.


### -----------------------[ Get chat history ]-----------------------------

### type RTMAudioFileInfo

	type RTMAudioFileInfo struct {
		IsRTMaudio bool   // Is it an rtm voice message
		Codec      string // rtm voice message has this value encoding format
		Srate      int32  // This value is available for rtm voice messages. Sampling rate
		Lang       string // If it is rtm voice, there will be this value
		Duration   int32  // ms，If it is rtm voice, there will be this value
	}

### type FileMsgInfo

	type FileMsgInfo struct {
		Url      string `json:"url"`
		FileSize int64  `json:"size"` // Byte size
		Surl     string `json:"surl"` // The address of the thumbnail, if it is a picture type, it will have this value
		RTMAudioFileInfo
	}

### type RTMMessage

	type RTMMessage struct {
		FromUid			int64
		ToId        	int64
		MessageType 	int8
		MessageId   	int64
		Message     	string
		Attrs       	string
		ModifiedTime	int64
		FileInfo     	*FileMsgInfo
	}

### type HistoryMessageUnit

	type HistoryMessageUnit struct {
		CursorId int64
		RTMMessage
	}

Chat history data unit.

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num				int16						//-- The number of entries actually returned
		LastCursorId	int64						//-- When the polling continues, the value of the lastid parameter used in the next call
		Begin			int64						//-- When continuing to poll, the value of the begin parameter used in the next call
		End				int64						//-- When continuing to poll, the value of the end parameter used in the next call
		Messages		[]*HistoryMessageUnit
	}

The chat history returns results.

### func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get group chat history.

Required parameters:

+ `desc bool`

	false: Starting from the timestamp of begin, turn pages sequentially.  
	true： Starting from the timestamp of end, turn pages in reverse order.

+ `num int16`

	Get the number of entries. 10 pieces are recommended, 20 pieces at a time at most.

+ `begin int64`

	Start timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: >=

+ `end int64`

	End timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: <=

+ `lastid int64`

	The id of the last message, default 0 for the first time. Condition:> or <

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.


### func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get room chat history.

Required parameters:

+ `desc bool`

	false: Starting from the timestamp of begin, turn pages sequentially.  
	true： Starting from the timestamp of end, turn pages in reverse order.

+ `num int16`

	Get the number of entries. 10 pieces are recommended, 20 pieces at a time at most.

+ `begin int64`

	Start timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: >=

+ `end int64`

	End timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: <=

+ `lastid int64`

	The id of the last message, default 0 for the first time. Condition:> or <

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.



### func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get broadcast chat history.

Required parameters:

+ `desc bool`

	false: Starting from the timestamp of begin, turn pages sequentially.  
	true： Starting from the timestamp of end, turn pages in reverse order

+ `num int16`

	Get the number of entries. 10 pieces are recommended, 20 pieces at a time at most.

+ `begin int64`

	Start timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: >=

+ `end int64`

	End timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: <=

+ `lastid int64`

	The id of the last message, default 0 for the first time. Condition:> or <

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.



### func (client *RTMServerClient) GetP2PChat(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetP2PChat(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get P2P chat history.

Required parameters:

+ `desc bool`

	false: Starting from the timestamp of begin, turn pages sequentially.  
	true： Starting from the timestamp of end, turn pages in reverse order

+ `num int16`

	Get the number of entries. 10 pieces are recommended, 20 pieces at a time at most.

+ `begin int64`

	Start timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: >=

+ `end int64`

	End timestamp, accurate to **milliseconds**, default 0. Use the current server time. Condition: <=

+ `lastid int64`

	The id of the last message, default 0 for the first time. Condition:> or <

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.


### -----------------------[ delete chat history ]-----------------------------

### type HistoryMessageUnit

	type MessageType int

	const (
		MessageType_P2P MessageType = iota
		MessageType_Group
		MessageType_Room
		MessageType_Broadcast
	)

Chat type definition.

### func (client *RTMServerClient) DelP2PChat(mid int64, fromUid int64, to int64, rest ... interface{}) error

	func (client *RTMServerClient) DelP2PChat(mid int64, fromUid int64, to int64, rest ... interface{}) error

Delete/retract chat messages.

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelGroupChat(mid int64, fromUid int64, gid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupChat(mid int64, fromUid int64, gid int64, rest ... interface{}) error

Delete/withdraw group messages.

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelRoomChat(mid int64, fromUid int64, rid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomChat(mid int64, fromUid int64, rid int64, rest ... interface{}) error

Delete/withdraw broadcast messages.

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelBroadcastChat(mid int64, fromUid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelBroadcastChat(mid int64, fromUid int64, rest ... interface{}) error

Delete/withdraw broadcast messages.

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### -----------------------[ Translate chat message ]-----------------------------

### type TranslateResult

	type TranslateResult struct {
		SourceLanguage		string
		TargetLanguage		string
		SourceText			string
		TargetText			string
	}

The result for Translate chat message.

### func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string,
	textType string, profanity string, postProfanity bool, uid int64, rest ...interface{}) (result *TranslateResult, err error)

	func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string, textType string, profanity string, rest ... interface{}) (result *TranslateResult, err error)

Translate chat message.

Required parameters:

+ `sourceLanguage string`

	Source language ISO 639-1 code.

	If it is an empty string, the system will automatically detect the source language.

+ `targetLanguage string`

	The target language ISO 639-1 code.

+ `textType string`

	Source data type.

	Optional values:

		* chat：'\t'、'\n'、' ' May be modified in the output text；
		* mail：In the output text，'\t'、'\n'、' ' will remain unchanged.

	If it is an empty string, the default is `chat`。

+ `profanity string`

	Whether to trigger sensitive word filtering.

	Optional value：

		* off：No sensitive word detection；
		* stop：When a sensitive word is found, the interface returns an error；
		* censor：When a sensitive word is found, the sensitive word will be replaced by `*`

	If it is an empty string, the default is `off`。

+ `postProfanity bool`

	Whether to filter the translated text.


The acceptable parameters are：


+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (result *TranslateResult, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *TranslateResult and error information is returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *TranslateResult result will be passed through callback.


### -----------------------[ Sensitive word filtering ]-----------------------------

### func (client *RTMServerClient) Profanity(text string, classify bool, uid int64, rest ... interface{}) (string, error)

	func (client *RTMServerClient) Profanity(text string, classify bool, uid int64, rest ... interface{}) (string, error)

Sensitive word filtering.

Note: 

* maybe in after version this interface will be deprecated，recommend use TextCheck interface replace

Required parameters:

+ `classify bool`

	Whether to perform text classification detection.

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (text string, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the filtered text and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" and error information. The real filtered text will be passed through callback.


### -----------------------[ Speech Recognition ]-------------------------------

### func (client *RTMServerClient) Speech2Text(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (string, string, error)

	func (client *RTMServerClient) Speech2Text(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (string, string, error)

Speech Recognition.

Required parameters:

+ `audio []byte`

	Audio URL or content (lang&codec&srate is required)

+ `audioType int32`

	Voice data type. 1: url 2: content

+ `lang RTMTranslateLanguage`

	language

Optional parameters
+ `codec string`

	If codec is empty, the default is AMR_WB

+ `srate int32`

	If srate is 0 or empty, the default is 16000

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (text string, lang string, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the text and lang and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" "" and error information. The real text will be passed through callback.


### -----------------------[ Text review ]-----------------------------

### func (client *RTMServerClient) TextCheck(text string, uid int64, rest ...interface{}) (int32, string, []int32, []string, error)

	func (client *RTMServerClient) TextCheck(text string, uid int64, rest ...interface{}) (result int32, text string, tags []int32, wlist []string, error)

Text review.

Required parameters:

+ `text string`

	Text content

Optional parameters:

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (result int32, text string, tags []int32, wlist []string, errorCode int, errorMessage string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the filtered text and lang and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" "" and error information. The real filtered text will be passed through callback. result: 0 pass 2 fail; text: sensitive word filtered text content, the sensitive words contained in it will be replaced with *, if it is not starred, there will be no field; tags: trigger classification, such as pornographic and political Wait, see text review classification for details; wlist: list of sensitive words

Note: If the returned result=2, the normal processing is: if the text is not empty, it can be sent directly (using the returned text), otherwise it is blocked (may be an advertisement or obscure pornography, etc.)

Note: If you need detailed return results, please call the native interface of the audit product


### -----------------------[ Picture review ]-----------------------------

### func (client *RTMServerClient) ImageCheck(image string, imageType int32, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) ImageCheck(image string, imageType int32, uid int64, rest ...interface{}) (int32, []int32, error)

Picture review.

Required parameters:

+ `image string`

	The url or content of the picture

+ `imageType int32`

	1, url, 2, content

Optional parameters:

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the filtered text and lang and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" "" and error information. The real filtered text will be passed through callback. Result: 0 passed, 2 failed; tags: triggered categories, such as pornographic and political, etc., see image review category for details

Note: If you need detailed return results, please call the native interface of the audit product


### -----------------------[ Video review ]-----------------------------

### func (client *RTMServerClient) VideoCheck(video string, videoType int32, videoName string, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) VideoCheck(video string, videoType int32, videoName string, uid int64, rest ...interface{}) (int32, []int32, error)

Video review

Required parameters:

+ `video string`

	The url or content of the video

+ `videoType int32`

	1, url, 2, content

+ `videoName string`

	Video file name, required when type=2, file format can be obtained by file name

Optional parameters:

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the filtered text and lang and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" "" and error information. The real result will be passed through callback. Result: 0 passed, 2 failed; tags: triggered categories, such as pornography and politics, etc., see the video review category for details.

Note: If you need detailed return results, please call the native interface of the audit product


### -----------------------[ Audio review ]-----------------------------

### func (client *RTMServerClient) AudioCheck(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) AudioCheck(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (int32, []int32, error)

Audio review

Required parameters:

+ `audio []byte`

	Audio URL or content (lang&codec&srate is required)

+ `audioType int32`

	Voice data type. 1: url 2: content

+ `lang RTMTranslateLanguage`

	language

Optional parameters:
+ `codec string`

	If codec is empty, the default is AMR_WB

+ `srate int32`

	If srate is 0 or empty, the default is 16000

+ `uid int64`

	user id

The acceptable parameters are：

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the filtered text and lang and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" "" and error information. The real result will be passed through callback. Result: 0 passed, 2 failed; tags: triggered categories, such as pornography and politics, etc., see audio review categories for details

Note: If you need detailed return results, please call the native interface of the audit product
