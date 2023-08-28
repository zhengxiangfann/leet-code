package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func threeSumClosest1(nums []int, target int) int {
	// 对数组进行排序
	sort.Ints(nums)

	// 初始化一个变量来保存与目标值最接近的和
	closestSum := nums[0] + nums[1] + nums[2]

	// 遍历数组中的每个元素，将问题转化为两数之和问题
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// 更新与目标值最接近的和
			if math.Abs(float64(sum-target)) < math.Abs(float64(closestSum-target)) {
				closestSum = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// 如果找到了和等于目标值的情况，直接返回结果
				return target
			}
		}
	}

	return closestSum
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	near := 0
	minNear := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			near = nums[i] + nums[l] + nums[r]
			if abs(near-target) < abs(minNear-target) {
				minNear = near
			}

			if abs(near-target) == 0 {
				return near
			} else if near > target {
				r--
			} else if near < target {
				l++
			}
		}
	}
	return minNear
}

func main() {
	nums := []int{-1, 2, 1, -4}
	ans := threeSumClosest(nums, 1)
	fmt.Println(ans)
}
