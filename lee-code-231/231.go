package main

import "fmt"

func isPowerOfTwo(n int) bool {

	if n <= 2 {
		return true
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		} else {
			n >>= 1
		}
	}

	return true
}

func main() {
	ans := isPowerOfTwo(3)
	fmt.Println(ans)
}
