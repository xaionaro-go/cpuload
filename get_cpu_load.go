package cpuload

import (
	"context"
	"time"
)

type Monitor struct {
	currentCPULoad atomicFloat64
	previousTotal  uint64
	previousIdle   uint64
}

func NewMonitor(ctx context.Context, interval time.Duration) *Monitor {
	m := &Monitor{}
	m.startMonitoring(ctx, interval)
	return m
}

func (m *Monitor) startMonitoring(ctx context.Context, interval time.Duration) {
	m.currentCPULoad.Set(m.getAccumulatedCPULoad())

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				m.currentCPULoad.Set(m.getAccumulatedCPULoad())
			}
		}
	}()
}

func (m *Monitor) GetCPULoad() float64 {
	return m.currentCPULoad.Get()
}
