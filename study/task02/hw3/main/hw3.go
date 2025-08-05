package main

import (
	"fmt"
	"time"
)

/*
*
使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/

func main() {
	i := 1
	j := 1
	go func() {
		//奇数
		for i <= 10 {
			if i%2 != 0 {
				fmt.Printf("这是一个奇数: %d\n", i)
			}
			i++
		}
	}()
	go func() {
		//奇数
		for j <= 10 {
			if j%2 == 0 {
				fmt.Printf("这是一个偶数: %d\n", j)
			}
			j++
		}

	}()
	time.Sleep(2 * time.Second)
}
