package main

import "fmt"

/**
最长公共前缀
*/

// 暴力解法
// 查询三个中最短的字符串进行循环
func method1(strArr []string) string {
	result := ""
	if len(strArr) < 2 {
		return result
	}
	//比对第一个和第二个字符串
	result = findCommonStr(strArr[0], strArr[1])
	for i := 2; i < len(strArr); i++ {
		result = findCommonStr(result, strArr[i])
	}
	return result
}

func findCommonStr(str1 string, str2 string) string {
	length := func() int {
		if len(str1) > len(str2) {
			return len(str2)
		}
		return len(str1)
	}()
	commonStr := ""
	for i := 0; i < length; i++ {
		if str1[i] == str2[i] {
			commonStr += string(str1[i])
		} else {
			break
		}
	}
	return commonStr
}

func main() {
	strArr := []string{"flower", "flow", "flight"}
	fmt.Println(method1(strArr))
}
