package main

import (
	"fmt"
	"strings"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")

	fmt.Println(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) > 0 {
			return len(arr[i])
		}
	}
	return 0
}

func main() {
	//s := "a"
	//ret := lengthOfLastWord(s)
	//fmt.Println(ret)

	ret1 := mySqrt(168)
	fmt.Println(ret1)

}
