# RTM Server-End Go SDK Friends API Docs

# Index

[TOC]

### -----------------------[ Friendship ]-----------------------------

### func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error

	func (client *RTMServerClient) AddFriends(uid int64, firendUids []int64, rest ... interface{}) error

add friend. **Up to 100 people are added each time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error

	func (client *RTMServerClient) DelFriends(uid int64, firendUids []int64, rest ... interface{}) error

delete friend. **Up to 100 people are deleted each time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetFriends(uid int64, rest ... interface{}) ([]int64, error)

Get a list of friends.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(uids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the buddy list and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real friend list will be passed through callback.

### func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsFriend(uid int64, peerUid int64, rest ... interface{}) (bool, error)

Judge the friendship.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real friendship will be passed through callback.

### func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) IsFriends(uid int64, uids []int64, rest ... interface{}) ([]int64, error)

Judge the friendship. A maximum of 100 people are judged each time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(uids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the buddy list and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real friend list will be passed through callback.
