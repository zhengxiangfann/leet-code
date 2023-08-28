package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	acc := 1
	sum := 0
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += digit
		acc *= digit
	}
	return acc - sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func massage1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < length; i++ {
		// 预约本次:    dp[i-2] + nums[i]
		// 不预约本次:   dp[i-1]
		// dp[i] =  max( 预发本次，不预约本次 )
		// dp[i] = max(dp[i-2] + nums[i],  dp[i-1])
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}

func massage(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < len(nums); i++ {

		cur0 := max(dp0, dp1)
		cur1 := dp0 + nums[i]

		dp0 = cur0
		dp1 = cur1

	}
	return max(dp0, dp1)
}

func Test_massage() {
	nums := []int{2, 1, 1, 2}
	ans := massage(nums)
	fmt.Println(ans)
}

func main() {
	Test_massage()
}
