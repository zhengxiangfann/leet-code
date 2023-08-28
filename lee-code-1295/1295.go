package main

import "fmt"

func findNumbers(nums []int) int {
	ans := 0
	for _, n := range nums {
		cnt := 0
		for n > 0 {
			n = n / 10
			cnt++
		}
		if cnt%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
