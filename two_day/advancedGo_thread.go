package main

//协程的使用方法
import (
	"fmt"
	"sync"
	"time"
)

//1. 通信实现共享内存
func CalSquare() {
	src := make(chan int)
	//缓冲区
	dest := make(chan int, 3)

	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	// 生产者
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	// 消费者
	for i := range dest {
		println(i)
	}
}

//2.通信实现共享内存(实现并发加一操作)
var (
	x    int64
	lock sync.Mutex
)

func addwithlock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addwithoutlock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addwithoutlock()
	}
	println(x)
	time.Sleep(time.Second)
	x = 0
	for i := 0; i < 5; i++ {
		go addwithlock()
	}
	time.Sleep(time.Second)
	println(x)
}

func hello(i int) {
	println("hello, goroutine:" + fmt.Sprint(i))
}

func main() {
	add()
	// 协程
	//for i := 0; i < 5; i++ {
	//	// 开启协程
	//	go func(j int) {
	//		hello(j)
	//	}(i)
	//}
	//time.Sleep(time.Second)
}
