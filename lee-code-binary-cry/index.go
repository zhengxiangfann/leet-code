package main

import (
	"fmt"
)

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	binaryLen := len(colsum)
	var ret [][]int
	for i := 0; i < 2; i++ {
		ret1 := make([]int, 0, binaryLen)
		for j := 0; j < binaryLen; j++ {
			ret1 = append(ret1, 0)
		}
		ret = append(ret, ret1)
	}

	if upper == 0 && lower == 0 {
		return ret
	}

	for i := 0; i < binaryLen; i++ {
		if colsum[i] == 2 {
			ret[i][0] = 1
			ret[i][1] = 1
			upper -= 1
			lower -= 1
		} else if colsum[i] == 1 {
			if upper > lower {
				ret[0][i] = 1
				upper -= 1
			} else {
				ret[1][i] = 1
				lower -= 1
			}
		}
	}

	return ret
}

func reconstructMatrix1(upper int, lower int, colsum []int) [][]int {
	binaryLen := len(colsum)
	var ret [][]int
	for i := 0; i < 2; i++ {
		ret1 := make([]int, binaryLen)
		ret = append(ret, ret1)
	}

	if upper == 0 && lower == 0 {
		return ret
	}

	for i := 0; i < binaryLen; i++ {
		if colsum[i] == 2 {
			ret[0][i] = 1
			ret[1][i] = 1
			upper -= 1
			lower -= 1
		}
	}

	for i := 0; i < binaryLen; i++ {
		if colsum[i] == 1 {
			if upper > 0 {
				ret[0][i] = 1
				upper -= 1
			} else if lower > 0 {
				ret[1][i] = 1
				lower -= 1
			} else {
				return [][]int{}
			}
		}
	}

	if upper == 0 && lower == 0 {
		return ret
	}

	return [][]int{}
}

func main() {
	upper := 2
	lower := 3
	colsum := []int{2, 2, 1, 1}
	ret := reconstructMatrix1(upper, lower, colsum)
	fmt.Println(ret)
}
