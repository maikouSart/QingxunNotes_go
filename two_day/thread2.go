package main

import "sync"

//协程的进阶
func ManyGoWait() {
	var wg sync.WaitGroup
	//指定5个协程
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			//告知协程完成
			defer wg.Done()
			// 相关方法
		}()
	}
	wg.Wait()
}
