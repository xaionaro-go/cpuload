package main

import (
	"fmt"
	"time"

	"github.com/trafficstars/cpuload"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println(cpuload.GetCPULoad())
		<-ticker.C
	}
}
