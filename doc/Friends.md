# RTM Server-End Go SDK Friends API Docs

# Index

[TOC]

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
