package rtm

import (
	"log"
	"time"
	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const (
	dupFliterCleanIntervalSeconds = 5 * 60
	dupFilterTriggerCleanCount = 1000
)

//=======================[ dupMessageFilter ]========================//

type dupP2PMessageKey struct {
	sender		int64
	receiver	int64
	mid			int64
}

type dupGroupMessageKey struct {
	sender		int64
	groupId		int64
	mid			int64
}

type dupRoomMessageKey struct {
	sender		int64
	roomId		int64
	mid			int64
}

type dupMessageFilter struct {
	p2pCache		map[dupP2PMessageKey]int64
	groupCache		map[dupGroupMessageKey]int64
	roomCache		map[dupRoomMessageKey]int64
}

func newDupMessageFilter() *dupMessageFilter {
	filter := &dupMessageFilter{}
	filter.p2pCache = make(map[dupP2PMessageKey]int64)
	filter.groupCache = make(map[dupGroupMessageKey]int64)
	filter.roomCache = make(map[dupRoomMessageKey]int64)
	return filter
}

func (filter *dupMessageFilter) checkP2PMessage(sender int64, receiver int64, mid int64) bool {

	key := dupP2PMessageKey{sender, receiver, mid}
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

func (filter *dupMessageFilter) checkGroupPMessage(sender int64, groupId int64, mid int64) bool {

	key := dupGroupMessageKey{sender, groupId, mid}
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

func (filter *dupMessageFilter) checkRoomMessage(sender int64, roomId int64, mid int64) bool {
	
	key := dupRoomMessageKey{sender, roomId, mid}
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
	monitor			RTMServerMonitor
	dupFilter		*dupMessageFilter
	logger			*log.Logger
}

func newRTMServerQuestProcessor() *rtmServerQuestProcessor {
	processor := &rtmServerQuestProcessor{}
	processor.dupFilter = newDupMessageFilter()
	return processor
}

func (processor *rtmServerQuestProcessor) Process(method string) func(*fpnn.Quest) (*fpnn.Answer, error) {

	if processor.monitor == nil {
		processor.logger.Printf("[ERROR] RTMServerMonitor is unconfiged.")
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
	case "ping":
		return processor.processPing
	default:
		return nil
	}
}

func (processor *rtmServerQuestProcessor) processPushMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	fromUid := quest.WantInt64("from")
	toUid := quest.WantInt64("to")
	mtype := quest.WantInt8("mtype")

	mid := quest.WantInt64("mid")
	message := quest.WantString("msg")
	attrs := quest.WantString("attrs")
	mtime := quest.WantInt64("mtime")

	if processor.dupFilter.checkP2PMessage(fromUid, toUid, mid) {
		go processor.monitor.P2PMessage(fromUid, toUid, mtype, mid, message, attrs, mtime)
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPushGroupMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	fromUid := quest.WantInt64("from")
	groupId := quest.WantInt64("gid")
	mtype := quest.WantInt8("mtype")

	mid := quest.WantInt64("mid")
	message := quest.WantString("msg")
	attrs := quest.WantString("attrs")
	mtime := quest.WantInt64("mtime")

	if processor.dupFilter.checkGroupPMessage(fromUid, groupId, mid) {
		go processor.monitor.GroupMessage(fromUid, groupId, mtype, mid, message, attrs, mtime)
	}

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPushRoomMessage(quest *fpnn.Quest) (*fpnn.Answer, error) {

	fromUid := quest.WantInt64("from")
	roomId := quest.WantInt64("rid")
	mtype := quest.WantInt8("mtype")

	mid := quest.WantInt64("mid")
	message := quest.WantString("msg")
	attrs := quest.WantString("attrs")
	mtime := quest.WantInt64("mtime")

	if processor.dupFilter.checkRoomMessage(fromUid, roomId, mid) {
		go processor.monitor.RoomMessage(fromUid, roomId, mtype, mid, message, attrs, mtime)
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

	go processor.monitor.Event(pid, event, uid, time, endpoint, data)

	return fpnn.NewAnswer(quest), nil
}

func (processor *rtmServerQuestProcessor) processPing(quest *fpnn.Quest) (*fpnn.Answer, error) {
	return fpnn.NewAnswer(quest), nil
}