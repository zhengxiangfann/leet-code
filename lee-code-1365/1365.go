package main

import "fmt"

func smallerNumbersThanCurrent1(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		small := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] && i != j {
				small++
			}
		}
		ans = append(ans, small)
	}
	return ans
}

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)

	for _, num := range nums {
		count[num]++
	}

	for i := 1; i <= 100; i++ {
		count[i] += count[i-1]
	}
	fmt.Println(count)

	ans := make([]int, len(nums))
	for i, num := range nums {
		if num > 0 {
			ans[i] = count[num-1]
		}
	}

	return ans
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	ans := smallerNumbersThanCurrent(nums)
	fmt.Println(ans)
}
