package main

import "fmt"

//func arrangeCoins(n int) int {
//	k := 0
//	for n > k {
//		k++
//		n -= k
//	}
//	return k
//}

func arrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		k := mid * (mid + 1) / 2
		if k == n {
			return mid
		} else if k < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i, arrangeCoins(i))
	}

}
