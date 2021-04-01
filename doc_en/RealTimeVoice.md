# RTM Server-End Go SDK RealTimeVoice API Docs

# Index

[TOC]

### -----------------------[Real-time voice room management interface]-----------------------------

### func (client *RTMServerClient) InviteUserIntoVoiceRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) InviteUserIntoVoiceRoom(roomId int64, fromUid int64, uids []int64, rest ...interface{}) error

Invite users to join the voice room.

Parameter Description:

+ `fromUid int64`:

	The user who initiated the invitation (must be in the voice room to initiate the invitation command)

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

### func (client *RTMServerClient) PullUserIntoVoiceRoom(roomId int64, uids []int64, rest ...interface{}) error

	func (client *RTMServerClient) PullUserIntoVoiceRoom(roomId int64, uids []int64, rest ...interface{}) error

Force the user to enter the voice room (the room will be created automatically if the room does not exist).

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) CloseVoiceRoom(roomId int64, rest ...interface{}) error

	func (client *RTMServerClient) CloseVoiceRoom(roomId int64, rest ...interface{}) error

Close the voice room

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) KickoutFromVoiceRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

	func (client *RTMServerClient) KickoutFromVoiceRoom(roomId int64, uid int64, fromUid int64, rest ...interface{}) error

Kicked out from the voice room.

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

### func (client *RTMServerClient) SetVoiceRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

	func (client *RTMServerClient) SetVoiceRoomMicStatus(roomId int64, status bool, rest ...interface{}) error

Set the default microphone state of the voice room (whether the microphone is turned on by default when the user just enters the room).

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### -----------------------[Real-time voice room information query interface]-----------------------------

### func (client *RTMServerClient) GetVoiceRoomList(rest ...interface{}) ([]int64, error))

	func (client *RTMServerClient) GetVoiceRoomList(rest ...interface{}) (rids []int64, error)

Get the current voice room id list of the project

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(rids []int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the voice room id list and error message are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request and returns nil and error information. The real voice room id list will be passed through callback.

### func (client *RTMServerClient) GetVoiceRoomMembers(roomId int64, rest ...interface{}) ([]int64, []int64, error)

	func (client *RTMServerClient) GetVoiceRoomMembers(roomId int64, rest ...interface{}) (uids []int64, managers []int64, error)

Get the current user id list and administrator id list of the voice room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (uids []int64, managers []int64 errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request. Return the current voice room user id list, administrator id list and error information
If the **callback** parameter **exists**, it is an **asynchronous** request. The real data is returned through callback

### func (client *RTMServerClient) GetVoiceRoomMemberCount(roomId int64, rest ...interface{}) (int32, error)

	func (client *RTMServerClient) GetVoiceRoomMemberCount(roomId int64, rest ...interface{}) (count int32, error)

Get the number of people in the current voice room.

The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func (count int32, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning the current number of people in the room and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request, and the real number of people in the room will be returned through callback.