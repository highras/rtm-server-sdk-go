package rtm

type RTMErrorCode int

const (
	RTM_EC_INVALID_PIDORUID              RTMErrorCode = 200001
	RTM_EC_INVALID_PID_OR_SIGN           RTMErrorCode = 200002
	RTM_EC_INVALID_FILE_OR_SIGN_OR_TOKEN RTMErrorCode = 200003
	RTM_EC_ATTRS_WITHOUT_SIGN_OR_EXT     RTMErrorCode = 200004
	RTM_EC_INVALID_MTYPE                 RTMErrorCode = 200005
	RTM_EC_SAME_SIGN                     RTMErrorCode = 200006
	RTM_EC_INVALID_FILE_MTYPE            RTMErrorCode = 200007
	RTM_EC_INVALID_SERVER_TIME           RTMErrorCode = 200008

	RTM_EC_FREQUENCY_LIMITED     RTMErrorCode = 200010
	RTM_EC_REFRESH_SCREEN_LIMITE RTMErrorCode = 200011
	RTM_EC_KICKOUT_SELF          RTMErrorCode = 200012

	RTM_EC_FORBIDDEN_METHOD  RTMErrorCode = 200020
	RTM_EC_PERMISSION_DENIED RTMErrorCode = 200021
	RTM_EC_UNAUTHORIZED      RTMErrorCode = 200022
	RTM_EC_DUPLCATED_AUTH    RTMErrorCode = 200023
	RTM_EC_AUTH_DENIED       RTMErrorCode = 200024
	RTM_EC_ADMIN_LOGIN       RTMErrorCode = 200025
	RTM_EC_ADMIN_ONLY        RTMErrorCode = 200026

	RTM_EC_LARGE_MESSAGE_OR_ATTRS       RTMErrorCode = 200030
	RTM_EC_LARGE_FILE_OR_ATTRS          RTMErrorCode = 200031
	RTM_EC_TOO_MANY_ITEMS_IN_PARAMETERS RTMErrorCode = 200032
	RTM_EC_EMPTY_PARAMETER              RTMErrorCode = 200033

	RTM_EC_NOT_IN_ROOM            RTMErrorCode = 200040
	RTM_EC_NOT_GROUP_MEMBER       RTMErrorCode = 200041
	RTM_EC_MAX_GROUP_MEMBER_COUNT RTMErrorCode = 200042
	RTM_EC_NOT_FRIEND             RTMErrorCode = 200043
	RTM_EC_BANNED_IN_GROUP        RTMErrorCode = 200044
	RTM_EC_BANNED_IN_ROOM         RTMErrorCode = 200045
	RTM_EC_EMPTY_GROUP            RTMErrorCode = 200046
	RTM_EC_MAX_ROOM_COUNT         RTMErrorCode = 200047
	RTM_EC_MAX_FRIEND_COUNT       RTMErrorCode = 200048

	RTM_EC_UNSUPPORTED_LANGUAGE RTMErrorCode = 200050
	RTM_EC_EMPTY_TRANSLATION    RTMErrorCode = 200051
	RTM_EC_SEND_TO_SELF         RTMErrorCode = 200052
	RTM_EC_DUPLCATED_MID        RTMErrorCode = 200053
	RTM_EC_SENSITIVE_WORDS      RTMErrorCode = 200054
	RTM_EC_NOT_ONLINE           RTMErrorCode = 200055
	RTM_EC_TRANSLATION_ERROR    RTMErrorCode = 200056
	RTM_EC_PROFANITY_STOP       RTMErrorCode = 200057

	RTM_EC_NO_CONFIG_IN_CONSOLE        RTMErrorCode = 200060
	RTM_EC_UNSUPPORTED_TRASNCRIBE_TYPE RTMErrorCode = 200061

	RTM_EC_MESSAGE_NOT_FOUND RTMErrorCode = 200070

	RTM_EC_UNKNOWN_ERROR RTMErrorCode = 200999
)
