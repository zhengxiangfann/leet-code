package main

import "fmt"

func removePalindromeSub(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	isPalindrome := func(str string) bool {
		l, r := 0, len(str)-1
		for l < r {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	if isPalindrome(s) {
		return 1
	}

	return 2
}

func main() {
	ans := removePalindromeSub("abbaab")
	fmt.Println(ans)
}
