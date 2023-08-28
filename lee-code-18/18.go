package main

import (
	"fmt"
	"sort"
)

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i, j := 0, len(nums)-1; i < j; {

		l, r := i+1, j-1

		for l < r {
			if nums[i]+nums[j]+nums[l]+nums[r] == target {
				ans = append(ans, []int{nums[i], nums[l], nums[r], nums[j]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}

			} else if nums[i]+nums[j]+nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}

		if nums[i]+nums[i+1]+nums[j]+nums[j-1] > target {
			j--
		} else {
			i++
		}

	}

	return ans
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := len(nums) - 1; j > i+2; j-- {

			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}

			l, r := i+1, j-1

			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					ans = append(ans, []int{nums[i], nums[l], nums[r], nums[j]})
					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return ans
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	ans := fourSum(nums, 0)
	fmt.Println(ans)
}
