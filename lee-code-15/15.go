package main

import (
	"fmt"
)

func sort3(x [3]int) (ans [3]int) {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	ans = x
	return ans
}

func threeSum(nums []int) (ans [][]int) {
	m := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					tmp := [3]int{nums[i], nums[j], nums[k]}
					tmp1 := sort3(tmp)
					fmt.Println("sort after: ", tmp1)

					if !m[tmp1] {
						ans = append(ans, []int{nums[i], nums[j], nums[k]})
					} else {
						m[tmp1] = true
					}

				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}
