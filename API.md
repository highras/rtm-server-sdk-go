# RTM Server-End Go SDK API Docs

# Package rtm

	import "github.com/highras/rtm-server-sdk-go/src/rtm"

rtm 包提供go连接和访问 RTM 后端服务的能力。

# Index

[TOC]

## Constants

	const SDKVersion = "0.1.0"

## Variables

FPNN 的 Config 对象，将会影响 rtm 客户端的行为。  
具体请参见：[FPNN Go SDK - API Docs](https://github.com/highras/fpnn-sdk-go/blob/master/API.md)

## type RTMServerMonitor

	type RTMServerMonitor interface {
		P2PMessage(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
		GroupMessage(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64)
		RoomMessage(fromUid int64, roomIid int64, mtype int8, mid int64, message string, attrs string, mtime int64)
		Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
	}

消息监听接口。

请通过 RTM Console 进行配置设置，并在建立连接后，调用 `func AddListen(...)` 或 `func SetListen(...)` 进行代码设置。


## type RTMServerClient

	type RTMServerClient struct {
		//-- same hidden fields
	}

RTM Server 客户端。


### func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

	func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

创建 RTM Server 客户端。

所有参数请通过 RTM Console 获取。

### -----------------------[ 配置接口 ]-----------------------------

### func (client *RTMServerClient) SetMonitor(monitor RTMServerMonitor)

	func (client *RTMServerClient) SetMonitor(monitor RTMServerMonitor)

配置消息监听接口。
具体参考：[RTMServerMonitor](#type-RTMServerMonitor)

### func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

配置 RTM Server Client 的连接超时。  
未配置时，默认采用 fpnn.Config 的连接超时参数。

### func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

配置 RTM Server Client 的请求超时。  
未配置时，默认采用 fpnn.Config 的请求超时参数。

### func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

	func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

配置连接建立事件的回调函数。

### func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

	func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

配置连接断开事件的回调函数。

### func (client *RTMServerClient) SetLogger(logger *log.Logger)

	func (client *RTMServerClient) SetLogger(logger *log.Logger)

配置 RTM Server Client 的日志路由。

### func (client *RTMServerClient) Endpoint() string

	func (client *RTMServerClient) Endpoint() string

获取 RTM Server Client 连接/目标地址。

### func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

	func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

配置使用加密链接。

可接受的参数为：

	+ `pemKeyPath string`

		服务器公钥文件路径。PEM 格式。与 pemKeyData 参数互斥。

	+ `pemKeyData []byte`

		服务器公钥文件内容。PEM 格式。与 pemKeyPath 参数互斥。

	+ `reinforce bool`

		true 采用 256 位密钥加密，false 采用 128 位密钥加密。  
		默认为 true


### -----------------------[ 发送消息接口 ]-----------------------------

### func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendMessage(fromUid int64, toUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

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

### func (client *RTMServerClient) SendBoradcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBoradcastMessage(fromUid int64, mtype int8, message string, rest ... interface{}) (int64, error)

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


### -----------------------[ 历史消息 ]-----------------------------

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

历史消息数据单元。

### type HistoryMessageResult

	type HistoryMessageResult struct {
		Num			int16						//-- 实际返回的条目数量
		LastId		int64						//-- 继续轮询时，下次调用使用的 lastid 参数的值
		Begin		int64						//-- 继续轮询时，下次调用使用的 begin 参数的值
		End			int64						//-- 继续轮询时，下次调用使用的 end 参数的值
		Messages	[]*HistoryMessageUnit
	}

历史消息返回结果。

### func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetGroupMessage(groupId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetRoomMessage(roomId int64, desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。



### func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

	func (client *RTMServerClient) GetBroadcastMessage(desc bool, num int16, begin int64, end int64, lastid int64, rest ... interface{}) (*HistoryMessageResult, error)

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

可接受的参数为：

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

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(result *HistoryMessageResult, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 *HistoryMessageResult 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 *HistoryMessageResult 结果，将通过 callback 传递。


### -----------------------[ 好友关系 ]-----------------------------

### func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error

	func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error

添加好友。每次**最多**添加100人。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error

	func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error

删除好友。每次**最多**删除100人。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error)

获取好友列表。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(uids []int64, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 好友列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 好友列表，将通过 callback 传递。

### func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error)

判断好友关系。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(ok bool, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 bool 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 false 及 error 信息。真实的 好友关系，将通过 callback 传递。

### func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error)

判断好友关系。每次**最多**判断100人。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(uids []int64, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 好友列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 好友列表，将通过 callback 传递。


### -----------------------[ 群组接口 ]-----------------------------


### func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

添加群组成员。每次**最多**添加100人。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

删除群组成员。每次**最多**删除100人。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error

删除群组。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error)

获取群组成员。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(uids []int64, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 群组成员列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 群组成员列表，将通过 callback 传递。

### func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error)

判断群组关系。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(ok bool, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 bool 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 false 及 error 信息。真实的 群组关系，将通过 callback 传递。

### func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error)

获取用户加入的群组。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(groupIds []int64, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 群组列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 群组列表，将通过 callback 传递。


### -----------------------[ 房间接口 ]-----------------------------

### func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error

添加房间成员。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

删除房间成员。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### -----------------------[ 管理接口 ]-----------------------------

### func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

禁止用户指定群组内发言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error

解除用户指定群组内禁言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

禁止用户指定房间内发言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error
	
	func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error

解除用户指定房间内禁言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ... interface{}) error

将用户加入项目黑名单。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ... interface{}) error

将用户移出项目黑名单。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error)

判断用户是否在指定群组中被禁言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(ok bool, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 bool 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 false 及 error 信息。真实的 禁言状态，将通过 callback 传递。

### func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error)

判断用户是否在指定房间中被禁言。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(ok bool, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 bool 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 false 及 error 信息。真实的 禁言状态，将通过 callback 传递。


### func (client *RTMServerClient) IsProjectBlack(uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsProjectBlack(uid int64, rest ... interface{}) (bool, error)

判断用户是否在项目黑名单中。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(ok bool, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 bool 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 false 及 error 信息。真实的 黑名单状态，将通过 callback 传递。

### func (client *RTMServerClient) DelMessage(mid int64, fromUid int64, xid int64, messageType MessageType, rest ... interface{}) error

	func (client *RTMServerClient) DelMessage(mid int64, fromUid int64, xid int64, messageType MessageType, rest ... interface{}) error

删除/撤回消息。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) Kickout(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) Kickout(uid int64, rest ... interface{}) error

踢用户下线。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### -----------------------[ 杂项工具接口 ]-----------------------------

### func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

	func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

获取用户登陆 token。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(token string, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 token 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" 及 error 信息。真实的 token，将通过 callback 传递。

### func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ... interface{}) ([]int64, error)

获取在线用户列表。每次**最多**获取200个。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(uids []int64, errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 在线用户列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 在线用户列表，将通过 callback 传递。

### func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error

	func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error

**增量**添加监听设置。

必须参数：

	+ `groupIds []int64`

		增加监听的群组。

	+ `roomIds []int64`

		增加监听的房间。

	+ `p2p bool`

		当 true 时，监听 P2P 消息。忽略 false 参数。

	+ `events []string`

		需要监听的事件。  
		可监听的事件列表，请参考 RTM 服务文档。


可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error

	func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error

**增量**取消监听设置。

必须参数：

	+ `groupIds []int64`

		取消监听的群组。

	+ `roomIds []int64`

		取消监听的房间。

	+ `p2p bool`

		当 true 时，取消监听 P2P 消息。忽略 false 参数。

	+ `events []string`

		需要取消监听的事件。  
		可监听的事件列表，请参考 RTM 服务文档。


可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error

	func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, p2p bool, events []string, rest ... interface{}) error 

设置监听状态。该接口将**覆盖**以前的设置。

必须参数：

	+ `groupIds []int64`

		取消监听的群组。

	+ `roomIds []int64`

		取消监听的房间。

	+ `p2p bool`

		当 true 时，取消监听 P2P 消息。忽略 false 参数。

	+ `events []string`

		需要取消监听的事件。  
		可监听的事件列表，请参考 RTM 服务文档。


可接受的参数为：

	+ `all bool`

		true: 忽略其他参数，监听所有消息，所有事件。  
		false: 忽略其他参数，取消监听所有消息，所有事件。

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ... interface{}) error

	func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ... interface{}) error

添加设备&应用信息。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ... interface{}) error

	func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ... interface{}) error

删除设备&应用信息。

可接受的参数为：

	+ `timeout time.Duration`

		请求超时。  
		缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
		若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

	+ `callback func(errorCode int, errInfo string)`

		异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### -----------------------[ 文件接口 ]-----------------------------

### func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送 P2P 文件。

可接受的参数为：

	+ `mtype int8`

		消息类型。默认 50。

	+ `extension string`

		文件的扩展名。  
		如果缺失，会自动从 filename 参数中提取。

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

### func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送多人 P2P 文件。

可接受的参数为：

	+ `mtype int8`

		消息类型。默认 50。

	+ `extension string`

		文件的扩展名。  
		如果缺失，会自动从 filename 参数中提取。

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

### func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送群组文件。

可接受的参数为：

	+ `mtype int8`

		消息类型。默认 50。

	+ `extension string`

		文件的扩展名。  
		如果缺失，会自动从 filename 参数中提取。

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

### func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送房间文件。

可接受的参数为：

	+ `mtype int8`

		消息类型。默认 50。

	+ `extension string`

		文件的扩展名。  
		如果缺失，会自动从 filename 参数中提取。

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

### func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送广播文件。

可接受的参数为：

	+ `mtype int8`

		消息类型。默认 50。

	+ `extension string`

		文件的扩展名。  
		如果缺失，会自动从 filename 参数中提取。

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
