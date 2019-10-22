# RTM Server-End Go SDK Users API Docs

# Index

[TOC]

### -----------------------[ 用户信息接口 ]-----------------------------

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

### func (client *RTMServerClient) SetUserInfo(uid int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetUserInfo(uid int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

设置用户公开信息和私有信息。

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

### func (client *RTMServerClient) GetUserInfo(uid int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetUserInfo(uid int64, rest ... interface{}) (string, string, error)

获取用户公开信息和私有信息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 用户公开信息、用户私有信息、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 ""、""、error 信息。真实的 用户公开信息和私有信息，将通过 callback 传递。


### func (client *RTMServerClient) GetUserPublicInfo(uids []int64, rest ... interface{}) (map[string]string, error)

	func (client *RTMServerClient) GetUserPublicInfo(uids []int64, rest ... interface{}) (map[string]string, error)

批量获取用户公开信息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (map[string]string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 用户公开信息 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 nil 及 error 信息。真实的 用户公开信息，将通过 callback 传递。

### -----------------------[ 管理接口 ]-----------------------------

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
