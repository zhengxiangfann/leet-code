package main

import (
	"fmt"
)

func cellsInRange(s string) []string {
	var list []string
	col1, row1 := s[0], s[1]
	col2, row2 := s[3], s[4]

	for j := col1; j <= col2; j++ {
		for i := row1; i <= row2; i++ {
			list = append(list, fmt.Sprintf("%c%c", j, i))
		}
	}

	return list
}

func main() {
	ans := cellsInRange("K1:L2")
	fmt.Println(ans)
}
