# RTM Server-End Go SDK Chat API Docs

# Index

[TOC]

### -----------------------[ 发送聊天信息接口 ]-----------------------------

### func (client *RTMServerClient) SendChat(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendChat(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

发送 P2P 聊天消息。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendCmd(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendCmd(fromUid int64, toUid int64, message string, rest ... interface{}) (int64, error)

发送 P2P 控制命令。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendChats(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendChats(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

发送多人 P2P 聊天消息。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendCmds(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendCmds(fromUid int64, toUids []int64, message string, rest ... interface{}) (int64, error)

发送多人 P2P 控制命令。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendGroupChat(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupChat(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

发送群组聊天消息。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendGroupCmd(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupCmd(fromUid int64, groupId int64, message string, rest ... interface{}) (int64, error)

发送群组控制命令。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendRoomChat(fromUid int64, roomId int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomChat(fromUid int64, roomId int64, message string, rest ... interface{}) (int64, error)

发送房间聊天消息。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendBroadcastChat(fromUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastChat(fromUid int64, message string, rest ... interface{}) (int64, error)

发送广播聊天消息。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendBroadcastCmd(fromUid int64, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastCmd(fromUid int64, message string, rest ... interface{}) (int64, error)

发送广播控制命令。

可接受的参数为：

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。


### -----------------------[ 获取聊天历史记录 ]-----------------------------

### type RTMAudioFileInfo

	type RTMAudioFileInfo struct {
		IsRTMaudio bool   // 是否是rtm语音消息
		Codec      string // rtm语音消息时有此值编 码格式
		Srate      int32  // rtm语音消息时有此值 采样率
		Lang       string // 如果是rtm语音会有此值
		Duration   int32  // ms，如果是rtm语音会有此值
	}

### type FileMsgInfo

	type FileMsgInfo struct {
		Url      string `json:"url"`
		FileSize int64  `json:"size"` // 字节大小
		Surl     string `json:"surl"` // 缩略图的地址，如果是图片类型会有此值
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

聊天历史数据单元。

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num				int16						//-- 实际返回的条目数量
		LastCursorId	nt64						//-- 继续轮询时，下次调用使用的 lastid 参数的值
		Begin			int64						//-- 继续轮询时，下次调用使用的 begin 参数的值
		End				int64						//-- 继续轮询时，下次调用使用的 end 参数的值
		Messages		[]*HistoryMessageUnit
	}

聊天历史返回结果。

### func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取群组聊天历史。

必选参数：

+ `desc bool`

	false: 从begin的时间戳开始，顺序翻页。  
	true： 从end的时间戳开始，倒序翻页。

+ `num int16`

	获取条目数量。建议10条，最多一次20条。

+ `begin int64`

	开始时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：>=

+ `end int64`

	结束时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：<=

+ `lastid int64`

	最后一条消息的id，第一次填默认0。条件：> 或者 <

+ `uid int64`

	用户id

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。


### func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取房间聊天历史。

必选参数：

+ `desc bool`

	false: 从begin的时间戳开始，顺序翻页。  
	true： 从end的时间戳开始，倒序翻页。

+ `num int16`

	获取条目数量。建议10条，最多一次20条。

+ `begin int64`

	开始时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：>=

+ `end int64`

	结束时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：<=

+ `lastid int64`

	最后一条消息的id，第一次填默认0。条件：> 或者 <

+ `uid int64`

	用户id

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取广播聊天历史。

必选参数：

+ `desc bool`

	false: 从begin的时间戳开始，顺序翻页。  
	true： 从end的时间戳开始，倒序翻页。

+ `num int16`

	获取条目数量。建议10条，最多一次20条。

+ `begin int64`

	开始时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：>=

+ `end int64`

	结束时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：<=

+ `lastid int64`

	最后一条消息的id，第一次填默认0。条件：> 或者 <

+ `uid int64`

	用户id

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetP2PChat(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetP2PChat(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取 P2P 聊天历史。

必选参数：

+ `desc bool`

	false: 从begin的时间戳开始，顺序翻页。  
	true： 从end的时间戳开始，倒序翻页。

+ `num int16`

	获取条目数量。建议10条，最多一次20条。

+ `begin int64`

	开始时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：>=

+ `end int64`

	结束时间戳，精确到**毫秒**，默认0。使用服务器当前时间。条件：<=

+ `lastid int64`

	最后一条消息的id，第一次填默认0。条件：> 或者 <

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。


### -----------------------[ 删除聊天记录 ]-----------------------------

### type HistoryMessageUnit

	type MessageType int

	const (
		MessageType_P2P MessageType = iota
		MessageType_Group
		MessageType_Room
		MessageType_Broadcast
	)

聊天类型定义。

### func (client *RTMServerClient) DelP2PChat(mid int64, fromUid int64, to int64, rest ... interface{}) error

	func (client *RTMServerClient) DelP2PChat(mid int64, fromUid int64, to int64, rest ... interface{}) error

删除/撤回聊天消息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelGroupChat(mid int64, fromUid int64, gid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupChat(mid int64, fromUid int64, gid int64, rest ... interface{}) error

删除/撤回组消息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelRoomChat(mid int64, fromUid int64, rid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomChat(mid int64, fromUid int64, rid int64, rest ... interface{}) error

删除/撤回房间消息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelBroadcastChat(mid int64, fromUid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelBroadcastChat(mid int64, fromUid int64, rest ... interface{}) error

删除/撤回广播消息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### -----------------------[ 翻译聊天系消息 ]-----------------------------

### type TranslateResult

	type TranslateResult struct {
		SourceLanguage		string
		TargetLanguage		string
		SourceText			string
		TargetText			string
	}

翻译聊天消息返回结果。

### func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string,
	textType string, profanity string, postProfanity bool, uid int64, rest ...interface{}) (result *TranslateResult, err error)

	func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string, textType string, profanity string, rest ... interface{}) (result *TranslateResult, err error)

翻译聊天消息。

必选参数：

+ `sourceLanguage string`

	源语言 ISO 639-1 代码。

	如果为空字符串，则系统将自动检测源语言种类。

+ `targetLanguage string`

	目标语言 ISO 639-1 代码。

+ `textType string`

	源数据类型。

	可选值：

		* chat：'\t'、'\n'、' ' 在输出文本中可能被修改；
		* mail：输出文本中，'\t'、'\n'、' ' 将保持不变。

	如果为空字符串，则默认为 `chat`。

+ `profanity string`

	是否触发敏感词过滤。

	可选值：

		* off：不做敏感词检测；
		* stop：当发现敏感词时，接口返回错误；
		* censor：当发现敏感词时，敏感词将被 `*` 替代。

	如果为空字符串，则默认为 `off`。

+ `postProfanity bool`

	是否把翻译后的文本过滤。


可接受的参数为：


+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (result *TranslateResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *TranslateResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *TranslateResult 结果，将通过 callback 传递。


### -----------------------[ 敏感词主动过滤 ]-----------------------------

### func (client *RTMServerClient) Profanity(text string, classify bool, uid int64, rest ... interface{}) (string, error)

	func (client *RTMServerClient) Profanity(text string, action string, rest ... interface{}) (string, error)

敏感词过滤。

Note: 

* maybe in after version this interface will be deprecated，recommend use TextCheck interface replace

必选参数：

+ `classify bool`

	是否进行文本分类检测。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (text string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" 及 error 信息。真实的 过滤后文本，将通过 callback 传递。


### -----------------------[ 语音转文字 ]-------------------------------

### func (client *RTMServerClient) Speech2Text(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (string, string, error)

	func (client *RTMServerClient) Speech2Text(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (string, string, error)

语音识别。

必选参数：

+ `audio []byte`

	音频的url或者内容（需要提供lang&codec&srate）

+ `audioType int32`

	语音数据类型. 1:url 2:内容

+ `lang RTMTranslateLanguage`

	语言

可选参数
+ `codec string`

	codec为空则默认为AMR_WB

+ `srate int32`

	srate为0或者空则默认为16000

+ `uid int64`

	用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (text string, lang string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 文本及lang及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" "" 及 error 信息。真实的文本，将通过 callback 传递。


### -----------------------[ 文本审核 ]-----------------------------

### func (client *RTMServerClient) TextCheck(text string, uid int64, rest ...interface{}) (int32, string, []int32, []string, error)

	func (client *RTMServerClient) TextCheck(text string, uid int64, rest ...interface{}) (result int32, text string, tags []int32, wlist []string, error)

文本审核。

必选参数：

+ `text string`

	文本内容

可选参数

+ `uid int64`

	用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (result int32, text string, tags []int32, wlist []string, errorCode int, errorMessage string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本及lang及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" "" 及 error 信息。真实的 过滤后文本，将通过 callback 传递。result: 0 通过 2 不通过；text: 敏感词过滤后的文本内容，含有的敏感词会被替换为*，如果没有被标星，则无此字段；tags: 触发的分类，比如涉黄涉政等等，具体见文本审核分类；wlist: 敏感词列表

说明：如果返回的result=2，正常处理是：如果text不为空则可以直接发出(用返回的text)，否则拦截（可能是广告或者隐晦色情等等）
注：如果需要详细的返回结果，请调用审核产品原生接口


### -----------------------[ 图片审核 ]-----------------------------

### func (client *RTMServerClient) ImageCheck(image string, imageType int32, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) ImageCheck(image string, imageType int32, uid int64, rest ...interface{}) (int32, []int32, error)

图片审核。

必选参数：

+ `image string`

	图片的url 或者内容

+ `imageType int32`

	1, url, 2, 内容

可选参数

+ `uid int64`

	用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本及lang及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" "" 及 error 信息。真实的 过滤后文本，将通过 callback 传递。result: 0 通过，2 不通过；tags：触发的分类，比如涉黄涉政等等，具体见图片审核分类

注：如果需要详细的返回结果，请调用审核产品原生接口


### -----------------------[ 视频审核 ]-----------------------------

### func (client *RTMServerClient) VideoCheck(video string, videoType int32, videoName string, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) VideoCheck(video string, videoType int32, videoName string, uid int64, rest ...interface{}) (int32, []int32, error)

视频审核。

必选参数：

+ `video string`

	视频的url或者内容

+ `videoType int32`

	1, url, 2, 内容

+ `videoName string`

	视频文件名，type=2时候必选，可以通过文件名获取文件格式

可选参数

+ `uid int64`

	用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本及lang及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" "" 及 error 信息。真实的结果，将通过 callback 传递。result: 0 通过，2 不通过；tags：触发的分类，比如涉黄涉政等等，具体见视频审核分类

注：如果需要详细的返回结果，请调用审核产品原生接口


### -----------------------[ 音频审核 ]-----------------------------

### func (client *RTMServerClient) AudioCheck(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (int32, []int32, error)

	func (client *RTMServerClient) AudioCheck(audio string, audioType int32, lang RTMTranslateLanguage, codec string, srate int32, uid int64, rest ...interface{}) (int32, []int32, error)

音频审核。

必选参数：

+ `audio []byte`

	音频的url或者内容（需要提供lang&codec&srate）

+ `audioType int32`

	语音数据类型. 1:url 2:内容

+ `lang RTMTranslateLanguage`

	语言

可选参数
+ `codec string`

	codec为空则默认为AMR_WB

+ `srate int32`

	srate为0或者空则默认为16000

+ `uid int64`

	用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (result int32, tags []int32, errorCode int, errorMessage string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本及lang及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" "" 及 error 信息。真实的结果，将通过 callback 传递。result: 0 通过，2 不通过；tags：触发的分类，比如涉黄涉政等等，具体见音频审核分类

注：如果需要详细的返回结果，请调用审核产品原生接口
