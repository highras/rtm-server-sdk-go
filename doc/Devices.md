# RTM Server-End Go SDK Devices API Docs

# Index

[TOC]

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

### func (client *RTMServerClient) AddDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

	func (client *RTMServerClient) AddDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

设置设备推送属性

必选参数：

+ `pushType int8`

	pushType**可选为(0,1)**, pushType == 0,设置某个p2p不推送，pushType == 1, 设置某个group不推送

+ `xid int64`

	当pushType为0时，xid为from，当pushType为1时，xid为groupId

+ `mtype []int8`

	mtype为nil或者为空，则为所有mtype均不推送，其他值，则指定mtype不推送，注意聊天涉及到文本、cmd，文件几个mtype

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) RemoveDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

	func (client *RTMServerClient) RemoveDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

删除设备推送属性

必选参数：

+ `pushType int8`

	pushType**可选为(0,1)**, pushType == 0,删除某个p2p不推送，pushType == 1, 删除某个group不推送

+ `xid int64`

	当pushType为0时，xid为from，当pushType为1时，xid为groupId

+ `mtype []int8`

	mtype为nil或者为空，则为取消所有mtype均不推送，其他值，则取消指定mtype不推送，注意聊天涉及到文本、cmd，文件几个mtype

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) GetDevicePushOption(uid int64, rest ...interface{}) (map[int64][]int8, map[int64][]int8, error)

	func (client *RTMServerClient) GetDevicePushOption(uid int64, rest ...interface{}) (map[int64][]int8, map[int64][]int8, error)

获取设备推送属性。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(map[int64][]int8, map[int64][]int8, int, string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回p2p的不推送设置key为from,value为mtypes, value为0，则代表所有mtype不推送，返回group的不推送设置key为groupId， value为mtypes, value为0，则代表所有mtype不推送。  
如果 **callback** 参数**存在**，则为**异步**请求。，返回 nil、nil及 error 信息。真实的设备推送属性，将通过 callback 传递。
