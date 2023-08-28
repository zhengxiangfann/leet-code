package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestAltitude1(gain []int) (ans int) {
	preHigh := make([]int, len(gain)+1)
	preHigh[0] = 0
	for i := 1; i <= len(gain); i++ {
		preHigh[i] = preHigh[i-1] + gain[i-1]
		ans = max(preHigh[i], ans)
	}
	return ans
}

func largestAltitude(gain []int) (ans int) {

	sum := 0
	for _, v := range gain {
		sum += v
		ans = max(ans, sum)
	}
	return ans
}

func main() {
	nums := []int{-5, 1, 5, 0, -7}
	ans := largestAltitude(nums)
	fmt.Println(ans)
}
