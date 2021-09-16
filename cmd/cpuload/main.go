package main

import (
	"context"
	"fmt"
	"time"

	"github.com/xaionaro-go/cpuload"
)

func main() {
	m := cpuload.NewMonitor(context.Background(), time.Second)

	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println(m.GetCPULoad())
		<-ticker.C
	}
}
