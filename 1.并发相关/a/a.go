package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
		//开启协程，有的时候协程中的程序还未运行，主程序就已经结束了，但是mu.Lock还未运行，就运行了mu.Unlock，就导致了互斥锁出错了
	}()

	mu.Unlock()
}
