# RTM Server-End Go SDK RealTimeRTC API Docs

# Index

[TOC]

### -----------------------[Real-time Communication room management interface]-----------------------------

### func (client *RTMServerClient) InviteUserIntoRTCRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) InviteUserIntoRTCRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

Invite users to join the RTC room.

Parameter Description:

+ `fromUid int64`:

	The user who initiated the invitation (must be in the RTC room to initiate the invitation command)

+ `uids []int64`:

	List of invited users

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) PullUserIntoRTCRoom(roomId int64, uids []int64, roomType int32, rest ...interface{}) error

	func (client *RTMServerClient) PullUserIntoRTCRoom(roomId int64, uids []int64,roomType int32, rest ...interface{}) error

Force the user to enter the RTC room (the room will be created automatically if the room does not exist).

The acceptable parameters are:

+ `roomType int32`

	room type 1 voice, 2 video

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) CloseRTCRoom(roomId int64, rest ...interface{}) error

	func (client *RTMServerClient) CloseRTCRoom(roomId int64, rest ...interface{}) error

Close the RTC room

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) KickoutFromRTCRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

	func (client *RTMServerClient) KickoutFromRTCRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

Kicked out from the RTC room.

Parameter Description:

+ `uid int64`:

	Kicked user

+ `fromUid int64`:

	The user who initiated the kick command

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) SetRTCRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

	func (client *RTMServerClient) SetRTCRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

Set the default microphone state of the RTC room (whether the microphone is turned on by default when the user just enters the room).

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### -----------------------[Real-time Communication room information query interface]-----------------------------

### func (client *RTMServerClient) GetRTCRoomList(rest ...interface{}) ([]int64, error))

	func (client *RTMServerClient) GetRTCRoomList(rest ...interface{}) (rids []int64, error)

Get the current RTC room id list of the project

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(rids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the RTC room id list and error message are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real RTC room id list will be passed through callback.

### func (client *RTMServerClient) GetRTCRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, int64, error)

	func (client *RTMServerClient) GetRTCRoomMembers(roomId int64, rest ...interface{}) (uids []int64, managers []int64, owner int64, error)

Get the current user id list and administrator id list and owner id of the RTC room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (uids []int64, managers []int64 errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request. Return the current RTC room user id list, administrator id list and error information
If the **callback** parameter **exists**, it is an **asynchronous** request. The real data is returned through callback

### func (client *RTMServerClient) GetRTCRoomMemberCount(roomId int64, rest ...interface{}) (int32, error)

	func (client *RTMServerClient) GetRTCRoomMemberCount(roomId int64, rest ...interface{}) (count int32, error)

Get the number of people in the current RTC room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (count int32, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning the current number of people in the room and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request, and the real number of people in the room will be returned through callback.

### func (client *RTMServerClient) AdminCommand(roomId int64, uids []int64, command int32, rest ...interface{}) error

	func (client *RTMServerClient) AdminCommand(roomId int64, uids []int64, command int32, rest ...interface{}) error

Room manager operation

The acceptable parameters are:

+ `uids []int64` 
	
	List of user ids operated

+ `command int32`

	Operation type: 0 grant administrator rights, 1 deprive administrator rights, 2 prohibit sending audio data, 3 allow sending audio data, 4 prohibit sending video data, 5 allow sending video data, 6 turn off others’ microphones, 7 turn off others’ cameras

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.