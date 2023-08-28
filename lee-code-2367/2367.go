package main

import "fmt"

func arithmeticTriplets1(nums []int, diff int) (ans int) {
	for i := 0; i < len(nums); i++ {
		findHalf := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+diff == nums[j] {
				findHalf = true
			} else if nums[i]+2*diff == nums[j] && findHalf {
				ans++
				break
			}
		}
	}
	return ans
}

func arithmeticTriplets(nums []int, diff int) int {
	ans := 0
	indices := make(map[int]bool)

	for _, num := range nums {
		indices[num] = true
	}

	for i := 0; i < len(nums); i++ {

		if indices[nums[i]+diff] && indices[nums[i]+2*diff] {
			ans++
		}
	}

	return ans
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 9}
	diff := 2
	ans := arithmeticTriplets(nums, diff)
	fmt.Println(ans)
}
