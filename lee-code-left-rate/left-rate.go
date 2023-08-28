package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	left, right := s[:n], s[n:]
	return right + left
}

func main() {
	ans := reverseLeftWords("abcdefg", 2)
	fmt.Println(ans)
}
