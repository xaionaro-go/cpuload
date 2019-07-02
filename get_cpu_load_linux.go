// +build linux

package cpuload

import (
	"sync/atomic"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

var (
	previousTotal uint64
	previousIdle  uint64
)

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

	totalDiff := total - atomic.SwapUint64(&previousTotal, total)
	idleDiff := idle - atomic.SwapUint64(&previousIdle, idle)

	return 1 - float64(idleDiff)/float64(totalDiff)
}
