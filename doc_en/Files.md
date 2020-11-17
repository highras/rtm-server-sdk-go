# RTM Server-End Go SDK Files API Docs

# Index

[TOC]

### -----------------------[ File sending interface ]-----------------------------

### func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

Send P2P files.

The optional parameters are:

+ `filename string`

	File name, preferably a file name

+ `attr string`

	Custom attribute **must be a json string if not empty**

+ `audioInfo *RTMAudioFileInfo`

	rtm Voice message parameters, see structure content：[RTMAudioFileInfo](Messages.md#type-RTMAudioFileInfo)

The acceptable parameters are:

+ `mtype int8`

	Message type. The default is 50.

+ `extension string`

	The extension of the file.
	If it is missing, it will be automatically extracted from the filename parameter.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

Send multi-person P2P files.

The optional parameters are:

+ `filename string`

	File name, preferably a file name

+ `attr string`

	Custom attribute **must be a json string if not empty**

+ `audioInfo *RTMAudioFileInfo`

	rtm Voice message parameters, see structure content：[RTMAudioFileInfo](Messages.md#type-RTMAudioFileInfo)

The acceptable parameters are:

+ `mtype int8`

	Message type. The default is 50.

+ `extension string`

	The extension of the file.
	If it is missing, it will be automatically extracted from the filename parameter.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

Send group files.

The optional parameters are:

+ `filename string`

	File name, preferably a file name

+ `attr string`

	Custom attribute **must be a json string if not empty**

+ `audioInfo *RTMAudioFileInfo`

	rtm Voice message parameters, see structure content：[RTMAudioFileInfo](Messages.md#type-RTMAudioFileInfo)

The acceptable parameters are:

+ `mtype int8`

	Message type. The default is 50.

+ `extension string`

	The extension of the file.
	If it is missing, it will be automatically extracted from the filename parameter.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

Send room files.

The optional parameters are:

+ `filename string`

	File name, preferably a file name

+ `attr string`

	Custom attribute **must be a json string if not empty**

+ `audioInfo *RTMAudioFileInfo`

	rtm Voice message parameters, see structure content：[RTMAudioFileInfo](Messages.md#type-RTMAudioFileInfo)

The acceptable parameters are:：

+ `mtype int8`

	Message type. The default is 50.

+ `extension string`

	The extension of the file.
	If it is missing, it will be automatically extracted from the filename parameter.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function.  

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exists**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.

### func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, attr string, audioInfo *RTMAudioFileInfo, rest ... interface{}) (int64, error)

Send broadcast files.

The optional parameters are:

+ `filename string`

	File name, preferably a file name

+ `attr string`

	Custom attribute **must be a json string if not empty**

+ `audioInfo *RTMAudioFileInfo`

	rtm Voice message parameters, see structure content：[RTMAudioFileInfo](Messages.md#type-RTMAudioFileInfo)

The acceptable parameters are:

+ `mtype int8`

	Message type. The default is 50.

+ `extension string`

	The extension of the file.
	If it is missing, it will be automatically extracted from the filename parameter.

+ `timeout time.Duration`

	Send timeout.
	When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be adopted.
	If the RTM Server Client instance is not configured, the corresponding configuration of fpnn.Config will be adopted.

+ `callback func(mtime int64, errorCode int, errInfo string)`

	Asynchronous callback function. 

If the **callback** parameter ** does not exist**, it is sent as **synchronized**, and mtime and error messages are returned.
If the **callback** parameter **exist**, it is **asynchronous** sending, and 0 and error information are returned. The real mtime will be passed through callback.
