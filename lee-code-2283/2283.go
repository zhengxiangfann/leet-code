package main

import (
	"fmt"
)

func digitCount(num string) bool {

	m := make([]int, 10)

	for _, s := range num {
		digit := int(s - '0')
		m[digit]++
	}

	for i, s := range num {
		digit := int(s - '0')
		if m[i] != digit {
			return false
		}
	}

	return true
}

func main() {
	ans := digitCount("030")
	fmt.Println(ans)
}
