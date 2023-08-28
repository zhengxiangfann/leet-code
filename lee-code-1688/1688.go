package main

import "fmt"

func numberOfMatches(n int) int {
	ans := 0

	for n > 1 {
		if n%2 == 0 {
			ans += n / 2
			n = n / 2
		} else {
			ans += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return ans
}
func numberOfMatches1(n int) int {
	//因为在锦标赛中，每轮淘汰一半的队伍，最终只剩下一支队伍，所以比赛场次为 'n - 1'。
	return n - 1
}

func main() {
	ans := numberOfMatches(14)
	fmt.Println(ans)
}
