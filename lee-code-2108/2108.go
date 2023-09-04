package main

import (
	"fmt"
)

func firstPalindrome(words []string) string {
	for _, word := range words {

		for i := 0; i < len(word)/2; i++ {
			if word[i] != word[len(word)-1-i] {
				goto next
			}
		}

		return word
	next:
	}
	return ""
}

func main() {
	nums := []string{"abc", "car", "ada", "racecar", "cool"}
	ans := firstPalindrome(nums)
	fmt.Println(ans)
}
