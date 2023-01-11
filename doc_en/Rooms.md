# RTM Server-End Go SDK Rooms API Docs

# Index

[TOC]

### -----------------------[ Room Relationship Interface]-----------------------------

### func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomMember(roomId int64, uid int64, rest ... interface{}) error

Add room members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) AddRoomMembers(roomId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomMembers(roomId int64, uids []int64, rest ... interface{}) error

Add room members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

Delete room members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(successedUids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DelRoomMembers(roomId int64, uids []int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMembers(roomId int64, uids []int64, rest ... interface{}) error

Delete room members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(successedUids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) AddUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

	func (client *RTMServerClient) AddUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

Add a user in some rooms.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) DeleteUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

	func (client *RTMServerClient) DeleteUserRooms(roomIds []int64, uid int64, rest ...interface{}) error

Delete a user from some rooms.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### -----------------------[Management Interface]-------------------- ---------

### func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

The user is prohibited from speaking in the designated room.

Parameter Description:

+ `roomId int64`:

	When roomId <= 0, all rooms are muted

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) RemoveRoomBan(roomId int64, uid int64, rest ... interface{}) error

Remove the mute in the room specified by the user.

Parameter Description:

+ `roomId int64`:

	When roomId <= 0, all rooms will be muted

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error)

	func (client *RTMServerClient) IsBanOfRoom(roomId int64, uid int64, rest ... interface{}) (bool, error)

Determine whether the user is muted in the designated room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(ok bool, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool and error messages are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and false and error messages are returned. The real mute status will be passed through callback.

### -----------------------[Room Information Interface]------------------- ----------

### func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

	func (client *RTMServerClient) SetRoomInfo(roomId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error

Set the room public information and private information.

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

### func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ... interface{}) (string, string, error)

	func (client *RTMServerClient) GetRoomInfo(roomId int64, rest ... interface{}) (string, string, error)

Get public and private information of the room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (publicInfo string, privateInfo string, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning the room public information, room private information, and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns "", "", error information. The real room public information and private information will be passed through callback.

### func (client *RTMServerClient) GetRoomMembers(roomId int64, rest ...interface{}) ([]int64, error)

	func (client *RTMServerClient) GetRoomMembers(roomId int64, rest ...interface{}) ([]int64, error)

Get all members of the room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (uids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the room member and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, and uids []int64, error information is returned. The real room members will be passed through callback.

### func (client *RTMServerClient) GetRoomCount(roomId []int64, rest ...interface{}) (map[int64]int32, error)

	func (client *RTMServerClient) GetRoomCount(roomId []int64, rest ...interface{}) (map[int64]int32, error)

Get the number of users in the room

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(map[int64]int32, int, string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning the number of users in the room and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request and return count map[int64]int32, error information. The number of users in the real room will be passed through callback.