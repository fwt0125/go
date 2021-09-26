package main

import (
	"sync"
)

var wg sync.WaitGroup

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	ch := make(chan string) // 实例化一个管道

	go say("Hello", ch)
	go say("World", ch)

	for {
		println(<-ch) //循环从管道取数据
	}

	wg.Wait()
}
