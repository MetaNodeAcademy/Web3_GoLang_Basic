package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() //defer 关键词会在方法执行结束之后执行
	c.count++
}

// GetCount 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// UnsafeCounter 线程不安全的计数
type UnsafeCounter struct {
	count int
}

func (c *UnsafeCounter) Increment() {
	c.count++
}
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

func main() {
	counter := SafeCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}
	//
	time.Sleep(time.Second)
	fmt.Println(counter.GetCount())
}
