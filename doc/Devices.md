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
