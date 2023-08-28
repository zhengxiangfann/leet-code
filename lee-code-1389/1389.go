package main

import (
	"fmt"
)

//func createTargetArray(nums []int, index []int) []int {
//	ans := make([]int, 0)
//	for _, idx := range index {
//		ans = append(ans[:idx], append([]int{nums[idx]}, ans[idx:]...)...)
//	}
//
//	return ans
//}

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		idx := index[i]
		target = append(target[:idx], append([]int{nums[i]}, target[idx:]...)...)
	}

	return target
}

func main() {

	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}

	ans := createTargetArray(nums, index)
	fmt.Println(ans)
}
