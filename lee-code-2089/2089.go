package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return ans
		} else if nums[i] == target {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 5, 2, 4}
	ans := targetIndices(nums, 3)
	fmt.Println(ans)
}
