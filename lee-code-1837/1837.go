package main

import (
	"fmt"
)

func sumBase(n int, k int) int {
	ans := 0
	for n > 0 {
		digit := n % k
		ans += digit
		n = n / k
	}
	return ans
}

func main() {
	ans := sumBase(10, 10)
	fmt.Println(ans)
}
