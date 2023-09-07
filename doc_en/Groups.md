# RTM Server-End Go SDK Groups API Docs

# Index

[TOC]

### -----------------------[Group Relationship Interface]-----------------------------

### func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

Add group members. **Up to 100 people are added each time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error

Delete group members. **Up to 100 people are deleted each time.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error

	func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error

Delete the group.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error)

Get group members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(uids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the group member list and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real list of group members will be passed through callback.

### func (client *RTMServerClient) GetGroupCount(groupId int64, rest ... interface{}) (int32, error)

	func (client *RTMServerClient) GetGroupCount(groupId int64, rest ... interface{}) (int32, error)

Get group members count.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(count int32, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the group member count and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns -1 and error information. The real count of group members will be passed through callback.

### func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error)

Determine the group relationship.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real group relationship will be passed through callback.

### func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error)

	func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error)

Get the group that the user has joined.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(groupIds []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the group list and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real group list will be passed through callback.

### -----------------------[Management Interface]-------------------- ---------

### func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

Forbid users to speak in the specified group.

Parameter Description:

+ `groupId int64`:

	When groupId <= 0, all groups will be muted

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error

Unblock the user-specified group.

Parameter Description:

+ `groupId int64`:

	When groupId <= 0, all groups will be unblocked

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error)

Determine whether the user is banned in the specified group.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real mute status will be passed through callback.

### -----------------------[Group Information Interface]------------------ -----------

### func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

Set group public information and private information.

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

### func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ... interface{}) (string, string, error)

Get public and private information of the group.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning group public information, group private information, and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "", "", error information. The true public and private information of the group will be passed through callbacks.

### func (client *RTMServerClient) ClearProjectGroup(rest ... interface{}) error

	func (client *RTMServerClient) ClearProjectGroup(rest ... interface{}) error

Clear all the group info in project, note: group message is not included, use ClearProjectMessage if necessary.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.
