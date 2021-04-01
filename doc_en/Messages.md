# RTM Server-End Go SDK Messages API Docs

# Index

[TOC]

### -----------------------[ Send message interface ]-----------------------------

### func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessageByBinary(fromUid int64, toUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

Send P2P messages.

The acceptable parameters are:

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be used

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessages(fromUid int64, toUids []int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessagesByBinary(fromUid int64, toUids []int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

Send multi-person P2P messages.

The acceptable parameters are:

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be used

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupMessageByBinary(fromUid int64, groupId int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

Send group messages.

The acceptable parameters are:

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be used

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomMessageByBinary(fromUid int64, roomId int64, messageType int8, message []byte, rest ...interface{}) (int64, error)

Send room message

The acceptable parameters are:

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be used

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastMessageByBinary(fromUid int64, messageType int8, message []byte, rest ...interface{}) (int64, error)
	
Send broadcast message

The acceptable parameters are:

+ `attrs string`

	The attribute information of the message. The default is an empty string.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be used.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.


### -----------------------[ Get message historical ]-----------------------------

### type RTMAudioFileInfo

	type RTMAudioFileInfo struct {
		IsRTMaudio bool   // Is it an rtm voice message
		Codec      string // There is this value when rtm voice message coding format
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
		FileInfo    	*FileMsgInfo
	}

### type HistoryMessageUnit

	type HistoryMessageUnit struct {
		CursorId int64
		RTMMessage
	}

Historical message data unit.

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num				int16						//-- Actual number of entries returned
		LastCursorId	int64						//-- When the polling continues, the value of the lastid parameter used in the next call
		Begin			int64						//-- When continuing to poll, the value of the begin parameter used in the next call
		End				int64						//-- When continuing to poll, the value of the end parameter used in the next call
		Messages		[]*HistoryMessageUnit
	}

Historical messages return results.

### func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get group history messages.

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

The acceptable parameters are:

+ `mtypes []int8`

	Specify the obtained mtype type.

+ `uid int64`

	User id

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.



### func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get room history information.

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

The acceptable parameters are:

+ `mtypes []int8`

	Specify the obtained mtype type.

+ `uid int64`

	user id

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.



### func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, uid int64, rest ... interface{}) (*HistoryMessageResult, error)

Get broadcast history messages.

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

The acceptable parameters are:

+ `mtypes []int8`

	Specify the obtained mtype type.

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.



### func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetP2PMessage(uid int64, peerUid int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

获取 P2P 历史消息。

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

The acceptable parameters are:

+ `mtypes []int8`

	Specify the obtained mtype type.

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and *HistoryMessageResult and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real *HistoryMessageResult result will be passed through callback.

### -----------------------[ Delete message record ]-----------------------------

### type HistoryMessageUnit

	type MessageType int

	const (
		MessageType_P2P MessageType = iota
		MessageType_Group
		MessageType_Room
		MessageType_Broadcast
	)

Chat type definition.

### func (client *RTMServerClient) DelP2PMessage(mid int64, fromUid int64, to int64, rest ... interface{}) error

	func (client *RTMServerClient) DelP2PMessage(mid int64, fromUid int64, to int64, rest ... interface{}) error

Delete/withdraw P2P messages.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelGroupMessage(mid int64, fromUid int64, gid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupMessage(mid int64, fromUid int64, gid int64, rest ... interface{}) error

Delete/withdraw group messages.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelRoomMessage(mid int64, fromUid int64, rid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMessage(mid int64, fromUid int64, rid int64, rest ... interface{}) error

Delete/withdraw room messages.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelBroadcastMessage(mid int64, fromUid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelBroadcastMessage(mid int64, fromUid int64, rest ... interface{}) error

Delete/withdraw broadcast messages.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetMsgCount(msgType MessageType, xid int64, begin int64, end int64, mtype []int8, rest ...interface{}) (sender int32, count int32, err error)

	func (client *RTMServerClient) GetMsgCount(msgType MessageType, xid int64, begin int64, end int64, mtype []int8, rest ...interface{}) (sender int32, count int32, err error)

Get statistics of messages sent in a room or group

Note：
	Only the saved messages will be counted. The current chat messages are saved by default, plus the message type configured by the user.

Required parameters:

+ `msgType MessageType`: 
	
	Get the type of message, **accept rtm.MessageType_Group, rtm.MessageType_Room**

+ `xid int64`: 
	
	When msgType is **rtm.MessageType_Group**, it is **groupId**; when msgType is **rtm.MessageType_Room**, it is **roomId**

+ `begin int64`: 
	
	Millisecond timestamp, start time, if **0**, the time is ignored

+ `end int64`: 
	
	Millisecond timestamp, end time, if **0**, the time is ignored

+ `mtype []int8`: 

	If mtype is nil or empty, return all

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (sender int32, count int32, errorCode int, errInfo string)`

	Asynchronous callback function.   

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning the number of people who sent messages in the room or group (after deduplication), and the number of messages
If the **callback** parameter **exists**, it is an **asynchronous** request. The returned result will be returned through callback
