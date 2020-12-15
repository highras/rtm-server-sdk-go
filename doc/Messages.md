# RTM Server-End Go SDK Messages API Docs

# Index

[TOC]

### -----------------------[ 发送消息接口 ]-----------------------------

### func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessageByBinary(fromUid int64, toUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

发送 P2P 消息。

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

### func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessagesByBinary(fromUid int64, toUids []int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

发送多人 P2P 消息。

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

### func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupMessageByBinary(fromUid int64, groupId int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

发送群组消息。

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

### func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomMessageByBinary(fromUid int64, roomId int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

发送房间消息。

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

### func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastMessageByBinary(fromUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error)
	
发送广播消息。

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


### -----------------------[ 获取历史消息 ]-----------------------------

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
		FileInfo    	*FileMsgInfo
	}

### type HistoryMessageUnit

	type HistoryMessageUnit struct {
		CursorId int64
		RTMMessage
	}

历史消息数据单元。

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num				int16						//-- 实际返回的条目数量
		LastCursorId	nt64						//-- 继续轮询时，下次调用使用的 lastid 参数的值
		Begin			int64						//-- 继续轮询时，下次调用使用的 begin 参数的值
		End				int64						//-- 继续轮询时，下次调用使用的 end 参数的值
		Messages		[]*HistoryMessageUnit
	}

历史消息返回结果。

### func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取群组历史消息。

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

+ `mtypes []int8`

	指定获取的 mtype 类型。

+ `uid int64`

	用户id

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取房间历史消息。

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

+ `mtypes []int8`

	指定获取的 mtype 类型。

+ `uid int64`

	用户id

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取广播历史消息。

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

+ `mtypes []int8`

	指定获取的 mtype 类型。

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取 P2P 历史消息。

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

+ `mtypes []int8`

	指定获取的 mtype 类型。

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。

### -----------------------[ 删除消息记录 ]-----------------------------

### type HistoryMessageUnit

	type MessageType int

	const (
		MessageType_P2P MessageType = iota
		MessageType_Group
		MessageType_Room
		MessageType_Broadcast
	)

聊天类型定义。

### func (client *RTMServerClient) DelP2PMessage(mid int64, fromUid int64, to int64, rest ... interface{}) error

	func (client *RTMServerClient) DelP2PMessage(mid int64, fromUid int64, to int64, rest ... interface{}) error

删除/撤回P2P消息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelGroupMessage(mid int64, fromUid int64, gid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupMessage(mid int64, fromUid int64, gid int64, rest ... interface{}) error

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

### func (client *RTMServerClient) DelRoomMessage(mid int64, fromUid int64, rid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMessage(mid int64, fromUid int64, rid int64, rest ... interface{}) error

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

### func (client *RTMServerClient) DelBroadcastMessage(mid int64, fromUid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelBroadcastMessage(mid int64, fromUid int64, rest ... interface{}) error

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
