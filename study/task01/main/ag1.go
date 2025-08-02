package main

import "fmt"

//只出现过一次的数字，给定一个非空整数数组，除了某个元素出现一次之外，其余每个元素均出现两次，找出那个只出现一次的元素

// 第一种：暴力破解
func method1(arr []int) int {
	for i := 0; i < len(arr); i++ {
		index := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if index == arr[j] {
				break
			}
			if j == len(arr)-1 {
				return index
			}
		}
	}
	return -1
}

// 第二种：使用map存储key value 节省时间复杂度
func method2(arr []int) int {
	var compareM map[int]int = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		_, exist := compareM[arr[i]]
		if exist {
			delete(compareM, arr[i])
		} else {
			compareM[arr[i]] = 1
		}
	}
	result := -1
	for k := range compareM {
		result = k
	}
	return result
}

// 第三种：二进制的方式解决
func method3(arr []int) int {
	result := 0
	for k, _ := range arr {
		result ^= arr[k]
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 2, 3, 1}
	fmt.Println(method1(arr))
	fmt.Println(method2(arr))
	fmt.Println(method3(arr))
}
