# RTM Server-End Go SDK Devices API Docs

# Index

[TOC]

### func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ... interface{}) error

	func (client *RTMServerClient) AddDevice(uid int64, appType string, deviceToken string, rest ... interface{}) error

Add device & application information.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ... interface{}) error

	func (client *RTMServerClient) RemoveDevice(uid int64, deviceToken string, rest ... interface{}) error

Delete device & application information.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.
