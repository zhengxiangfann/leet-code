package main

import "fmt"

// 交替数字和
// alternate

func alternateDigitSum(n int) int {
	sum, flag := 0, 1
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += flag * digit
		flag = -flag
	}
	return -flag * sum
}

func main() {
	ret := alternateDigitSum(1)
	fmt.Println(ret)
}
