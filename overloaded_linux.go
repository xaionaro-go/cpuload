// +build linux

package cpuload

import (
	"log"
	"os"
	"sync"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

var (
	currentCPULoad      = atomicFloat64(0)
	defaultCPULoadLimit = atomicFloat64(0)
	limits              = sync.Map{}
	ErrorLogger         = log.New(os.Stderr, "[cpuload] ", 0)
)

func init() {
	defaultCPULoadLimit.Set(0.9)
	currentCPULoad.Set(getCPULoad())

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			currentCPULoad.Set(getCPULoad())
		}
	}()
}

func getCPULoad() float64 {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		ErrorLogger.Printf("getCPULoad(): %v", err)
		return -1
	}

	total := uint64(0)
	idle := uint64(0)
	for _, s := range stat.CPUStats {
		total += s.User
		total += s.Nice
		total += s.System
		total += s.Idle
		total += s.IOWait

		idle += s.Idle
	}

	return 1 - float64(idle)/float64(total)
}

func IsOverloadedFor(key string) bool {
	limit := defaultCPULoadLimit.Get()
	if limitI, ok := limits.Load(key); ok {
		limit = limitI.(float64)
	}

	return currentCPULoad.Get() > limit
}

func IsOverloaded() bool {
	return IsOverloadedFor(`everything_else`)
}

func SetCPULoadLimitFor(key string, limit float64) {
	limits.Store(key, limit)
}
