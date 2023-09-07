package rtm

import (
	"errors"
	"time"
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
func (client *RTMServerClient) AddGroupMembers(groupId int64, uids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddGroupMembers() function.")
		}
	}

	quest := client.genServerQuest("addgroupmembers")
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
func (client *RTMServerClient) DelGroupMembers(groupId int64, uids []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DelGroupMembers() function.")
		}
	}

	quest := client.genServerQuest("delgroupmembers")
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
func (client *RTMServerClient) DelGroup(groupId int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.DelGroup() function.")
		}
	}

	quest := client.genServerQuest("delgroup")
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
func (client *RTMServerClient) GetGroupMembers(groupId int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetGroupMembers() function.")
		}
	}

	quest := client.genServerQuest("getgroupmembers")
	quest.Param("gid", groupId)

	return client.sendSliceQuest(quest, timeout, "uids", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (count int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (-1, error);
		else this function work in sync mode, and return (count int32, err error)
*/
func (client *RTMServerClient) GetGroupCount(groupId int64, rest ...interface{}) (int32, error) {

	var timeout time.Duration
	var callback func(int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int, string):
			callback = value
		default:
			return -1, errors.New("Invaild params when call RTMServerClient.GetGroupCount() function.")
		}
	}

	quest := client.genServerQuest("getgroupcount")
	quest.Param("gid", groupId)

	return client.sendIntQuest(quest, timeout, "cn", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (ok bool, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (true, error);
		else this function work in sync mode, and return (ok bool, err error)
*/
func (client *RTMServerClient) IsGroupMember(groupId int64, uid int64, rest ...interface{}) (bool, error) {

	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invaild params when call RTMServerClient.IsGroupMember() function.")
		}
	}

	quest := client.genServerQuest("isgroupmember")
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
func (client *RTMServerClient) GetUserGroups(uid int64, rest ...interface{}) ([]int64, error) {

	var timeout time.Duration
	var callback func([]int64, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func([]int64, int, string):
			callback = value
		default:
			return nil, errors.New("Invaild params when call RTMServerClient.GetUserGroups() function.")
		}
	}

	quest := client.genServerQuest("getusergroups")
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
func (client *RTMServerClient) AddGroupBan(groupId int64, uid int64, bannedSeconds int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.AddGroupBan() function.")
		}
	}

	quest := client.genServerQuest("addgroupban")
	if groupId > 0 {
		quest.Param("gid", groupId)
	}
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
func (client *RTMServerClient) RemoveGroupBan(groupId int64, uid int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.RemoveGroupBan() function.")
		}
	}

	quest := client.genServerQuest("removegroupban")
	if groupId > 0 {
		quest.Param("gid", groupId)
	}
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
func (client *RTMServerClient) IsBanOfGroup(groupId int64, uid int64, rest ...interface{}) (bool, error) {

	var timeout time.Duration
	var callback func(bool, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(bool, int, string):
			callback = value
		default:
			return false, errors.New("Invaild params when call RTMServerClient.IsBanOfGroup() function.")
		}
	}

	quest := client.genServerQuest("isbanofgroup")
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
func (client *RTMServerClient) SetGroupInfo(groupId int64, publicInfo *string, privateInfo *string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.SetGroupInfo() function.")
		}
	}

	quest := client.genServerQuest("setgroupinfo")
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
func (client *RTMServerClient) GetGroupInfo(groupId int64, rest ...interface{}) (string, string, error) {

	var timeout time.Duration
	var callback func(string, string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(string, string, int, string):
			callback = value
		default:
			return "", "", errors.New("Invaild params when call RTMServerClient.GetGroupInfo() function.")
		}
	}

	quest := client.genServerQuest("getgroupinfo")
	quest.Param("gid", groupId)

	return client.sendGetObjectInfoQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) ClearProjectGroup(rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("Invaild params when call RTMServerClient.ClearProjectGroup() function.")
		}
	}

	quest := client.genServerQuest("clearprojectgroup")

	return client.sendSilentQuest(quest, timeout, callback)
}
