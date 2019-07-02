// +build !linux,!darwin

package cpuload

func getCPULoad() float64 {
	panic("Not implemented")
}
