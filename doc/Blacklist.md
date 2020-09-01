# RTM Server-End Go SDK Blacklist API Docs

# Index

[TOC]

### ---------------------------[拉黑用户操作相关接口]---------------------------

### func (client *RTMServerClient) AddBlacks(uid int64, blacks []int64, rest ...interface{}) error
    func (client *RTMServerClient) AddBlacks(uid int64, blacks []int64, rest ...interface{}) error

拉黑用户，每次**最多拉黑100人**, 拉黑后对方不能给自己发消息，自己可以给对方发送聊天消息，
双方能正常获取session及历史信息。

可接受的参数：

* `timeout time.Duration`
  
        请求超时。
        缺少timeout参数，或timeout参数为0时，将采用RTM Server Client实例的配置。
        若RTM Server Client实例未配置，将采用fpnn.Config的超时配置。

* `callback func(errorCode int, errorInfo string)`
  
        异步回掉函数。

如果 **callback** 参数**不存在**，则为**同步**请求。
如果 **callback** 参数**存在**，则为**异步**请求。  

### func (client *RTMServerClient) DelBlacks(uid int64, blacks []int64, rest ...interface{}) error
    func (client *RTMServerClient) DelBlacks(uid int64, blacks []int64, rest ...interface{}) error

解除拉黑，每次**最多解除拉黑100人**。

可接受的参数：

* `timeout time.Duration`
  
        请求超时。
        缺少timeout参数，或timeout参数为0时，将采用RTM Server Client实例的配置。
        若RTM Server Client实例未配置，将采用fpnn.Config的超时配置。

* `callback func(errorCode int, errorInfo string)`
  
        异步回掉函数。

如果 **callback** 参数**不存在**，则为**同步**请求。
如果 **callback** 参数**存在**，则为**异步**请求。

### func (client * RTMServerClient) GetBlacks(uid int64, rest ...interface{}) ([]int64, error)
    func (client * RTMServerClient) GetBlacks(uid int64, rest ...interface{}) ([]int64, error)

获取被uid拉黑的用户列表。

可接受的参数：

* `timeout time.Duration`
  
        请求超时。
        缺少timeout参数，或timeout参数为0时，将采用RTM Server Client实例的配置。
        若RTM Server Client实例未配置，将采用fpnn.Config的超时配置。

* `callback func(buids []int64, errorCode int, errorInfo string)`
  
        异步回掉函数。

如果 **callback** 参数**不存在**，则为**同步**请求，返回拉黑列表以及error信息。
如果 **callback** 参数**存在**，则为**异步**请求，返回nil以及error信息，真实的拉黑列表将通过callback返回。

### func (client *RTMServerClient) IsBlacks(uid int64, buids []int64, rest ...interface{}) ([]int64, error)
    func (client *RTMServerClient) IsBlacks(uid int64, buids []int64, rest ...interface{}) ([]int64, error)

获取拉黑关系，检查uid是否被buids拉黑，在发送多人聊天消息的时候使用。每次**最多获取100人**的拉黑关系

可接受的参数：

* `timeout time.Duration`
  
        请求超时。
        缺少timeout参数，或timeout参数为0时，将采用RTM Server Client实例的配置。
        若RTM Server Client实例未配置，将采用fpnn.Config的超时配置。

* `callback func(buids []int64, errorCode int, errorInfo string)`
  
        异步回掉函数。

如果 **callback** 参数**不存在**，则为**同步**请求，返回拉黑uid的用户列表(返回的这些buids收不到uid发送的聊天消息)以及error信息。
如果 **callback** 参数**存在**，则为**异步**请求，返回nil以及error信息，真实拉黑uid的用户列表(返回的这些buids收不到uid发送的聊天消息)将通过callback返回。

### func (client *RTMServerClient) IsBlack(uid, buid int64, rest ...interface{}) (bool, error)
    func (client *RTMServerClient) IsBlack(uid, buid int64, rest ...interface{}) (bool, error)

获取拉黑关系，检查uid是否被buid拉黑，在发送单人聊天消息的时候使用。

可接受的参数：

* `timeout time.Duration`
  
        请求超时。
        缺少timeout参数，或timeout参数为0时，将采用RTM Server Client实例的配置。
        若RTM Server Client实例未配置，将采用fpnn.Config的超时配置。

* `callback func(ok bool, errorCode int, errorInfo string)`
  
        异步回掉函数。

如果 **callback** 参数**不存在**，则为**同步**请求，返回bool(如果返回true，则buid收不到uid发送的聊天消息)以及error信息。
如果 **callback** 参数**存在**，则为**异步**请求，返回false以及error信息，拉黑关系(如果callback返回的bool为true，则buid收不到uid发送的聊天消息)将通过callback返回。
