package rtm

import (
	"time"
)

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddListen() function.")
		}
	}

	quest := client.genServerQuest("addlisten")

	if groupIds != nil && len(groupIds) > 0 {
		quest.Param("gids", groupIds)	
	}
	if roomIds != nil && len(roomIds) > 0 {
		quest.Param("rids", roomIds)	
	}
	if uids != nil && len(uids) > 0 {
		quest.Param("uids", uids)	
	}
	if events != nil && len(events) > 0 {
		quest.Param("events", events)	
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) RemoveListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveListen() function.")
		}
	}

	quest := client.genServerQuest("removelisten")

	if groupIds != nil && len(groupIds) > 0 {
		quest.Param("gids", groupIds)	
	}
	if roomIds != nil && len(roomIds) > 0 {
		quest.Param("rids", roomIds)	
	}
	if uids != nil && len(uids) > 0 {
		quest.Param("uids", uids)	
	}
	if events != nil && len(events) > 0 {
		quest.Param("events", events)	
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetListen(groupIds []int64, roomIds []int64, uids []int64, events []string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SetListen() function.")
		}
	}

	quest := client.genServerQuest("setlisten")

	if groupIds != nil {
		quest.Param("gids", groupIds)
	}
	if roomIds != nil {
		quest.Param("rids", roomIds)
	}
	if uids != nil {
		quest.Param("uids", uids)
	}
	if events != nil {
		quest.Param("events", events)
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetListenStatus(allGroups bool, allRrooms bool, allP2P bool, allEvents bool, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SetListen() function.")
		}
	}

	quest := client.genServerQuest("setlisten")
	quest.Param("group", allGroups)
	quest.Param("room", allRrooms)	
	quest.Param("p2p", allP2P)	
	quest.Param("ev", allEvents)

	return client.sendSilentQuest(quest, timeout, callback)
}
