# RTM Server-End Go SDK Listening API Docs

# Index

[TOC]

## type IRTMServerMonitor

	type IRTMServerMonitor interface {
		P2PMessage(messageInfo *RTMMessage)
		GroupMessage(messageInfo *RTMMessage)
		RoomMessage(messageInfo *RTMMessage)

		P2PChat(messageInfo *RTMMessage)
		GroupChat(messageInfo *RTMMessage)
		RoomChat(messageInfo *RTMMessage)

		P2PCmd(messageInfo *RTMMessage)
		GroupCmd(messageInfo *RTMMessage)
		RoomCmd(messageInfo *RTMMessage)

		Event(pid int32, event string, uid int64, time int32, endpoint string, data string)
	}

Message monitoring interface.

Please configure the settings through the RTM Console, and after the connection is established, call `func AddListen(...)` or `func SetListen(...)` to set the code.


### func (client *RTMServerClient) SetServerPushMonitor(monitor IRTMServerMonitor)

	func (client *RTMServerClient) SetServerPushMonitor(monitor IRTMServerMonitor)

Configure the message monitoring interface.
Specific reference: [IRTMServerMonitor](#type-IRTMServerMonitor)


### func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

**Increment** Add monitoring settings.

Required parameters:

+ `groupIds []int64`

	Increase the monitoring group.

+ `roomIds []int64`

	Increase the monitoring room.

+ `uids []int64`

	Increase monitoring P2P users.

+ `events []string`

	The event that needs to be monitored.
	For the list of events that can be monitored, please refer to the RTM service documentation.


The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

**Increment** Cancel monitoring settings.

Required parameters:

+ `groupIds []int64`

	Cancel the monitoring group.

+ `roomIds []int64`

	Cancel the monitoring room.

+ `uids []int64`

	Cancel the P2P users

+ `events []string`

	The event that needs to be monitored.
	For the list of events that can be monitored, please refer to the RTM service documentation.


The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.

### func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error

	func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error 

Set the monitoring state. This interface will **overwrite** the previous settings.

Required parameters:

+ `groupIds []int64`

	Set the monitoring group.

+ `roomIds []int64`

	Set the monitoring room.

+ `uids []int64`

	Set the monitoring P2P users.

+ `events []string`

	The event that needs to be monitored.
	For the list of events that can be monitored, please refer to the RTM service documentation.


The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.


### func (client *RTMServerClient) SetListenStatus(allGroups bool, allRrooms bool, allP2P bool, allEvents bool, rest ... interface{}) error

	func (client *RTMServerClient) SetListenStatus(allGroups bool, allRrooms bool, allP2P bool, allEvents bool, rest ... interface{}) error 

Set the monitoring state. This interface will **overwrite** the previous settings.

Required parameters:

+ `allGroups bool`

	Set whether to monitor all groups.

+ `allRrooms bool`

	Set whether to monitor all rooms.

+ `allP2P bool`

	Set whether to listen to all P2P messages.

+ `allEvents bool`

	Set whether to listen to all events.
	For the list of events that can be monitored, please refer to the RTM service documentation.


The acceptable parameters are:

+ `timeout time.Duration`

	Request timed out.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is a **synchronization** request.
If the **callback** parameter **exists**, it is an **asynchronous** request.
