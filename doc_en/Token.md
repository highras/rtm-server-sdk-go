# RTM Server-End Go SDK Token API Docs

# Index

[TOC]


### func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

	func (client *RTMServerClient) GetToken(uid int64, rest ... interface{}) (string, error)

Get user login token.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(token string, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and token and error information is returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "" and error information. The real token will be passed through callback.


### func (client *RTMServerClient) RemoveToken(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveToken(uid int64, rest ... interface{}) error

Delete the user login token.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.
