package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	ball := make(chan int)

	go work(ball,"尤旋")
	go work(ball, "杨嘉乐")

	time.Sleep(time.Second * 1)
	fmt.Println("开始")
	ball <- 1

	wg.Wait()
}

func work(ball chan int, play string) {
	defer  wg.Done()
	for  {
		value,ok := <- ball
		if !ok {
			fmt.Printf("%s 赢了\n", play)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0 {
			close(ball)
			return
		}
		fmt.Printf("%s 击打 %d\n", play,value)
		value += 1
		// 将球打给对方
		ball <- value
	}
}
