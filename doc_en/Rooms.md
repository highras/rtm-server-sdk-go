# RTM Server-End Go SDK Rooms API Docs

# Index

[TOC]

### -----------------------[ Room relationship interface ]-----------------------------

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

### func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

	func (client *RTMServerClient) DelRoomMember(roomId int64, uid int64, rest ... interface{}) error

Delete room members.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### -----------------------[ manage interface ]-----------------------------

### func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

	func (client *RTMServerClient) AddRoomBan(roomId int64, uid int64, bannedSeconds int32, rest ... interface{}) error

The user is prohibited from speaking in the designated room.

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

### -----------------------[ Room information interface ]-----------------------------

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
