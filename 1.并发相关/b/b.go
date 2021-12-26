package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("下山的路又堵起了")
		ch <- 1
	}()
	<-ch
}
