package rtm

import (
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	dupFliterCleanIntervalSeconds = 5 * 60
	dupFilterTriggerCleanCount    = 1000
)

//=======================[ dupMessageFilter ]========================//

type dupP2PMessageKey struct {
	sender    int64
	receiver  int64
	messageId int64
}

type dupGroupMessageKey struct {
	sender    int64
	groupId   int64
	messageId int64
}

type dupRoomMessageKey struct {
	sender    int64
	roomId    int64
	messageId int64
}

type dupMessageFilter struct {
	p2pCache   map[dupP2PMessageKey]int64
	groupCache map[dupGroupMessageKey]int64
	roomCache  map[dupRoomMessageKey]int64
}

func newDupMessageFilter() *dupMessageFilter {
	filter := &dupMessageFilter{}
	filter.p2pCache = make(map[dupP2PMessageKey]int64)
	filter.groupCache = make(map[dupGroupMessageKey]int64)
	filter.roomCache = make(map[dupRoomMessageKey]int64)
	return filter
}

func (filter *dupMessageFilter) checkP2PMessage(sender int64, receiver int64, messageId int64) bool {

	key := dupP2PMessageKey{sender, receiver, messageId}
	_, ok := filter.p2pCache[key]
	curr := time.Now().Unix()

	if len(filter.p2pCache) > dupFilterTriggerCleanCount {

		threshold := curr - dupFliterCleanIntervalSeconds
		oldKeys := make(map[dupP2PMessageKey]int64)

		for k, v := range filter.p2pCache {
			if v <= threshold {
				oldKeys[k] = v
			}
		}

		for k, _ := range oldKeys {
			delete(filter.p2pCache, k)
		}
	}

	filter.p2pCache[key] = curr

	return !ok
}

func (filter *dupMessageFilter) checkGroupMessage(sender int64, groupId int64, messageId int64) bool {

	key := dupGroupMessageKey{sender, groupId, messageId}
	_, ok := filter.groupCache[key]
	curr := time.Now().Unix()

	if len(filter.groupCache) > dupFilterTriggerCleanCount {

		threshold := curr - dupFliterCleanIntervalSeconds
		oldKeys := make(map[dupGroupMessageKey]int64)

		for k, v := range filter.groupCache {
			if v <= threshold {
				oldKeys[k] = v
			}
		}

		for k, _ := range oldKeys {
			delete(filter.groupCache, k)
		}
	}

	filter.groupCache[key] = curr

	return !ok
}

func (filter *dupMessageFilter) checkRoomMessage(sender int64, roomId int64, messageId int64) bool {

	key := dupRoomMessageKey{sender, roomId, messageId}
	_, ok := filter.roomCache[key]
	curr := time.Now().Unix()

	if len(filter.roomCache) > dupFilterTriggerCleanCount {

		threshold := curr - dupFliterCleanIntervalSeconds
		oldKeys := make(map[dupRoomMessageKey]int64)

		for k, v := range filter.roomCache {
			if v <= threshold {
				oldKeys[k] = v
			}
		}

		for k, _ := range oldKeys {
			delete(filter.roomCache, k)
		}
	}

	filter.roomCache[key] = curr

	return !ok
}

//=======================[ rtmServerQuestProcessor ]========================//

type rtmServerQuestProcessor struct {
	monitor    RTMServerMonitor
	newMonitor IRTMServerMonitor
	dupFilter  *dupMessageFilter
	logger     RTMLogger
}

func newRTMServerQuestProcessor() *rtmServerQuestProcessor {
	processor := &rtmServerQuestProcessor{}
	processor.dupFilter = newDupMessageFilter()
	return processor
}

func (processor *rtmServerQuestProcessor) Process(method string) func(*fpnn.Quest) (*fpnn.Answer, error) {

	if method == "ping" {
		return processor.processPing
	}

	if processor.monitor == nil && processor.newMonitor == nil {
		processor.logger.Println("[ERROR] RTMServerMonitor is unconfiged.")
		return nil
	}

	switch method {
	case "pushmsg":
		return processor.processPushMessage
	case "pushgroupmsg":
		return processor.processPushGroupMessage
	case "pushroommsg":
		return processor.processPushRoomMessage
	case "pushevent":
		return processor.processPushEvent
	default:
		return nil
	}
}

func (processor *rtmServerQuestProcessor) processPushMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	rtmMessage := &RTMMessage{}
	rtmMessage.FromUid = quest.WantInt64("from")
	rtmMessage.ToId = quest.WantInt64("to")
	rtmMessage.MessageType = quest.WantInt8("mtype")

	rtmMessage.MessageId = quest.WantInt64("mid")
	rtmMessage.Attrs = quest.WantString("attrs")
	rtmMessage.ModifiedTime = quest.WantInt64("mtime")
	rtmMessage.Message = quest.WantString("msg")

	if processor.dupFilter.checkP2PMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId) {
		if rtmMessage.MessageType == defaultMtype_Chat {
			if processor.monitor != nil {
				go processor.monitor.P2PChat(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.P2PChat(rtmMessage)
			}
		} else if rtmMessage.MessageType == defaultMtype_Cmd {
			if processor.monitor != nil {
				go processor.monitor.P2PCmd(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.P2PCmd(rtmMessage)
			}
		} else if rtmMessage.MessageType >= defaultMtype_Image && rtmMessage.MessageType <= defaultMtype_File {
			if processor.monitor != nil {
				go processor.monitor.P2PFile(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				fileInfo := processFileInfo(rtmMessage.Message, rtmMessage.Attrs, rtmMessage.MessageType, processor.logger)
				rtmMessage.FileInfo = fileInfo
				rtmMessage.Attrs = fetchFileCustomAttrs(rtmMessage.Attrs, processor.logger)
				go processor.newMonitor.P2PFile(rtmMessage)
			}
		} else {
			if processor.monitor != nil {
				go processor.monitor.P2PMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.P2PMessage(rtmMessage)
			}
		}
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPushGroupMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	rtmMessage := &RTMMessage{}
	rtmMessage.FromUid = quest.WantInt64("from")
	rtmMessage.ToId = quest.WantInt64("gid")
	rtmMessage.MessageType = quest.WantInt8("mtype")
	rtmMessage.MessageId = quest.WantInt64("mid")
	rtmMessage.Attrs = quest.WantString("attrs")
	rtmMessage.ModifiedTime = quest.WantInt64("mtime")

	rtmMessage.Message = quest.WantString("msg")
	if processor.dupFilter.checkGroupMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId) {
		if rtmMessage.MessageType == defaultMtype_Chat {
			if processor.monitor != nil {
				go processor.monitor.GroupChat(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.GroupChat(rtmMessage)
			}
		} else if rtmMessage.MessageType == defaultMtype_Cmd {
			if processor.monitor != nil {
				go processor.monitor.GroupCmd(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.GroupCmd(rtmMessage)
			}
		} else if rtmMessage.MessageType >= defaultMtype_Image && rtmMessage.MessageType <= defaultMtype_File {
			if processor.monitor != nil {
				go processor.monitor.GroupFile(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				fileInfo := processFileInfo(rtmMessage.Message, rtmMessage.Attrs, rtmMessage.MessageType, processor.logger)
				rtmMessage.FileInfo = fileInfo
				rtmMessage.Attrs = fetchFileCustomAttrs(rtmMessage.Attrs, processor.logger)
				go processor.newMonitor.GroupFile(rtmMessage)
			}
		} else {
			if processor.monitor != nil {
				go processor.monitor.GroupMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.GroupMessage(rtmMessage)
			}
		}
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPushRoomMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	rtmMessage := &RTMMessage{}
	rtmMessage.FromUid = quest.WantInt64("from")
	rtmMessage.ToId = quest.WantInt64("rid")
	rtmMessage.MessageType = quest.WantInt8("mtype")
	rtmMessage.MessageId = quest.WantInt64("mid")
	rtmMessage.Attrs = quest.WantString("attrs")
	rtmMessage.ModifiedTime = quest.WantInt64("mtime")

	rtmMessage.Message = quest.WantString("msg")
	if processor.dupFilter.checkRoomMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId) {
		if rtmMessage.MessageType == defaultMtype_Chat {
			if processor.monitor != nil {
				go processor.monitor.RoomChat(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.RoomChat(rtmMessage)
			}
		} else if rtmMessage.MessageType == defaultMtype_Cmd {
			if processor.monitor != nil {
				go processor.monitor.RoomCmd(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.RoomCmd(rtmMessage)
			}
		} else if rtmMessage.MessageType >= defaultMtype_Image && rtmMessage.MessageType <= defaultMtype_File {
			if processor.monitor != nil {
				go processor.monitor.RoomFile(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				fileInfo := processFileInfo(rtmMessage.Message, rtmMessage.Attrs, rtmMessage.MessageType, processor.logger)
				rtmMessage.FileInfo = fileInfo
				rtmMessage.Attrs = fetchFileCustomAttrs(rtmMessage.Attrs, processor.logger)
				go processor.newMonitor.RoomFile(rtmMessage)
			}
		} else {
			if processor.monitor != nil {
				go processor.monitor.RoomMessage(rtmMessage.FromUid, rtmMessage.ToId, rtmMessage.MessageType, rtmMessage.MessageId,
					rtmMessage.Message, rtmMessage.Attrs, rtmMessage.ModifiedTime)
			} else if processor.newMonitor != nil {
				go processor.newMonitor.RoomMessage(rtmMessage)
			}
		}
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPushEvent(quest *fpnn.Quest) (*fpnn.Answer, error) {

	pid := quest.WantInt32("pid")
	event := quest.WantString("event")
	uid := quest.WantInt64("uid")

	time := quest.WantInt32("time")
	endpoint := quest.WantString("endpoint")
	data, _ := quest.GetString("data")

	if processor.monitor != nil {
		go processor.monitor.Event(pid, event, uid, time, endpoint, data)
	} else if processor.newMonitor != nil {
		go processor.newMonitor.Event(pid, event, uid, time, endpoint, data)
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPing(quest *fpnn.Quest) (*fpnn.Answer, error) {
	return fpnn.NewAnswer(quest), nil
}
