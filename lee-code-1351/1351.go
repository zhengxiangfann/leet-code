package main

import "fmt"

func countNegatives(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				ans++
			} else {
				break
			}
		}
	}
	return ans
}

func binarySearch(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] <= 0 {
			return mid
		} else {
			left = mid + 1
		}
	}
	return 0
}

func main() {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	ans := countNegatives(grid)
	fmt.Println(ans)
}
