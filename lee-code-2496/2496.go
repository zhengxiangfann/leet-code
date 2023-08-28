package main

import (
	"fmt"
	"strconv"
)

func maximumValue(strs []string) int {
	ans := 0

	var checkDigit = func(str string) bool {
		for _, s := range str {
			if '9' < s || s < '0' {
				return false
			}
		}
		return true
	}

	for _, s := range strs {

		if checkDigit(s) {
			if n, err := strconv.Atoi(s); err == nil && n > ans {
				ans = n
			}
		} else if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}

func main() {
	nums := []string{"alic3", "bob", "3", "4", "00000"}
	ans := maximumValue(nums)
	fmt.Println(ans)
}
