package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0)
	lastLeft := intervals[0][0]
	lastRight := intervals[0][1]

	for i := 1; i < n; i++ {
		current := intervals[i]
		// 1. 当前节点的区间 > 前面一个区间
		if lastRight < current[0] {
			ret = append(ret, []int{lastLeft, lastRight})
			lastLeft = current[0]
			lastRight = current[1]
		} else if lastLeft <= current[0] && current[1] >= lastRight {
			// 2. 当前节点的区间 和前面一个区间重合
			lastRight = current[1]
		}
		//else if lastLeft <= current[0] && current[1] < lastRight {
		//	// 3. 当前节点的区间被前一个区间包含
		//
		//}
	}
	ret = append(ret, []int{lastLeft, lastRight})

	return ret
}

func main() {
	//arr := [][]int{{1, 6}, {2, 4}, {5, 10}, {15, 18}}
	arr := [][]int{{1, 4}, {5, 6}}
	merge(arr)
}
