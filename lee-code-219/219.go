package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok && (i-v) <= k {
			return true
		} else {
			m[nums[i]] = i
		}
	}
	return false

}

func main() {
	nums := []int{1, 2, 3, 1}
	ans := containsNearbyDuplicate(nums, 3)
	fmt.Println(ans)
}
