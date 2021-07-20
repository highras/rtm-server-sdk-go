package rtm

import (
	"math/rand"
	"sync"
	"time"
)

type midGenerator struct {
	mutex        sync.Mutex
	randId       int32
	count        int32
	randBits     int32
	sequenceBits int32
	sequenceMask int32
	lastTime     int64
}

func (gen *midGenerator) genMid() int64 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.randId == 0 {
		rand.Seed(time.Now().UnixNano())
		gen.randId = rand.Int31n(255) + 1
	}

	currentMillis := gen.getNewStamp()
	gen.count = (gen.count + 1) & gen.sequenceMask
	if gen.count == 0 {
		currentMillis = gen.getNextTimeMillis(gen.lastTime)
	}
	gen.lastTime = currentMillis
	mid := (currentMillis << (gen.randBits + gen.sequenceBits)) | (int64)((gen.randId)<<(gen.sequenceBits)) | int64(gen.count)
	return mid
}

func (gen *midGenerator) getNewStamp() int64 {
	now := time.Now()
	currentMillis := now.UnixNano() / 1000000
	return currentMillis
}

func (gen *midGenerator) getNextTimeMillis(lastTime int64) int64 {
	curr := gen.getNewStamp()
	for curr <= lastTime {
		curr = gen.getNewStamp()
	}
	return curr
}

var idGen = &midGenerator{
	randId:       0,
	count:        0,
	randBits:     8,
	sequenceBits: 6,
	sequenceMask: -1 ^ (-1 << 6),
	lastTime:     0,
}
