# RTM Server-End Go SDK Users API Docs

# Index

[TOC]

### -----------------------[ User information interface ]-----------------------------

### func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetOnlineUsers(uids []int64, rest ... interface{}) ([]int64, error)

Get a list of online users. Get up to 200 at a time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(uids []int64, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and a list of online users and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The list of real online users will be passed through callback.

### func (client *RTMServerClient) SetUserInfo(uid int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetUserInfo(uid int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

Set user public information and private information.

Required parameters:

+ `publicInfo *string`

	The public information that needs to be set. `nil` means that this call does not operate public information. Maximum 65535 bytes.

+ `privateInfo *string`

	The private information that needs to be set. `nil` means that this call does not operate private information. Maximum 65535 bytes.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetUserInfo(uid int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetUserInfo(uid int64, rest ... interface{}) (string, string, error)

Obtain public and private information from users.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning user public information, user private information, and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "", "", error information. Real user public information and private information will be passed through callback.


### func (client *RTMServerClient) GetUserPublicInfo(uids []int64, rest ... interface{}) (map[string]string, error)

	func (client *RTMServerClient) GetUserPublicInfo(uids []int64, rest ... interface{}) (map[string]string, error)

Get user public information in batches.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (map[string]string, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning user public information and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real user public information will be passed through callback.

### -----------------------[ Management interface ]-----------------------------

### func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddProjectBlack(uid int64, bannedSeconds int32, rest ... interface{}) error

Add users to the project blacklist.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveProjectBlack(uid int64, rest ... interface{}) error

Remove users from the project blacklist.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) IsProjectBlack(uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsProjectBlack(uid int64, rest ... interface{}) (bool, error)

Determine whether the user is in the project blacklist.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real blacklist status will be passed through callback.

### func (client *RTMServerClient) AddProjectBan(uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddProjectBan(uid int64, bannedSeconds int32, rest ... interface{}) error

Ban a user in the project.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) RemoveProjectBan(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveProjectBan(uid int64, rest ... interface{}) error

Unban a user from the project.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) IsProjectBan(uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsProjectBan(uid int64, rest ... interface{}) (bool, error)

Determine whether the user is banned by the project.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real banned status will be passed through callback.

### func (client *RTMServerClient) Kickout(uid int64, rest ... interface{}) error

	func (client *RTMServerClient) Kickout(uid int64, rest ... interface{}) error

Kick users offline.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.
