package cpuload

import (
	"testing"
)

func Test_getCPULoad(t *testing.T) {
	cpuLoad := getCPULoad()
	if cpuLoad < 0 || cpuLoad > 1 {
		t.Errorf("Erroneuos cpuload: %v", cpuLoad)
	}
}
