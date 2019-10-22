# RTM Server-End Go SDK Data API Docs

# Index

[TOC]

### -----------------------[ 用户数据接口 ]-----------------------------

### func (client *RTMServerClient) SetData(uid int64, key string, value string, rest ... interface{}) error

	func (client *RTMServerClient) SetData(uid int64, key string, value string, rest ... interface{}) error

存储用户数据。

必须参数：

+ `key string`

	最大 128 字节。

+ `value string`

	最大 65535 字节。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client *RTMServerClient) GetData(uid int64, key string, rest ... interface{}) (string, error)

	func (client *RTMServerClient) GetData(uid int64, key string, rest ... interface{}) (string, error)

获取用户存储的数据。

必须参数：

+ `key string`

	最大 128 字节。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func (text string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 用户存储的数据（文本）及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" 及 error 信息。真实的 用户存储的数据，将通过 callback 传递。

### func (client *RTMServerClient) DelData(uid int64, key string, rest ... interface{}) error

	func (client *RTMServerClient) DelData(uid int64, key string, rest ... interface{}) error

删除用户数据。

必须参数：

+ `key string`

	最大 128 字节。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。
