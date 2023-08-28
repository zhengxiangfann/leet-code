package main

import (
	"fmt"
)

func canBeIncreasing1(nums []int) bool {

	if len(nums) <= 2 {
		return true
	}

	preNum0 := nums[0]
	preNum1 := nums[1]
	numCur := nums[2]

	isDel := false
	for i := 2; i < len(nums); i++ {
		numCur = nums[i]
		if preNum0 < preNum1 && preNum1 < numCur {
			preNum0 = preNum1
			preNum1 = numCur
		} else if preNum0 < preNum1 && preNum1 > numCur && !isDel {
			preNum1 = numCur
			isDel = true
		} else if preNum0 < preNum1 && preNum1 > numCur && isDel {
			return false
		} else if preNum1 == numCur || numCur == preNum0 {
			return false
		} else if preNum0 > preNum1 && preNum1 > numCur {
			return false
		}
	}

	return true
}

func canBeIncreasing2(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	count := 0

	curNum := nums[0]
	preNum := nums[0]
	preNum0 := nums[0]

	for i := 1; i < len(nums); i++ {
		curNum = nums[i]
		preNum = nums[i-1]

		if curNum <= preNum {
			count++
			if count > 1 {
				return false
			}

			if i == 1 {
				preNum = curNum
			} else {
				preNum0 = nums[i-2]
			}

			if curNum > preNum0 {
				preNum = curNum
			}

		}
	}

	return true
}

func canBeIncreasing3(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			count++
			if count > 1 {
				return false
			}
			if i == 1 || nums[i] > nums[i-2] {
				nums[i-1] = nums[i]
			}
		}
	}

	return true
}

func canBeIncreasing(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	count := 0
	prevNum := nums[0]

	for i := 1; i < len(nums); i++ {
		currNum := nums[i]
		if currNum <= prevNum {
			count++
			if count > 1 {
				return false
			}
			if i == 1 || currNum > nums[i-2] {
				prevNum = currNum
			}
		} else {
			prevNum = currNum
		}
	}

	return true
}

func main() {

	nums := []int{1, 1, 2, 100}
	ans := canBeIncreasing(nums)
	fmt.Println(ans)
}
