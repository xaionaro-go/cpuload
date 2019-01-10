package cpuload

import (
	"log"
	"os"
	"sync"
	"time"
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
