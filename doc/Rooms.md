# RTM Server-End Go SDK Rooms API Docs

# Index

[TOC]

### -----------------------[ 房间关系接口 ]-----------------------------

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

### func (client *RTMServerClient) AddUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

	func (client *RTMServerClient) AddUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

用户加入多个房间。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) DeleteUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

	func (client *RTMServerClient) DeleteUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

用户离开多个房间。

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

### func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

禁止用户指定房间内发言。

参数说明：

+ `roomId int64`:

	roomId <= 0时，则对所有房间禁言

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

参数说明：

+ `roomId int64`:

	roomId <= 0时，则对所有房间解除禁言

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

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

### -----------------------[ 房间信息接口 ]-----------------------------

### func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

设置房间公开信息和私有信息。

必须参数：

+ `publicInfo *string`

	需要设置的公开信息。`nil` 表示本次调用不操作公开信息。最大 65535 字节。

+ `privateInfo *string`

	需要设置的私有信息。`nil` 表示本次调用不操作私有信息。最大 65535 字节。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ... interface{}) (string, string, error)

获取房间公开信息和私有信息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 房间公开信息、房间私有信息、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 ""、""、error 信息。真实的 房间公开信息和私有信息，将通过 callback 传递。

### func (client *RTMServerClient) GetRoomMembers(roomId int64, rest ...interface{}) ([]int64, error)

	func (client *RTMServerClient) GetRoomMembers(roomId int64, rest ...interface{}) ([]int64, error)

获取房间所有成员。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (uids []int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 房间成员、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 uids []int64、error 信息。真实的房间成员，将通过 callback 传递。

### func (client *RTMServerClient) GetRoomCount(roomId []int64, rest ...interface{}) (map[int64]int32, error)

	func (client *RTMServerClient) GetRoomCount(roomId []int64, rest ...interface{}) (map[int64]int32, error)

获取房间中用户数量

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(map[int64]int32, int, string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 房间中用户数量、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 count map[int64]int32、error 信息。真实的 房间中用户数量，将通过 callback 传递。
