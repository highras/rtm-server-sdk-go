# RTM Server-End Go SDK Configure API Docs

# Index

[TOC]

## Variables

The Config object of FPNN will affect the behavior of the rtm client. 
Please refer to：[FPNN Go SDK - API Docs](https://github.com/highras/fpnn-sdk-go/blob/master/API.md#variables)

### -----------------------[ Configure interface ]-----------------------------

### func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetConnectTimeOut(timeout time.Duration)

Configure the connection timeout of the RTM Server Client.  
When not configured, the connection timeout parameter of fpnn.Config is used by default.

### func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

	func (client *RTMServerClient) SetQuestTimeOut(timeout time.Duration)

Configure the request timeout of the RTM Server Client.  
When not configured, the request timeout parameter of fpnn.Config is used by default.

### func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

	func (client *RTMServerClient) SetOnConnectedCallback(onConnected func(connId uint64))

Configure the callback function of the connection establishment event.

### func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

	func (client *RTMServerClient) SetOnClosedCallback(onClosed func(connId uint64))

Configure the callback function for disconnection events.

### func (client *RTMServerClient) SetLogger(logger *log.Logger)

	func (client *RTMServerClient) SetLogger(logger *log.Logger)

Configure log routing for RTM Server Client.

### func (client *RTMServerClient) Endpoint() string

	func (client *RTMServerClient) Endpoint() string

Get the RTM Server Client connection/destination address.

### func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

	func (client *RTMServerClient) EnableEncryptor(rest ... interface{}) (err error)

Configure the use of encrypted links.

The acceptable parameters are:

+ `pemKeyPath string`

	The server public key file path. PEM format. Mutually exclusive with the pemKeyData parameter.

+ `pemKeyData []byte`

	The content of the server public key file. PEM format. Mutually exclusive with the pemKeyPath parameter.

+ `reinforce bool`

	True uses 256-bit key encryption, false uses 128-bit key encryption. Default is true

### func (client *RTMServerClient) SetAutoReconnect(autoReconnect bool)

	func (client *RTMServerClient) SetAutoReconnect(autoReconnect bool)

Configure the connection method. True means automatic connection, false requires an explicit call to the Connect() or Dial() method to establish a connection. When not configured, the connection is automatically established by default.

### func (client *RTMServerClient) Connect() bool

	func (client *RTMServerClient) Connect() bool

Explicitly establish a connection. If SetAutoReconnect() is set to false, this method must be called to establish a connection with the server.
Same as `func (client *RTMServerClient) Dial() bool`.

### func (client *RTMServerClient) Dial() bool

	func (client *RTMServerClient) Dial() bool

Explicitly establish a connection. If SetAutoReconnect() is set to false, this method must be called to establish a connection with the server.  
Same as `func (client *RTMServerClient) Connect() bool`.

### func (client *RTMServerClient) IsConnected() bool

	func (client *RTMServerClient) IsConnected() bool

Check whether the link has been **established**.