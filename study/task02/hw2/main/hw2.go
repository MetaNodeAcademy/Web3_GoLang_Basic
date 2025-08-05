package main

import (
	"fmt"
)

/*
*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func inc(data *[]int) {
	for i := 0; i < len(*data); i++ {
		(*data)[i] *= 2
	}
}

func main() {
	i := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &i)
	inc(&i)
	fmt.Println(i)
}
