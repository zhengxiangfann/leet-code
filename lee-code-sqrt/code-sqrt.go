package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	left, right := 0, num/2
	if num%2 == 1 {
		right += 1
	}
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		}
	}
	return false
}

//func guessNumber(n int) int {
//
//}

func main() {
	for i := 0; i < 200; i++ {
		ret := isPerfectSquare(i)
		if ret == true {
			fmt.Println(i, ret)
		}
	}

}
