# RTM Server-End Go SDK Groups API Docs

# Index

[TOC]

### -----------------------[ 群组关系接口 ]-----------------------------

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

### -----------------------[ 群组信息接口 ]-----------------------------

### func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

设置群组公开信息和私有信息。

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

### func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ... interface{}) (string, string, error)

获取群组公开信息和私有信息。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 群组公开信息、群组私有信息、error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 ""、""、error 信息。真实的 群组公开信息和私有信息，将通过 callback 传递。
