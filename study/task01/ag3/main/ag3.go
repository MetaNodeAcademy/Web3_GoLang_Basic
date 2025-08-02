package main

import (
	"Web3_GoLang_Basic/study/utils/stack"
	"fmt"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

// 使用栈进行操作
func method1(strs string) bool {
	var vKeyMap = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	var stack = stack.Stack{}
	for i := 0; i < len(strs); i++ {
		if 0 == i || stack.IsEmpty() {
			stack.Push(string(strs[i]))
		} else {
			if vKeyMap[stack.Peek().(string)] == string(strs[i]) {
				stack.Pop()
			} else {
				stack.Push(string(strs[i]))
			}
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}

func main() {
	str1 := "()[]{}"
	str2 := "{[()]}"
	str3 := "(()){}"
	str4 := "{(()]}"
	fmt.Println(method1(str1))
	fmt.Println(method1(str2))
	fmt.Println(method1(str3))
	fmt.Println(method1(str4))
}
