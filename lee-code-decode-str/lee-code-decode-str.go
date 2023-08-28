package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {

	stack := make([]string, 0)
	countStack := make([]int, 0)
	currStr := ""
	count := 0

	for _, char := range s {
		if char >= '0' && char <= '9' {
			// 处理数字
			digit, _ := strconv.Atoi(string(char))
			count = count*10 + digit
		} else if char == '[' {
			// 将当前的count和currStr分别入栈，并重置count和currStr
			stack = append(stack, currStr)
			countStack = append(countStack, count)
			currStr = ""
			count = 0
		} else if char == ']' {
			// 出栈，拼接字符串
			prevCount := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			currStr = prevStr + repeatString(currStr, prevCount)
		} else {
			// 字母字符，直接添加到currStr
			currStr += string(char)
		}
	}
	return currStr
}

// 辅助函数：重复字符串n次
func repeatString(str string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += str
	}
	return result
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}
