package main

import "fmt"

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func method1(input ...int) []int {
	//进行+1操作
	flag := true
	for i := len(input) - 1; i >= 0; i-- {
		if flag {
			input[i] = (input[i] + 1) % 10
		}
		if input[i] != 0 {
			//没有进位操作
			flag = false
			continue
		} else {
			//有进位操作
			flag = true
		}
	}
	if flag {
		input = append([]int{1}, input...)
	}
	return input
}

func main() {
	input1 := []int{9, 9, 9}
	input2 := []int{1, 3, 9}
	fmt.Println(method1(input1...))
	fmt.Println(method1(input2...))
}
