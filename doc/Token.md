# RTM Server-End Go SDK Token API Docs

# Index

[TOC]


### func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

	func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

获取用户登陆 token。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(token string, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求，返回 token 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**请求，返回 "" 及 error 信息。真实的 token，将通过 callback 传递。


### func (client *RTMServerClient) RemoveToken(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveToken(uid int64, rest ... interface{}) error

删除用户登陆 token。

可接受的参数为：

+ `timeout time.Duration`

	请求超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**请求。  
如果 **callback** 参数**存在**，则为**异步**请求。
