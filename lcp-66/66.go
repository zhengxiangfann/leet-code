package main

import "fmt"

func minNumBooths(demands []string) int {
	m1 := make(map[int32]int)
	for _, demand := range demands {
		m := make(map[int32]int)

		for _, plat := range demand {
			m[plat]++
		}
		for k, v := range m {
			if v > m1[k] {
				m1[k] = v
			}
		}
	}

	ans := 0
	for _, v := range m1 {
		ans += v
	}
	return ans
}

func main() {

	nums := []string{"ccluro", "mmjhp", "ln", "ayoqwqtqrh", "m", "luhnsb", "gyyy", "auuksw"}
	ans := minNumBooths(nums)
	fmt.Println(ans)
}
