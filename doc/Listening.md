# RTM Server-End Go SDK Listening API Docs

# Index

[TOC]

## type IRTMServerMonitor

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

消息监听接口。

请通过 RTM Console 进行配置设置，并在建立连接后，调用 `func AddListen(...)` 或 `func SetListen(...)` 进行代码设置。


### func (client *RTMServerClient) SetServerPushMonitor(monitor IRTMServerMonitor)

	func (client *RTMServerClient) SetServerPushMonitor(monitor IRTMServerMonitor)

配置消息监听接口。
具体参考：[IRTMServerMonitor](#type-IRTMServerMonitor)


### func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

**增量**添加监听设置。

必须参数：

+ `groupIds []int64`

	增加监听的群组。

+ `roomIds []int64`

	增加监听的房间。

+ `uids []int64`

	增加监听的 P2P 用户。

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


### func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

**增量**取消监听设置。

必须参数：

+ `groupIds []int64`

	取消监听的群组。

+ `roomIds []int64`

	取消监听的房间。

+ `uids []int64`

	取消监听的 P2P 用户。

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

### func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error 

设置监听状态。该接口将**覆盖**以前的设置。

必须参数：

+ `groupIds []int64`

	设置监听的群组。

+ `roomIds []int64`

	设置监听的房间。

+ `uids []int64`

	设置监听的 P2P 用户。

+ `events []string`

	设置监听的事件。  
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


### func (client *RTMServerClient) SetListenStatus(allGroups bool, allRrooms bool, allP2P bool, allEvents bool, rest ... interface{}) error

	func (client *RTMServerClient) SetListenStatus(allGroups bool, allRrooms bool, allP2P bool, allEvents bool, rest ... interface{}) error 

设置监听状态。该接口将**覆盖**以前的设置。

必须参数：

+ `allGroups bool`

	设置是否监听所有群组。

+ `allRrooms bool`

	设置是否监听所有房间。

+ `allP2P bool`

	设置是否监听所有的 P2P 消息。

+ `allEvents bool`

	设置是否监听所有的事件。  
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
