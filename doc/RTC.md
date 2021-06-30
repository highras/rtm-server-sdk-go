# RTM Server-End Go SDK RealTimeRTC API Docs

# Index

[TOC]

### -----------------------[ 实时音视频房间管理接口 ]-----------------------------

### func (client *RTMServerClient) InviteUserIntoRTCRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) InviteUserIntoRTCRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

邀请用户加入RTC房间。

参数说明：

+ `fromUid int64`: 
	
	发起邀请的用户(必须在RTC房间里才能发起邀请指令) 

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

### func (client *RTMServerClient) PullUserIntoRTCRoom(roomId int64, uids []int64, roomType int32, rest ...interface{}) error

	func (client *RTMServerClient) PullUserIntoRTCRoom(roomId int64, uids []int64, roomType int32, rest ...interface{}) error

强拉用户进入RTC房间(房间不存在会自动创建该房间)。

可接受的参数为：

+ `roomType int32`

	房间类型 1 voice, 2 video

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。


### func (client *RTMServerClient) CloseRTCRoom(roomId int64, rest ...interface{}) error

	func (client *RTMServerClient) CloseRTCRoom(roomId int64, rest ...interface{}) error

关闭RTC房间

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) KickoutFromRTCRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error
	
	func (client *RTMServerClient) KickoutFromRTCRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

从RTC房间里踢出。

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

### func (client *RTMServerClient) SetRTCRoomMicStatus(roomId int64, status bool, rest ...interface{}) error
	
	func (client *RTMServerClient) SetRTCRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

设置RTC房间默认麦克风状态(用户刚进入房间默认是否开启麦克风)。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### -----------------------[ 实时RTC房间信息查询接口 ]-----------------------------

### func (client *RTMServerClient) GetRTCRoomList(rest ...interface{}) ([]int64, error))

	func (client *RTMServerClient) GetRTCRoomList(rest ...interface{}) (rids []int64, error)

获取该项目当前的RTC房间id列表

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(rids []int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 RTC房间id列表 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 RTC房间id列表，将通过 callback 传递。

### func (client *RTMServerClient) GetRTCRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, int64, error)

	func (client *RTMServerClient) GetRTCRoomMembers(roomId int64, rest ...interface{}) (uids []int64, managers []int64, owner int64, error)

获取RTC房间当前的用户id列表以及管理员id列表和房主id。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (uids []int64, managers []int64, owner int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。返回当前RTC房间用户id列表以及管理员id列表和房主id及error信息 
如果 **callback** 参数**存在**，则为**异步**请求。真实的数据通过callback返回

### func (client *RTMServerClient) GetRTCRoomMemberCount(roomId int64, rest ...interface{}) (int32, error)

	func (client *RTMServerClient) GetRTCRoomMemberCount(roomId int64, rest ...interface{}) (count int32, error)

获取当前RTC房间人数。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (count int32, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回当前房间人数、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，真实的房间人数将通过 callback 返回。

### func (client *RTMServerClient) AdminCommand(roomId int64, uids []int64, command int32, rest ...interface{}) error

	func (client *RTMServerClient) AdminCommand(roomId int64, uids []int64, command int32, rest ...interface{}) error

房间管理员操作

可接受的参数为：

+ `uids []int64` 
	
	操作的用户id列表

+ `command int32`

	操作类型：0 赋予管理员权限，1 剥夺管理员权限，2 禁止发送音频数据，3 允许发送音频数据，4 禁止发送视频数据，5 允许发送视频数据，6 关闭他人麦克风，7 关闭他人摄像头

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，通过 callback 返回。
