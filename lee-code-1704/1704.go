package main

import "fmt"

func halvesAreAlike(s string) bool {
	m := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	l, r := 0, len(s)-1
	x, y := 0, 0
	for l < r {
		if m[s[l]] {
			x++
		}
		if m[s[r]] {
			y++
		}
		l++
		r--
	}
	return x == y
}

func main() {
	ans := halvesAreAlike("book")
	fmt.Println(ans)
}
