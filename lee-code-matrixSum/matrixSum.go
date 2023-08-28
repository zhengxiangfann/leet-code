package main

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	if len(nums) == 1 && len(nums[0]) == 1 {
		return nums[0][0]
	}
	ret := make([]int, 0)
	for len(nums[0]) > 0 {
		scores := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			IDX, maxV := maxIDX(nums[i])
			scores = append(scores, maxV)
			nums[i] = append(nums[i][:IDX], nums[i][IDX+1:]...)
		}
		_, maxScore := maxIDX(scores)
		ret = append(ret, maxScore)
	}
	return sum(ret)
}
func matrixSum1(nums [][]int) int {
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	ret := 0
	for i := 0; i < len(nums[0]); i++ {
		colMax := 0
		for j := 0; j < len(nums); j++ {
			if nums[j][i] > colMax {
				colMax = nums[j][i]
			}
		}
		ret += colMax
	}
	return ret
}

func sum(score []int) int {
	sumScore := 0
	for _, v := range score {
		sumScore += v
	}
	return sumScore
}

func maxIDX(arr []int) (int, int) {
	maxV := arr[0]
	idx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxV {
			maxV = arr[i]
			idx = i
		}
	}
	return idx, maxV
}

func main() {
	nums := [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}
	ret := matrixSum1(nums)
	fmt.Println(ret)
}
