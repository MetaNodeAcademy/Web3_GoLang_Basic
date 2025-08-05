package main

import (
	"fmt"
	"time"
)

// 接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func main() {

	ch := make(chan int, 3)

	go sendOnly(ch)
	go receiveOnly(ch)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("no data, waiting")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
