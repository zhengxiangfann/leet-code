package main

import (
	"fmt"
	"strings"
)

func countConsistentStrings(allowed string, words []string) int {
	ans := 0
	flag := false
	for _, word := range words {
		flag = true
		for _, b := range word {
			if !strings.Contains(allowed, string(b)) {
				flag = false
				break
			}
		}

		if flag {
			ans++
		}
	}
	return ans
}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	maxLocal := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		maxLocal[i] = make([]int, n-2)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			maxVal := grid[i][j]
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			maxLocal[i-1][j-1] = maxVal
		}
	}

	return maxLocal
}

func main() {
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	ans := countConsistentStrings(allowed, words)
	fmt.Println(ans)

}
