package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func countKDifference1(nums []int, k int) (ans int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}
	return
}

func countKDifference(nums []int, k int) int {
	freqMap := make(map[int]int)
	ans := 0
	for _, num := range nums {
		ans += freqMap[num-k] + freqMap[num+k]
		freqMap[num]++
	}
	fmt.Println(freqMap, ans)

	return ans
}

func main() {
	arr := []int{1, 3}
	countKDifference(arr, 1)
}
