package rtm

import (
	"errors"
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
func (client *RTMServerClient) IMServer_SetUserInfos(uid int64, infos map[string]string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_SetUserInfos() function")
		}
	}

	quest := client.genServerQuest("imserver_setuserinfos")

	if infos == nil {
		return errors.New("invaild params when call RTMServerClient.IMServer_SetUserInfos() function. infos must not be nil")
	}
	quest.Param("uid", uid)
	quest.Param("infos", infos)
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
func (client *RTMServerClient) IMServer_SetGroupInfos(groupId int64, infos map[string]string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_SetGroupInfos() function")
		}
	}

	quest := client.genServerQuest("imserver_setgroupinfos")

	if infos == nil {
		return errors.New("invaild params when call RTMServerClient.IMServer_SetGroupInfos() function. infos must not be nil")
	}
	quest.Param("gid", groupId)
	quest.Param("infos", infos)
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
func (client *RTMServerClient) IMServer_SetApplyGrant(mtype int32, xid int64, grant_type int32, rest ...interface{}) error {

	var timeout time.Duration
	var extra string
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		case string:
			extra = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_SetApplyGrant() function")
		}
	}

	quest := client.genServerQuest("imserver_setapplygrant")

	quest.Param("type", mtype)
	quest.Param("xid", xid)
	quest.Param("grant_type", grant_type)
	if len(extra) != 0 {
		quest.Param("extra", extra)
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
func (client *RTMServerClient) IMServer_SetInviteGrant(mtype int32, xid int64, invite_type int32, inviteManageType int32, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_SetApplyGrant() function")
		}
	}

	quest := client.genServerQuest("imserver_setinvitegrant")

	quest.Param("type", mtype)
	quest.Param("xid", xid)
	quest.Param("invite_type", invite_type)
	quest.Param("invite_manage_type", inviteManageType)
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
func (client *RTMServerClient) IMServer_CreateGroup(groupId int64, ownerUid int64, infos map[string]string, permissions map[string]string, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_CreateGroup() function")
		}
	}

	if infos == nil {
		return errors.New("invaild params when call RTMServerClient.IMServer_CreateGroup() function. infos must not be nil")
	}
	quest := client.genServerQuest("imserver_creategroup")

	quest.Param("gid", groupId)
	quest.Param("owner_uid", ownerUid)
	quest.Param("infos", infos)
	quest.Param("permissions", permissions)
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
func (client *RTMServerClient) IMServer_DismissGroup(groupId int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_DismissGroup() function")
		}
	}

	quest := client.genServerQuest("imserver_dismissgroup")
	quest.Param("gid", groupId)

	return client.sendSilentQuest(quest, timeout, callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (grantType int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) IMServer_GetApplyGrant(mType int32, xid int64, rest ...interface{}) (int32, error) {

	var timeout time.Duration
	var callback func(int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int, string):
			callback = value
		default:
			return -1, errors.New("invaild params when call RTMServerClient.IMServer_GetApplyGrant() function")
		}
	}

	quest := client.genServerQuest("imserver_getapplygrant")
	quest.Param("type", mType)
	quest.Param("xid", xid)

	return client.sendIntQuest(quest, timeout, "grant_type", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (inviteType int32, inviteManageType int32, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) IMServer_GetInviteGrant(mType int32, xid int64, rest ...interface{}) (int32, int32, error) {

	var timeout time.Duration
	var callback func(int32, int32, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int32, int32, int, string):
			callback = value
		default:
			return -1, -1, errors.New("invaild params when call RTMServerClient.IMServer_GetInviteGrant() function")
		}
	}

	quest := client.genServerQuest("imserver_getinvitegrant")
	quest.Param("type", mType)
	quest.Param("xid", xid)

	return client.sendDoubleIntQuest(quest, timeout, "invite_type", "invite_manage_type", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (infos map[string]map[string]string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) IMServer_GetUserInfos(userIds []int64, rest ...interface{}) (map[string]map[string]string, error) {

	var timeout time.Duration
	var callback func(map[string]map[string]string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(map[string]map[string]string, int, string):
			callback = value
		default:
			return nil, errors.New("invaild params when call RTMServerClient.IMServer_GetUserInfos() function")
		}
	}

	quest := client.genServerQuest("imserver_getuserinfos")
	quest.Param("uids", userIds)

	return client.sendMapNestedQuest(quest, timeout, "infos", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (infos map[string]map[string]string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) IMServer_GetGroupInfos(groupIds []int64, rest ...interface{}) (map[string]map[string]string, error) {

	var timeout time.Duration
	var callback func(map[string]map[string]string, int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(map[string]map[string]string, int, string):
			callback = value
		default:
			return nil, errors.New("invaild params when call RTMServerClient.IMServer_GetGroupInfos() function")
		}
	}

	quest := client.genServerQuest("imserver_getgroupinfos")
	quest.Param("gids", groupIds)

	return client.sendMapNestedQuest(quest, timeout, "infos", callback)
}

/*
	Params:
		rest: can be include following params:
			timeout time.Duration
			func (infos map[string]map[string]string, errorCode int, errInfo string)

		If include func param, this function will enter into async mode, and return (error);
		else this function work in sync mode, and return (err error)
*/
func (client *RTMServerClient) IMServer_AddGroupMembers(gid int64, members []int64, rest ...interface{}) error {

	var timeout time.Duration
	var callback func(int, string)

	for _, value := range rest {
		switch value := value.(type) {
		case time.Duration:
			timeout = value
		case func(int, string):
			callback = value
		default:
			return errors.New("invaild params when call RTMServerClient.IMServer_AddGroupMembers() function")
		}
	}

	quest := client.genServerQuest("imserver_addgroupmembers")
	quest.Param("gid", gid)
	quest.Param("uids", members)

	return client.sendSilentQuest(quest, timeout, callback)
}
