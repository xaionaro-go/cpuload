// +build darwin,!cgo

package cpuload

// getCPULoad for darwin is not supported without CGO
func getCPULoad() float64 {
	return 0
}
