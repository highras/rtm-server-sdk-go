# RTM Server-End Go SDK API Docs

# Package rtm

	import "github.com/highras/rtm-server-sdk-go/src/rtm"

The rtm package provides the ability to connect and access RTM back-end services with go

# Index

[TOC]

## Constants

	const SDKVersion = "0.9.4"

## Variables

The Config object of FPNN will affect the behavior of the rtm client。  
Please refer to：[FPNN Go SDK - API Docs](https://github.com/highras/fpnn-sdk-go/blob/master/API.md#variables)


## RTMServerClient

	type RTMServerClient struct {
		//-- same hidden fields
	}

RTM Server Client。


### func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

	func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

Create RTM Server Client。

Please get all parameters through RTM Console

## Configure

Configure RTMServerClient behavior。

Please refer to：[Config RTMServerClient](Config.md)


## Token Functions

Log in to Token related interface。

Please refer to：[Token Interface](Token.md)


## Chat Functions

Chat interface。

Please refer to：[Chat interface](Chat.md)


## Message Functions

Messaging interface。

Please refer to：[Messaging interface](Messages.md)


## Files Functions

File sending interface。

Please refer to：[File sending interface](Files.md)


## Friend Functions

Friendship interface。

Please refer to：[Friendship interface](Friends.md)


## Group Functions

Group Relationship Interface。

Please refer to：[Group Relationship Interface](Groups.md)


## Room Functions

Room relationship interface。

Please refer to：[Room relationship interface](Rooms.md)


## User Functions

User information interface。

Please refer to：[User information interface](Users.md)


## Data Functions

User data interface。

Please refer to：[User data interface](Data.md)


## Device Functions

Device information interface。

Please refer to：[Device information interface](Devices.md)


## Monitor Functions

Chat & Message Monitoring Interface。

Please refer to：[Chat & Message Monitoring Interface](Listening.md)


## Black Functions

Block user related interfaces。

Please refer to: [Block user related interfaces](Blacklist.md)

## RTC Functions

Real-time Communication related interface。

Please refer to: [Real-time Communication related interface](RTC.md)
