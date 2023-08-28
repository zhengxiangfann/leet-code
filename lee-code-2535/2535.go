package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func differenceOfSum(nums []int) int {

	arrSum := 0
	elemSum := 0

	for _, v := range nums {

		arrSum += v

		for v > 0 {
			elemSum += v % 10
			v = v / 10
		}
	}

	return abs(arrSum - elemSum)
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	ans := differenceOfSum(nums)
	fmt.Println(ans)
}
