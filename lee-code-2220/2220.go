package main

import (
	"fmt"
)

func minBitFlips1(start int, goal int) (ans int) {
	startBit := intToBits(start)
	goalBit := intToBits(goal)
	fmt.Println(startBit)
	fmt.Println(goalBit)

	return
}

func intToBits(num int) []int {
	var bits []int
	for num > 0 {
		bits = append(bits, num%2)
		num /= 2
	}
	return bits
}

func minBitFlips(start int, goal int) int {
	flips := 0
	xorResult := start ^ goal

	for xorResult != 0 {
		xorResult &= (xorResult - 1)
		flips++
	}

	return flips
}

func main() {
	ans := minBitFlips(10, 7)
	fmt.Println(ans)
}
