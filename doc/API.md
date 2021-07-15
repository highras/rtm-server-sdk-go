# RTM Server-End Go SDK API Docs

# Package rtm

	import "github.com/highras/rtm-server-sdk-go/src/rtm"

rtm 包提供go连接和访问 RTM 后端服务的能力。

# Index

[TOC]

## Constants

	const SDKVersion = "0.9.3"

## Variables

FPNN 的 Config 对象，将会影响 rtm 客户端的行为。  
具体请参见：[FPNN Go SDK - API Docs](https://github.com/highras/fpnn-sdk-go/blob/master/API.md#variables)


## RTMServerClient

	type RTMServerClient struct {
		//-- same hidden fields
	}

RTM Server 客户端。


### func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

	func NewRTMServerClient(pid int32, secretKey string, endpoint string) *RTMServerClient

创建 RTM Server 客户端。

所有参数请通过 RTM Console 获取。

## Configure

配置 RTMServerClient 行为。

具体请参见：[配置 RTMServerClient](Config.md)


## Token Functions

登陆 Token 相关接口。

具体请参见：[Token 接口](Token.md)


## Chat Functions

聊天接口。

具体请参见：[聊天接口](Chat.md)


## Message Functions

消息传送接口。

具体请参见：[消息传送接口](Messages.md)


## Files Functions

文件传送接口。

具体请参见：[文件传送接口](Files.md)


## Friend Functions

好友关系接口。

具体请参见：[好友关系接口](Friends.md)


## Group Functions

群组关系接口。

具体请参见：[群组关系接口](Groups.md)


## Room Functions

房间关系接口。

具体请参见：[房间关系接口](Rooms.md)


## User Functions

用户信息接口。

具体请参见：[用户信息接口](Users.md)


## Data Functions

用户数据接口。

具体请参见：[用户数据接口](Data.md)


## Device Functions

设备信息接口。

具体请参见：[设备信息接口](Devices.md)


## Monitor Functions

聊天 & 消息监控接口。

具体请参见：[聊天 & 消息监控接口](Listening.md)


## Black Functions

拉黑用户相关接口。

具体请参见: [拉黑用户相关接口](Blacklist.md)

## RTC Functions

实时音视频相关接口。

具体请参见: [实时音视频相关接口](RTC.md)
