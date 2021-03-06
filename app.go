package main

import (
	"fmt"
	"google"
	"runtime"
	"sync"
	"time"
)

func main() {
	cpuNumber := 3
	runtime.GOMAXPROCS(cpuNumber)
	var wait sync.WaitGroup
	wait.Add(cpuNumber)

	startTime := time.Now()

	go func() {
		for i := 0; i < 10; i++ {
			google.Crawler()
		}

		defer wait.Done()
	}()

	wait.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Printf("実行時間: %s\n", elapsedTime)
}
