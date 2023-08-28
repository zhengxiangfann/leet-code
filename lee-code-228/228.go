package main

import "fmt"

func summaryRanges1(nums []int) []string {
	ans := make([]string, 0)
	start := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > 1 {
			if start == nums[i-1] {
				ans = append(ans, fmt.Sprintf("%d", start))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[n-1] {
		ans = append(ans, fmt.Sprintf("%d", start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, nums[n-1]))
	}

	return ans
}

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	n := len(nums)
	if n == 0 {
		return ans
	}

	start := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > 1 {
			if start == nums[i-1] {
				ans = append(ans, fmt.Sprintf("%d", start))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[n-1] {
		ans = append(ans, fmt.Sprintf("%d", start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, nums[n-1]))
	}

	return ans
}

func main() {
	ans := summaryRanges1([]int{0, 3, 4, 6, 8, 10})
	fmt.Println(ans)
}
