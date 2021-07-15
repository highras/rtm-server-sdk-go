package rtm

import (
	"hash/crc32"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type midGenerator struct {
	mutex        sync.Mutex
	randId       int32
	randBits     int32
	macCode      int32
	macBits      int32
	ipCode       int32
	ipBits       int32
	count        int32
	sequenceBits int32
	sequenceMask int32
	lastTime     int64
}

func getMacAddrCode() int32 {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		return rand.Int31n(15) + 1
	}
	var macAddrs []string
	for _, netInterface := range netInterfaces {
		if (netInterface.Flags&net.FlagUp != 0) && (netInterface.Flags&net.FlagLoopback == 0) {
			macAddr := netInterface.HardwareAddr.String()
			if len(macAddr) == 0 {
				continue
			}

			macAddrs = append(macAddrs, macAddr)
		}
	}

	if len(macAddrs) > 0 {
		data := macAddrs[0]
		return (int32)(crc32.ChecksumIEEE([]byte(data)))%15 + 1
	} else {
		rand.Seed(time.Now().UnixNano())
		return rand.Int31n(15) + 1
	}
}

func getIPCode() int32 {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		return rand.Int31n(15) + 1
	}

	var ips []string
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	var ip_code int32 = 0
	if len(ips) > 0 {
		data := ips[0]
		s := strings.Split(data, ".")
		if len(s) > 3 {
			ip_i, _ := strconv.Atoi(s[3])
			ip_code = (int32)(ip_i)
		}
	}
	if ip_code == 0 {
		rand.Seed(time.Now().UnixNano())
		return rand.Int31n(15) + 1
	} else {
		return ip_code
	}
}

func (gen *midGenerator) genMid() int64 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.randId == 0 {
		rand.Seed(time.Now().UnixNano())
		gen.randId = rand.Int31n(15) + 1
		gen.ipCode = getIPCode()
		gen.macCode = getMacAddrCode()
		log.New(os.Stdout, "[RTM Go SDK] ", log.LstdFlags|log.Lshortfile).Println("GenMid Info: randId =", gen.randId, ",ipCode =", gen.ipCode, ",macCode =", gen.macCode)
	}

	currentMillis := gen.getNewStamp()
	gen.count = (gen.count + 1) & gen.sequenceMask
	if gen.count == 0 {
		currentMillis = gen.getNextTimeMillis(gen.lastTime)
	}
	gen.lastTime = currentMillis
	mid := (currentMillis << (gen.randBits + gen.sequenceBits + gen.macBits + gen.ipBits)) | (int64)(gen.randId<<(gen.sequenceBits+gen.macBits+gen.ipBits)) | (int64)(gen.macCode<<(gen.ipBits+gen.sequenceBits)) | (int64)(gen.ipCode<<gen.sequenceBits) | int64(gen.count)
	return mid
}

func (gen *midGenerator) getNewStamp() int64 {
	now := time.Now()
	currentMillis := now.UnixNano() / 1e6
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
	randBits:     4,
	macCode:      0,
	macBits:      4,
	ipCode:       0,
	ipBits:       8,
	count:        0,
	sequenceBits: 6,
	sequenceMask: -1 ^ (-1 << 6),
	lastTime:     0,
}
