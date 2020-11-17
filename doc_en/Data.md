# RTM Server-End Go SDK Data API Docs

# Index

[TOC]

### -----------------------[ User data interface ]-----------------------------

### func (client *RTMServerClient) SetData(uid int64, key string, value string, rest ... interface{}) error

	func (client *RTMServerClient) SetData(uid int64, key string, value string, rest ... interface{}) error

 Save user data interface

Required parameters:

+ `key string`

	Maximum 128 bytes.

+ `value string`

	Maximum 65535 bytes.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exist**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetData(uid int64, key string, rest ... interface{}) (string, error)

	func (client *RTMServerClient) GetData(uid int64, key string, rest ... interface{}) (string, error)

Get the data stored by the user.

Required parameters:

+ `key string`

	Maximum 128 bytes.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (text string, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, which returns the data (text) and error information stored by the user.
If the **callback** parameter **exist**, it is an **asynchronous** request and returns "" and error information. The real data stored by the user will be passed through the callback.

### func (client *RTMServerClient) DelData(uid int64, key string, rest ... interface{}) error

	func (client *RTMServerClient) DelData(uid int64, key string, rest ... interface{}) error

Delete user data.

Required parameters:

+ `key string`

	Maximum 128 bytes.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exist**, it is an **asynchronous** request.
