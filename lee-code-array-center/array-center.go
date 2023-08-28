package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	Sum := 0
	leftSum := 0
	idx := -1
	for _, v := range nums {
		Sum += v
	}

	for i, v := range nums {
		if Sum-leftSum-v == leftSum {
			idx = i
			break
		}
		leftSum += v
	}
	return idx
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	idx := 0
	if target < nums[0] {
		idx = 0
	}
	if target > nums[n-1] {
		idx = n
	} else {
		for i := 0; i < n; i++ {
			if target <= nums[i] {
				idx = i
				break
			}
		}
	}
	return idx
}

func main() {
	arr1 := []int{1, 3}
	ret := searchInsert(arr1, 2)
	fmt.Println(ret)
}
