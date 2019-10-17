package rtm

import (
	"time"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

//-----------[ Group functions ]-------------------//

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uids", uids)

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
func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uids", uids)

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
func (client *RTMServerClient) DelGroup(groupId int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.DelGroup() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("delgroup")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (uids []int64, err error)
*/
func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetGroupMembers() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getgroupmembers")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsGroupMember() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isgroupmember")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (uids []int64, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (nil, error);
		else this function work in sync mode, and return (groupIds []int64, err error)
*/
func (client *RTMServerClient) GetUserGroups(uid int64, rest ... interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func ([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func ([]int64, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetUserGroups() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getusergroups")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("uid", uid)

	return client.sendSliceQuest(quest, timeout, "gids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.AddGroupBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("addgroupban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)
	quest.Param("btime", bannedSeconds)

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
func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.RemoveGroupBan() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("removegroupban")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ... interface{}) (bool, error) {

	var timeout time.Duration
	var callback func (bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (bool, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.IsBanOfGroup() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("isbanofgroup")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)
	quest.Param("uid", uid)

	return client.sendOkQuest(quest, timeout, callback)
}

/*
	Params:
		publicInfo:
			Public info.
			Nil pointer means ignore the params when invoking.

		privateInfo:
			Private info.
			Nil pointer means ignore the params when invoking.

		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ... interface{}) error {

	var timeout time.Duration
	var callback func (int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.SetGroupInfo() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("setgroupinfo")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)

	quest.Param("gid", groupId)

	if publicInfo != nil {
		quest.Param("oinfo", *publicInfo)
	}
	
	if privateInfo != nil {
		quest.Param("pinfo", *privateInfo)
	}

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (publicInfo string, privateInfo string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return ("", "", error);
		else this function work in sync mode, and return (publicInfo string, privateInfo string, err error)
*/
func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ... interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func (string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
			case time.Duration:
				timeout = value
			case func (string, string, int, string):
				callback = value
			default:
				panic("Invaild params when call RTMServerClient.GetGroupInfo() function.")
		}
	}

	sign, salt := client.makeSignAndSalt()

	quest := fpnn.NewQuest("getgroupinfo")
	quest.Param("pid", client.pid)
	quest.Param("sign", sign)
	quest.Param("salt", salt)
	quest.Param("gid", groupId)

	return client.sendGetObjectInfoQuest(quest, timeout, callback)
}
