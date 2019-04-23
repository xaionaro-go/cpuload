package cpuload

import (
	"time"
)

var (
	currentCPULoad = atomicFloat64(0)
)

func init() {
	currentCPULoad.Set(getCPULoad())

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			currentCPULoad.Set(getCPULoad())
		}
	}()
}

func GetCPULoad() float64 {
	return currentCPULoad.Get()
}
