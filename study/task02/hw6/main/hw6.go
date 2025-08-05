package main

import (
	"fmt"
	"time"
)

/*
* 编写一个程序，使用通道实现两个协程之间的通信。
* 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
 */

func incNo(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func getNo(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 3)
	go getNo(ch)
	go incNo(ch)
	time.Sleep(2 * time.Second)
}
