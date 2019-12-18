# RTM Server-End Go SDK Configure API Docs

# Index

[TOC]

## Variables

FPNN 的 Config 对象，将会影响 rtm 客户端的行为。  
具体请参见：[FPNN Go SDK - API Docs](https://github.com/highras/fpnn-sdk-go/blob/master/API.md#variables)

### -----------------------[ 配置接口 ]-----------------------------

### func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

配置 RTM Server Client 的连接超时。  
未配置时，默认采用 fpnn.Config 的连接超时参数。

### func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

配置 RTM Server Client 的请求超时。  
未配置时，默认采用 fpnn.Config 的请求超时参数。

### func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

	func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

配置连接建立事件的回调函数。

### func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

	func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

配置连接断开事件的回调函数。

### func (client *RTMServerClient) SetLogger(logger *log.Logger)

	func (client *RTMServerClient) SetLogger(logger *log.Logger)

配置 RTM Server Client 的日志路由。

### func (client *RTMServerClient) Endpoint() string

	func (client *RTMServerClient) Endpoint() string

获取 RTM Server Client 连接/目标地址。

### func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

	func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

配置使用加密链接。

可接受的参数为：

+ `pemKeyPath string`

	服务器公钥文件路径。PEM 格式。与 pemKeyData 参数互斥。

+ `pemKeyData []byte`

	服务器公钥文件内容。PEM 格式。与 pemKeyPath 参数互斥。

+ `reinforce bool`

	true 采用 256 位密钥加密，false 采用 128 位密钥加密。  
	默认为 true

### func (client *RTMServerClient) SetAutoReconnect(autoReconnect bool)

	func (client *RTMServerClient) SetAutoReconnect(autoReconnect bool)

配置连接方式。true 为自动连接，false 需要显式调用 Connect() 或者 Dial() 方法建立连接。 
未配置时，默认自动建立连接。

### func (client *RTMServerClient) Connect() bool

	func (client *RTMServerClient) Connect() bool

显式建立连接。如果 SetAutoReconnect() 设置为 false，必须调用该方法才能与服务器建立连接。  
与 `func (client *RTMServerClient) Dial() bool` 相同。

### func (client *RTMServerClient) Dial() bool

	func (client *RTMServerClient) Dial() bool

显式建立连接。如果 SetAutoReconnect() 设置为 false，必须调用该方法才能与服务器建立连接。  
与 `func (client *RTMServerClient) Connect() bool` 相同。

### func (client *RTMServerClient) IsConnected() bool

	func (client *RTMServerClient) IsConnected() bool

查询链接是否**已经**建立。