package main

import (
	"fmt"
	"math"
	"strconv"
)

func buildArray1(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}

func buildArray2(nums []int) []int {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		// 0  0 + 0
		// 1  2 + 1000* 1 1001
		// 2  1 + (1001%1000)*1000  2001
		nums[i] = nums[i] + (nums[nums[i]]%1000)*1000

	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] / 1000
	}
	return nums
}

func buildArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] |= (nums[nums[i]&1023] & 1023) << 10
	}

	for i := 0; i < len(nums); i++ {
		nums[i] >>= 10
	}

	return nums
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

func convertTemperature(celsius float64) (ans []float64) {

	k, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", celsius+273.15), 64)
	h, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", celsius*1.80+32.00), 64)

	ans = append(ans, k, h)

	return ans
}

func main() {
	//arr := []int{0, 2, 1, 5, 3, 4}
	//ans := buildArray(arr)
	//fmt.Println(ans)
	//ans := findDelayedArrivalTime(18, 1)
	//fmt.Println(ans)

	ans := convertTemperature(36.50)
	fmt.Println(ans)

	fmt.Println(math.Round(1.1 - 0.5))
}
