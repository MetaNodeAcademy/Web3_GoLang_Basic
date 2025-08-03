package main

import (
	"fmt"
	"slices"
)

/**
合并区间
*/

// 暴力解法
func method1(input [][]int) (res [][]int) {
	slices.SortFunc(input, func(a, b []int) int {
		return a[0] - b[0]
	})
	for _, p := range input {
		m := len(res)
		if m > 0 && p[0] <= res[m-1][1] {
			res[m-1][1] = max(res[m-1][1], p[1])
		} else {
			res = append(res, p)
		}
	}
	return
}

func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(method1(input))
}
