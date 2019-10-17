# RTM Server-End Go SDK Files API Docs

# Index

[TOC]

### -----------------------[ 文件传送接口 ]-----------------------------

### func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFile(fromUid int64, toUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送 P2P 文件。

可接受的参数为：

+ `mtype int8`

	消息类型。默认 50。

+ `extension string`

	文件的扩展名。  
	如果缺失，会自动从 filename 参数中提取。

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendFiles(fromUid int64, toUids []int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送多人 P2P 文件。

可接受的参数为：

+ `mtype int8`

	消息类型。默认 50。

+ `extension string`

	文件的扩展名。  
	如果缺失，会自动从 filename 参数中提取。

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendGroupFile(fromUid int64, groupId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送群组文件。

可接受的参数为：

+ `mtype int8`

	消息类型。默认 50。

+ `extension string`

	文件的扩展名。  
	如果缺失，会自动从 filename 参数中提取。

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendRoomFile(fromUid int64, roomId int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送房间文件。

可接受的参数为：

+ `mtype int8`

	消息类型。默认 50。

+ `extension string`

	文件的扩展名。  
	如果缺失，会自动从 filename 参数中提取。

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。

### func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

	func (client *RTMServerClient) SendBroadcastFile(fromUid int64, fileContent []byte, filename string, rest ... interface{}) (int64, error)

发送广播文件。

可接受的参数为：

+ `mtype int8`

	消息类型。默认 50。

+ `extension string`

	文件的扩展名。  
	如果缺失，会自动从 filename 参数中提取。

+ `attrs string`

	消息的属性信息。默认为空字符串。

+ `timeout time.Duration`

	发送超时。  
	缺少 timeout 参数，或 timeout 参数为 0 时，将采用 RTM Server Client 实例的配置。  
	若 RTM Server Client 实例未配置，将采用 fpnn.Config 的相应配置。

+ `callback func(mtime int64, errorCode int, errInfo string)`

	异步回调函数。  

如果 **callback** 参数**不存在**，则为**同步**发送，返回 mtime 及 error 信息。  
如果 **callback** 参数**存在**，则为**异步**发送，返回 0 及 error 信息。真实的 mtime，将通过 callback 传递。
