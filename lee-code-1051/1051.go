package main

import (
	"fmt"
	"sort"
)

func heightChecker1(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	//expected := slices.Clone(heights)
	sort.Ints(expected)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			ans++
		}
	}
	return ans
}

func heightChecker(heights []int) int {
	var bucket = [101]int{}
	for i := 0; i < len(heights); i++ {
		bucket[heights[i]]++
	}
	ans := 0
	for i, idx := 0, 0; i < 101; i++ {
		bucket[i]--
		for bucket[i] > 0 {
			if heights[idx] != i {
				ans++
			}
			idx++
		}

	}
	return ans
}

func main() {
	nums := []int{1, 1, 4, 2, 1, 3}
	ans := heightChecker(nums)
	fmt.Println(ans)
}
