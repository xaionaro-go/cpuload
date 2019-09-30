package cpuload

import (
	"log"
	"os"
	"sync"
)

var (
	defaultCPULoadLimit = atomicFloat64(0)
	limits              = sync.Map{}
	ErrorLogger         = log.New(os.Stderr, "[cpuload] ", 0)
)

func init() {
	defaultCPULoadLimit.Set(0.9)
}

func IsOverloadedFor(key string) bool {
	limit := defaultCPULoadLimit.Get()
	if limitI, ok := limits.Load(key); ok {
		limit = limitI.(float64)
	}

	return GetCPULoad() > limit
}

// OverloadedStateFor - returns current CPU usage and corresponding limit
// for a given key. Used for custom logic to determine overloaded state.
func OverloadedStateFor(key string) (current, limit float64) {
	limit = defaultCPULoadLimit.Get()
	if limitI, ok := limits.Load(key); ok {
		limit = limitI.(float64)
	}

	return GetCPULoad(), limit
}

func IsOverloaded() bool {
	return IsOverloadedFor(`everything_else`)
}

func SetCPULoadLimitFor(key string, limit float64) {
	limits.Store(key, limit)
}
