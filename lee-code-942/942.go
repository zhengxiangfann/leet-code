package main

import "fmt"

func diStringMatch(s string) []int {
	countD, countI := len(s), 0
	ans := make([]int, 0, len(s)+1)
	for _, b := range s {
		if b == 'D' {
			ans = append(ans, countD)
			countD--
		} else {
			ans = append(ans, countI)
			countI++
		}
	}
	lastVal := countI
	if s[len(s)-1] == 'D' {
		lastVal = countD
	}
	ans = append(ans, lastVal)
	return ans
}

func main() {
	in := "IDID"
	ans := diStringMatch(in)
	fmt.Println(ans)
}
