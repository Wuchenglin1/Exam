package main

import "fmt"

//基于管道，我们可以把打印的协程拓展为N个。
//请在main函数中开启10个协程输出一段话，要求10行话全部输出完毕后再结束main函数。

func main() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println("Hello,golang!")
			ch <- 1
		}()
	}
	for i := 1; i <= 10; i++ {
		<-ch
	}
	fmt.Println("输出完毕")
}
