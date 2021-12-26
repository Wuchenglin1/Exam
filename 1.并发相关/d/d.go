package main

import (
	"fmt"
	"sync"
)

//使用并发操作可以提高程序的运行速度，请实现高并发求一百万之内的所有素数
//如果你参考了网络上的文章/程序，请附带上包含你个人理解的注释and被你参考文章/程序的链接

var wg sync.WaitGroup

func main() {
	//要求一百万之内的所有素数，我这里开50个协程来分别计算
	//每个协程计算20000个数

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go judge(i)
	}
	wg.Wait()
}

func judge(i int) {
	fmt.Println("i:", i)
	for k := (i-1)*20000 + 1; k <= i*20000; k++ {
		for m := 2; m <= k; m++ {
			flag := true
			if k%m == 0 {
				flag = false
				break
			}
			if flag {
				fmt.Println("k:", k)
			}
		}
	}
	wg.Done()
}
