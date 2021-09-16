// +build darwin,cgo
// @link https://github.com/mackerelio/go-osstat/blob/master/cpu/cpu_darwin_cgo.go

package cpuload

import (
	"sync/atomic"
	"unsafe"
)

// #include <mach/mach_host.h>
// #include <mach/host_info.h>
import "C"

func (m *Monitor) getAccumulatedCPULoad() float64 {
	var (
		cpuLoad C.host_cpu_load_info_data_t
		count   C.mach_msg_type_number_t = C.HOST_CPU_LOAD_INFO_COUNT
		ret                              = C.host_statistics(C.host_t(C.mach_host_self()), C.HOST_CPU_LOAD_INFO, C.host_info_t(unsafe.Pointer(&cpuLoad)), &count)
	)

	if ret != C.KERN_SUCCESS {
		ErrorLogger.Printf("getAccumulatedCPULoad(): host_statistics failed: %d", ret)
		return -1
	}

	total := uint64(cpuLoad.cpu_ticks[C.CPU_STATE_USER])
	total += uint64(cpuLoad.cpu_ticks[C.CPU_STATE_NICE])
	total += uint64(cpuLoad.cpu_ticks[C.CPU_STATE_SYSTEM])
	total += uint64(cpuLoad.cpu_ticks[C.CPU_STATE_IDLE])
	idle := uint64(cpuLoad.cpu_ticks[C.CPU_STATE_IDLE])

	totalDiff := total - atomic.SwapUint64(&m.previousTotal, total)
	idleDiff := idle - atomic.SwapUint64(&m.previousIdle, idle)

	return 1 - float64(idleDiff)/float64(totalDiff)
}
