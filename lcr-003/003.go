package main

func countBits(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}

func main() {
	countBits(5)
}
