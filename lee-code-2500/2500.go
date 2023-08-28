package main

import (
	"fmt"
	"sort"
)

func deleteGreatestValue1(grid [][]int) (ans int) {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	for j := 0; j < len(grid[0]); j++ {
		maxNum := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > maxNum {
				maxNum = grid[i][j]
			}
		}

		ans += maxNum
	}
	return
}

func deleteGreatestValue(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	ans := 0

	for j := 0; j < cols; j++ {
		maxNum := 0
		for i := 0; i < rows; i++ {
			if grid[i][j] > maxNum {
				maxNum = grid[i][j]
			}
		}
		fmt.Println("maxNUM", maxNum)
		ans += maxNum
	}

	return ans
}

func main() {
	nums := [][]int{{1, 2, 4}, {3, 3, 1}}
	ans := deleteGreatestValue(nums)
	fmt.Println(ans)
}
