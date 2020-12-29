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

### func (client *RTMServerClient) AddDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

	func (client *RTMServerClient) AddDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

Set device push properties

Required parameters:

+ `pushType int8`

	pushType** can be selected as (0,1)**, pushType == 0, set a p2p not to push, pushType == 1, set a group not to push

+ `xid int64`

	When pushType is 0, xid is from, when pushType is 1, xid is groupId

+ `mtype []int8`

	If mtype is nil or empty, all mtypes will not be pushed. For other values, mtype will not be pushed. Note that the chat involves text, cmd, and files.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.。

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronous** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) RemoveDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

	func (client *RTMServerClient) RemoveDevicePushOption(uid int64, pushType int8, xid int64, mtype []int8, rest ...interface{}) error 

Delete device push attributes

Required parameters:

+ `pushType int8`

	pushType** can be selected as (0,1)**, pushType == 0, cancel a p2p without pushing, pushType == 1, cancel a group without pushing

+ `xid int64`

	When pushType is 0, xid is from, when pushType is 1, xid is groupId

+ `mtype []int8`

	If mtype is nil or empty, all mtypes will not be pushed. For other values, the specified mtype will not be pushed. Note that the chat involves text, cmd, and files.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronous** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetDevicePushOption(uid int64, rest ...interface{}) (map[int64][]int8, map[int64][]int8, error)

	func (client *RTMServerClient) GetDevicePushOption(uid int64, rest ...interface{}) (map[int64][]int8, map[int64][]int8, error)

Get device push properties.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(map[int64][]int8, map[int64][]int8, int, string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the return p2p is not pushed. Set the key to from, value to mtypes, and value to 0, which means that all mtypes are not pushed and return to group If you do not push, set the key to groupId, value to mtypes, and value to 0, which means that all mtypes will not be pushed.  
If the **callback** parameter **exists**, it is an **asynchronous** request. , Return nil, nil and error information. The real device push attributes will be passed through callback.
