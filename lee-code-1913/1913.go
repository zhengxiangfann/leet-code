package main

import (
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	if len(nums) >= 4 {
		return nums[len(nums)-1]*nums[len(nums)-2] - nums[1]*nums[0]
	}
	return 0
}

func main() {

}
