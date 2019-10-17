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


### -----------------------[ 获取聊天历史记录 ]-----------------------------

### type HistoryMessageUnit

	type HistoryMessageUnit struct {
		Id			int64
		FromUid		int64
		MType		int8
		Mid			int64
		Deleted		bool						//-- 是否已被删除/撤销
		Message		string
		Attrs		string
		MTime		int64
	}

聊天历史数据单元。

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num			int16						//-- 实际返回的条目数量
		LastId		int64						//-- 继续轮询时，下次调用使用的 lastid 参数的值
		Begin		int64						//-- 继续轮询时，下次调用使用的 begin 参数的值
		End			int64						//-- 继续轮询时，下次调用使用的 end 参数的值
		Messages	[]*HistoryMessageUnit
	}

聊天历史返回结果。

### func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupChat(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。


### func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomChat(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastChat(desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

### func (client *RTMServerClient) DelChat(mid int64, fromUid int64, xid int64, messageType MessageType, rest ... interface{}) error

	func (client *RTMServerClient) DelChat(mid int64, fromUid int64, xid int64, messageType MessageType, rest ... interface{}) error

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

### -----------------------[ 翻译聊天系消息 ]-----------------------------

### type TranslateResult

	type TranslateResult struct {
		SourceLanguage		string
		TargetLanguage		string
		SourceText			string
		TargetText			string
	}

翻译聊天消息返回结果。

### func (client *RTMServerClient) Translate(text string, sourceLanguage string, targetLanguage string, textType string, profanity string, rest ... interface{}) (result *TranslateResult, err error)

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

### func (client *RTMServerClient) Profanity(text string, action string, rest ... interface{}) (string, error)

	func (client *RTMServerClient) Profanity(text string, action string, rest ... interface{}) (string, error)

敏感词过滤。

必选参数：

+ `action string`

	是否触发敏感词过滤。

	可选值：

		* stop：当发现敏感词时，接口返回错误；
		* censor：当发现敏感词时，敏感词将被 `*` 替代。

	如果为空字符串，则默认为 `censor`。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (text string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 过滤后的文本 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" 及 error 信息。真实的 过滤后文本，将通过 callback 传递。

