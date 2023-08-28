package main

import "fmt"

func reverseStr(s string) string {
	l, r := 0, len(s)-1

	str := []byte(s)
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	return string(str)
}

func finalString(s string) string {

	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			ans = reverseStr(ans)
		} else {
			ans += string(s[i])
		}
	}
	return ans
}

func main() {
	//ans := finalString("bbaiabb")
	//fmt.Println(">>>", ans)

	ans := divide(-2147483648, -1)
	fmt.Println(ans)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func divide(a int, b int) int {

	if a == -2147483648 {
		return 2147483647
	}

	positive := (a > 0) == (b > 0)
	a, b = abs(a), abs(b)
	ans := 0
	for a >= b {
		x := b
		multiple := 1
		for a >= (x << 1) {
			x <<= 1
			multiple <<= 1
		}
		a -= x
		ans += multiple
	}
	if positive {
		return ans
	} else {
		return -ans
	}
}

//
//def divide(a, b):
//INT_MAX = 2**31 - 1
//INT_MIN = -2**31
//
//# 处理特殊情况
//if a == INT_MIN and b == -1:
//return INT_MAX
//
//# 计算结果的符号
//positive = (a > 0) == (b > 0)
//a, b = abs(a), abs(b)
//
//result = 0
//while a >= b:
//x = b
//multiple = 1
//while a >= (x << 1):
//x <<= 1
//multiple <<= 1
//a -= x
//result += multiple
//
//return result if positive else -result
