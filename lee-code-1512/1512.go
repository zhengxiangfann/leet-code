package main

import "fmt"

func numIdenticalPairs1(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ans++
			}
		}
	}
	return ans
}

func numIdenticalPairs(nums []int) int {
	count := make(map[int]int)
	ans := 0

	for _, num := range nums {
		ans += count[num]
		count[num]++
	}

	return ans
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums))
}
