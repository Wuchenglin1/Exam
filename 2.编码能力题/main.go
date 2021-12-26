package main

import (
	"fmt"
	"math/rand"
)

var sum int
var k int

func main() {
	m := dua(1, 2)
	fmt.Println(m)
}

func dua(n, m int) int {
	fmt.Println(n, m)
	//n为层数，m为房间数
	for i := 1; i <= n; i++ {

		for {
			k = rand.Intn(m)
			if k != 0 {
				break
			}
		}
		//生成每一层有多少个有楼梯的房间

		for j := 1; j <= m; j++ {

		}
	}
}
