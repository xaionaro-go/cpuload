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

func Test_OverloadedStateFor(t *testing.T) {
	SetCPULoadLimitFor("qwerty", 100)
	current, limit := OverloadedStateFor("qwerty")
	if limit != 100 {
		t.Errorf("invalid state returned")
	}
	if current < 0 || current > 1 {
		t.Errorf("invalid current load returned")
	}
}