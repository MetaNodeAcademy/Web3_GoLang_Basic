package main

import "fmt"

/**
两数之和
*/

// 暴力解法
func method1(input []int, target int) []int {
	var matchMap = make(map[int]int)
	for i := 0; i < len(input); i++ {
		if i == 0 {
			matchMap[target-input[i]] = i
		} else {
			_, exist := matchMap[input[i]]
			if exist {
				//return
				return []int{matchMap[input[i]], i}
			} else {
				matchMap[target-input[i]] = i
			}
		}
	}
	return []int{}
}

func main() {
	input := []int{3, 3}
	fmt.Println(method1(input, 6))
}
