package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(1)
	wg2.Add(1)
	ballot := make(chan int)
	// 第一位接力者
	go Runner(ballot)

	// 开始比赛
	ballot <- 1

	wg2.Wait()
}

func Runner(ballot chan int) {
	var newRunner int


	runner := <- ballot

	fmt.Printf("Runnber %d Running with baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(ballot)
	}

	time.Sleep(100 * time.Microsecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg2.Done()
		return
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	ballot <- newRunner
}
