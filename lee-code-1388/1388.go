package main

import "fmt"

func maxSizeSlices(slices []int) int {
	n := len(slices)
	// 传入 start 和 end，表示只考虑从 slices[start] 到 slices[end] 之间的部分
	rob := func(start, end int) int {
		dp := make([][]int, end-start+1)
		for i := range dp {
			dp[i] = make([]int, n/3+1)
		}

		dp[0][1] = slices[start]
		dp[1][1] = max(slices[start], slices[start+1])

		for i := 2; i <= end-start; i++ {
			for j := 1; j <= n/3; j++ {
				dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[start+i])
			}
		}

		return dp[end-start][n/3]
	}
	return max(rob(0, n-2), rob(1, n-1))
}

func main() {
	//nums := []int{8, 9, 8, 6, 1, 1}
	//nums := []int{1, 2, 3, 4, 5, 6}
	nums := []int{1, 2, 3}
	ans := maxSizeSlices(nums)
	fmt.Println(ans)
}
