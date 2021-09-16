// +build linux

package cpuload

import (
	"log"
	"os"
	"sync/atomic"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

var (
	ErrorLogger = log.New(os.Stderr, "[cpuload] ", 0)
)

func (m *Monitor) getAccumulatedCPULoad() float64 {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		ErrorLogger.Printf("getAccumulatedCPULoad(): %v", err)
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

	totalDiff := total - atomic.SwapUint64(&m.previousTotal, total)
	idleDiff := idle - atomic.SwapUint64(&m.previousIdle, idle)

	return 1 - float64(idleDiff)/float64(totalDiff)
}
