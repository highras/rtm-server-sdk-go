# RTM Server-End Go SDK RealTimeVoice API Docs

# Index

[TOC]

### -----------------------[ 实时语音房间管理接口 ]-----------------------------

### func (client *RTMServerClient) InviteUserIntoVoiceRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) InviteUserIntoVoiceRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

邀请用户加入语音房间。

参数说明：

+ `fromUid int64`: 
	
	发起邀请的用户(必须在语音房间里才能发起邀请指令) 

+ `uids []int64`: 
	
	被邀请的用户列表

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) PullUserIntoVoiceRoom(roomId int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) PullUserIntoVoiceRoom(roomId int64, uids []int64, rest ...interface{}) error

强拉用户进入语音房间(房间不存在会自动创建该房间)。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) CloseVoiceRoom(roomId int64, rest ...interface{}) error

	func (client *RTMServerClient) CloseVoiceRoom(roomId int64, rest ...interface{}) error

关闭语音房间

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) KickoutFromVoiceRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error
	
	func (client *RTMServerClient) KickoutFromVoiceRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

从语音房间里踢出。

参数说明：

+ `uid int64`: 

	被踢的用户

+ `fromUid int64`:

	发起踢人指令的用户

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) SetVoiceRoomMicStatus(roomId int64, status bool, rest ...interface{}) error
	
	func (client *RTMServerClient) SetVoiceRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

设置语音房间默认麦克风状态(用户刚进入房间默认是否开启麦克风)。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### -----------------------[ 实时语音房间信息查询接口 ]-----------------------------

### func (client *RTMServerClient) GetVoiceRoomList(rest ...interface{}) ([]int64, error))

	func (client *RTMServerClient) GetVoiceRoomList(rest ...interface{}) (rids []int64, error)

获取该项目当前的语音房间id列表

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(rids []int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 语音房间id列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 语音房间id列表，将通过 callback 传递。

### func (client *RTMServerClient) GetVoiceRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, error)

	func (client *RTMServerClient) GetVoiceRoomMembers(roomId int64, rest ...interface{}) (uids []int64, managers []int64, error)

获取语音房间当前的用户id列表以及管理员id列表。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (uids []int64, managers []int64 errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。返回当前语音房间用户id列表以及管理员id列表及error信息 
如果 **callback** 参数**存在**，则为**异步**请求。真实的数据通过callback返回

### func (client *RTMServerClient) GetVoiceRoomMemberCount(roomId int64, rest ...interface{}) (int32, error)

	func (client *RTMServerClient) GetVoiceRoomMemberCount(roomId int64, rest ...interface{}) (count int32, error)

获取当前语音房间人数。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (count int32, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回当前房间人数、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，真实的房间人数将通过 callback 返回。
