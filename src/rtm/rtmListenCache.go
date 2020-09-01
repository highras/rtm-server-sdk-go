package rtm

var emptyExists = struct{}{}

type rtmListenCache struct {
	listenGroupIds map[int64]struct{}
	listenRoomIds  map[int64]struct{}
	listenUids     map[int64]struct{}
	listenEvents   map[string]struct{}
	allP2P         bool
	allGroup       bool
	allRoom        bool
	allEvent       bool
}

func newRtmListenCache() *rtmListenCache {

	cache := &rtmListenCache{}
	cache.listenEvents = make(map[string]struct{})
	cache.listenGroupIds = make(map[int64]struct{})
	cache.listenRoomIds = make(map[int64]struct{})
	cache.listenUids = make(map[int64]struct{})
	cache.allEvent = false
	cache.allGroup = false
	cache.allP2P = false
	cache.allRoom = false
	return cache
}

func (cache *rtmListenCache) addUids(uids []int64) {
	for _, v := range uids {
		cache.listenUids[v] = emptyExists
	}
}

func (cache *rtmListenCache) addGroupIds(groupIds []int64) {
	for _, v := range groupIds {
		cache.listenGroupIds[v] = emptyExists
	}
}

func (cache *rtmListenCache) addRoomIds(roomIds []int64) {
	for _, v := range roomIds {
		cache.listenRoomIds[v] = emptyExists
	}
}

func (cache *rtmListenCache) addEvents(events []string) {
	for _, v := range events {
		cache.listenEvents[v] = emptyExists
	}
}

func (cache *rtmListenCache) removeUids(uids []int64) {
	for _, v := range uids {
		delete(cache.listenUids, v)
	}
}

func (cache *rtmListenCache) removeRoomIds(roomIds []int64) {
	for _, v := range roomIds {
		delete(cache.listenRoomIds, v)
	}
}

func (cache *rtmListenCache) removeGroupIds(groupIds []int64) {
	for _, v := range groupIds {
		delete(cache.listenGroupIds, v)
	}
}

func (cache *rtmListenCache) removeEvents(events []string) {
	for _, v := range events {
		delete(cache.listenEvents, v)
	}
}

func (cache *rtmListenCache) setUids(uids []int64) {
	cache.listenUids = make(map[int64]struct{})
	for _, v := range uids {
		cache.listenUids[v] = emptyExists
	}
}

func (cache *rtmListenCache) setRoomIds(roomIds []int64) {
	cache.listenRoomIds = make(map[int64]struct{})
	for _, v := range roomIds {
		cache.listenRoomIds[v] = emptyExists
	}
}

func (cache *rtmListenCache) setGroupIds(groupIds []int64) {
	cache.listenGroupIds = make(map[int64]struct{})
	for _, v := range groupIds {
		cache.listenGroupIds[v] = emptyExists
	}
}

func (cache *rtmListenCache) setEvents(events []string) {
	cache.listenEvents = make(map[string]struct{})
	for _, v := range events {
		cache.listenEvents[v] = emptyExists
	}
}

func (cache *rtmListenCache) setAllUid(p2p bool) {
	cache.allP2P = p2p
	if p2p {
		cache.listenUids = make(map[int64]struct{})
	}
}

func (cache *rtmListenCache) setAllGroup(group bool) {
	cache.allGroup = group
	if group {
		cache.listenGroupIds = make(map[int64]struct{})
	}
}

func (cache *rtmListenCache) setAllRoom(room bool) {
	cache.allRoom = room
	if room {
		cache.listenRoomIds = make(map[int64]struct{})
	}
}

func (cache *rtmListenCache) setAllEvent(event bool) {
	cache.allEvent = event
	if event {
		cache.listenEvents = make(map[string]struct{})
	}
}

func (cache *rtmListenCache) empty() bool {
	if len(cache.listenUids) == 0 && len(cache.listenGroupIds) == 0 && len(cache.listenRoomIds) == 0 && len(cache.listenEvents) == 0 {
		return true
	}
	return false
}

func (cache *rtmListenCache) isAllFalse() bool {
	if !cache.allP2P && !cache.allGroup && !cache.allEvent && !cache.allRoom {
		return true
	}
	return false
}

func (cache *rtmListenCache) clear() {
	cache.listenEvents = make(map[string]struct{})
	cache.listenGroupIds = make(map[int64]struct{})
	cache.listenRoomIds = make(map[int64]struct{})
	cache.listenUids = make(map[int64]struct{})
	cache.allEvent = false
	cache.allGroup = false
	cache.allP2P = false
	cache.allRoom = false
}
