# RTM Server-End Go SDK Blacklist API Docs

# Index

[TOC]

### ---------------------------[Block user operation related interfaces]---------------------------

### func (client *RTMServerClient) AddBlacks(uid int64, blacks []int64, rest ...interface{}) error
    func (client *RTMServerClient) AddBlacks(uid int64, blacks []int64, rest ...interface{}) error

Block users, each **block 100 people** at most. After blocking, the other party can't send messages to themselves, and they can send chat messages to each other, but both parties can get session and historical information normally。

Acceptable parameters：

* `timeout time.Duration`
  
        Request timed out.
        When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be used.
        If the RTM Server Client instance is not configured, the timeout configuration of fpnn.Config will be used.

* `callback func(errorCode int, errorInfo string)`
  
        Asynchronous callback function。

If the **callback** parameter ** does not exist**, it is a **synchronous** request。
If the **callback** parameter **exists**, it is an **asynchronous** request。  

### func (client *RTMServerClient) DelBlacks(uid int64, blacks []int64, rest ...interface{}) error
    func (client *RTMServerClient) DelBlacks(uid int64, blacks []int64, rest ...interface{}) error

Unblocking, each **can remove up to 100 people**。

Acceptable parameters：

* `timeout time.Duration`
  
        Request timed out.
        When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be used.
        If the RTM Server Client instance is not configured, the timeout configuration of fpnn.Config will be used.

* `callback func(errorCode int, errorInfo string)`
  
        Asynchronous callback function。

If the **callback** parameter ** does not exist**, it is a **synchronous** request。
If the **callback** parameter **exists**, it is an **asynchronous** request。

### func (client * RTMServerClient) GetBlacks(uid int64, rest ...interface{}) ([]int64, error)
    func (client * RTMServerClient) GetBlacks(uid int64, rest ...interface{}) ([]int64, error)

Get the list of users blocked by uid。

Acceptable parameters：

* `timeout time.Duration`
  
        Request timed out.
        When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be used.
        If the RTM Server Client instance is not configured, the timeout configuration of fpnn.Config will be used.

* `callback func(buids []int64, errorCode int, errorInfo string)`
  
        Asynchronous callback function。

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and the blacklisted list and error information are returned.
If the **callback** parameter **exists**, it is an **asynchronous** request, returning nil and error information, and the real blacklisted list will be returned through callback.

### func (client *RTMServerClient) IsBlacks(uid int64, buids []int64, rest ...interface{}) ([]int64, error)
    func (client *RTMServerClient) IsBlacks(uid int64, buids []int64, rest ...interface{}) ([]int64, error)

Get the black relationship, check whether the uid is blacked by buids, and use it when sending multi-person chat messages. Get a maximum of 100 people** each time.

Acceptable parameters：

* `timeout time.Duration`
  
        Request timed out.
        When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be used.
        If the RTM Server Client instance is not configured, the timeout configuration of fpnn.Config will be used.

* `callback func(buids []int64, errorCode int, errorInfo string)`
  
        Asynchronous callback function。

If the **callback** parameter ** does not exist**, it is a **synchronization** request, returning a list of users who have blocked uid (the buids returned cannot receive chat messages sent by uid) and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request, returning nil and error information, and the list of users who have blocked the uid (the buids returned cannot receive the chat messages sent by the uid) Return via callback.

### func (client *RTMServerClient) IsBlack(uid, buid int64, rest ...interface{}) (bool, error)
    func (client *RTMServerClient) IsBlack(uid, buid int64, rest ...interface{}) (bool, error)

Get the black relationship, check whether the uid is blacked by the buid, and use it when sending a single chat message.

Acceptable parameters：

* `timeout time.Duration`
  
        Request timed out.
        When the timeout parameter is missing or the timeout parameter is 0, the configuration of the RTM Server Client instance will be used.
        If the RTM Server Client instance is not configured, the timeout configuration of fpnn.Config will be used.

* `callback func(ok bool, errorCode int, errorInfo string)`
  
        Asynchronous callback function.

If the **callback** parameter ** does not exist**, it is a **synchronization** request, and bool is returned (if it returns true, the chat message sent by uid cannot be received by buid) and error information.
If the **callback** parameter **exists**, it is an **asynchronous** request, returns false and error information, and blacks the relationship (if the bool returned by the callback is true, the buid cannot receive the chat message sent by the uid ) Will return via callback.
