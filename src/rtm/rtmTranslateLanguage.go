package rtm

type RTMTranslateLanguage int

const (
	_                          = iota
	AR    RTMTranslateLanguage = iota //阿拉伯语
	NL                                //荷兰语
	EN                                //英语
	FR                                //法语
	DE                                //德语
	EL                                //希腊语
	ID                                //印度尼西亚语
	IT                                //意大利语
	JA                                //日语
	KO                                //韩语
	NO                                //挪威语
	PL                                //波兰语
	PT                                //葡萄牙语
	RU                                //俄语
	ES                                //西班牙语
	SV                                //瑞典语
	TL                                //塔加路语（菲律宾语）
	TH                                //泰语
	TR                                //土耳其语
	VI                                //越南语
	ZH_CN                             //中文（简体）
	ZH_TW                             //中文（繁体）
)

func (lang RTMTranslateLanguage) String() string {
	switch lang {
	case AR:
		return "ar"
	case NL:
		return "nl"
	case EN:
		return "en"
	case FR:
		return "fr"
	case DE:
		return "de"
	case EL:
		return "el"
	case ID:
		return "id"
	case IT:
		return "it"
	case JA:
		return "ja"
	case KO:
		return "ko"
	case NO:
		return "no"
	case PL:
		return "pl"
	case PT:
		return "pt"
	case RU:
		return "ru"
	case ES:
		return "es"
	case SV:
		return "sv"
	case TL:
		return "tl"
	case TH:
		return "th"
	case TR:
		return "tr"
	case VI:
		return "vi"
	case ZH_CN:
		return "zh_cn"
	case ZH_TW:
		return "zh_tw"
	default:
		return ""
	}
}
