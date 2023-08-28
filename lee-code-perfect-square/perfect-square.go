package main

import "fmt"

func numSquares(n int) int {
	// 创建一个切片用于存储最小数量
	dp := make([]int, n+1)

	// 遍历从1到n的每个数字
	for i := 1; i <= n; i++ {
		// 将最小数量初始化为最大值，即假设每个数字都由1组成
		dp[i] = i

		// 遍历从1到i的每个完全平方数
		for j := 1; j*j <= i; j++ {
			// 更新最小数量为当前数量和使用完全平方数后的数量中的较小值
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ret := numSquares(12)
	fmt.Println(ret)
}
