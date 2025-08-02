package main

import (
	"fmt"
	"strconv"
)

//整数判断回文数

// 方法一，暴力解决
func method1(input int) bool {
	if input < 0 {
		return false
	}
	var strs string = strconv.Itoa(input)
	i := 0
	j := len(strs) - 1
	result := false
	for {
		if strs[i] == strs[j] {
			if i >= j {
				result = true
				break
			}
			i++
			j--
		} else {
			break
		}
	}
	return result
}

// 方法二 取余
func method2(input int) bool {
	if input < 0 {
		return false
	}
	revertedNumber := 0
	for input > revertedNumber {
		revertedNumber = revertedNumber*10 + input%10
		input /= 10
	}
	return input == revertedNumber || input == revertedNumber/10
}

func main() {
	input1 := 123321
	input2 := 12321
	input3 := 132321
	fmt.Println(method1(input1))
	fmt.Println(method1(input2))
	fmt.Println(method1(input3))
	fmt.Println(method2(input1))
	fmt.Println(method2(input2))
	fmt.Println(method2(input3))
}
