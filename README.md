# RTM Server-End Go SDK

[TOC]

## Depends & Install & Update

### Depends

	go get github.com/highras/fpnn-sdk-go/src/fpnn

### Install

	go get github.com/highras/rtm-server-sdk-go/src/rtm

### Update

	go get -u github.com/highras/rtm-server-sdk-go/src/rtm

### Use

	import "github.com/highras/rtm-server-sdk-go/src/rtm"

### Notice

* Before using the SDK, please make sure the server time is correct, RTM-Server will check whether the signature time has expired
	
## Usage

### Create

	client := rtm.NewRTMServerClient(pid int32, secretKey string, endpoint string)

Please get your project params from RTM Console.

### Configure (Optional)

* Basic configs

		client.SetAutoReconnect(autoReconnect bool) 
		client.SetConnectTimeOut(timeout time.Duration)
		client.SetQuestTimeOut(timeout time.Duration)
		client.SetLogger(logger *log.Logger)

	**Note**  
	**autoReconnect** means establishing the connection in implicit or explicit. NOT keep the connection.

* Set message monitor

		client.SetMonitor(monitor RTMServerMonitor)

* Set connection events' callbacks

		client.SetOnConnectedCallback(onConnected func(connId uint64))
		client.SetOnClosedCallback(onClosed func(connId uint64))

* Config encrypted connection
	
		client.EnableEncryptor(pemKeyPath string)
		client.EnableEncryptor(pemKeyData []byte)

	RTM Server-End Go SDK using **ECC**/**ECDH** to exchange the secret key, and using **AES-128** or **AES-256** in **CFB** mode to encrypt the whole session in **stream** way.


### Connect/Dial (Optional)

	client.Connect()
	client.Dial()

Call one of these methods to do an explicit connecting action.  
If client.SetAutoReconnect(false) is called, one of these explicit connecting methods MUST be called; otherwise, these methods are optional.

### Send messages

* Send P2P Message

		mtime, err := client.SendMessage(fromUid int64, toUid int64, mtype int8, message string)
		mtime, err := client.SendMessage(fromUid int64, toUid int64, mtype int8, message string, timeout time.Duration)

		_, err := client.SendMessage(fromUid int64, toUid int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string))
		_, err := client.SendMessage(fromUid int64, toUid int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string), timeout time.Duration)

* Send Multi-Receivers P2P Message

		mtime, err := client.SendMessages(fromUid int64, toUids []int64, mtype int8, message string)
		mtime, err := client.SendMessages(fromUid int64, toUids []int64, mtype int8, message string, timeout time.Duration)

		_, err := client.SendMessages(fromUid int64, toUids []int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string))
		_, err := client.SendMessages(fromUid int64, toUids []int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string), timeout time.Duration)

* Send Group Message
	
		mtime, err := client.SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string)
		mtime, err := client.SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, timeout time.Duration)

		_, err := client.SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string))
		_, err := client.SendGroupMessage(fromUid int64, groupId int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string), timeout time.Duration)

* Send Room Message

		mtime, err := client.SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string)
		mtime, err := client.SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, timeout time.Duration)

		_, err := client.SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string))
		_, err := client.SendRoomMessage(fromUid int64, roomId int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string), timeout time.Duration)

* Send Boradcast Message

		mtime, err := client.SendBoradcastMessage(fromUid int64, mtype int8, message string)
		mtime, err := client.SendBoradcastMessage(fromUid int64, mtype int8, message string, timeout time.Duration)

		_, err := client.SendBoradcastMessage(fromUid int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string))
		_, err := client.SendBoradcastMessage(fromUid int64, mtype int8, message string, callback func (mtime int64, errorCode int, errInfo string), timeout time.Duration)


### SDK Version

	fmt.Println("RTM Server-End Go SDK Version:", rtm.SDKVersion)

## API docs

Please refer: [API docs](doc/API.md)


## Directory structure

* **<rtm-server-sdk-go>/src**

	Codes of SDK.

* **<rtm-server-sdk-go>/example**

	Examples codes for using this SDK.

* **<rtm-server-sdk-go>/doc**

	API documents in markdown format.