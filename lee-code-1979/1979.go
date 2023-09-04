package main

import "fmt"

func findGCD(nums []int) int {
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}

		if num < minNum {
			minNum = num
		}
	}

	for minNum != 0 {
		maxNum, minNum = minNum, maxNum%minNum
	}

	return maxNum

}

func main() {

	nums := []int{2, 5, 6, 9, 10}
	ans := findGCD(nums)
	fmt.Println(ans)
}
