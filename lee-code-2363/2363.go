package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	m := make(map[int]int)

	for _, item := range items1 {
		m[item[0]] = item[1]
	}

	for _, item := range items2 {
		m[item[0]] += item[1]
	}

	ans := make([][]int, 0, len(m))
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}

func main() {
	items1 := [][]int{{1, 3}, {2, 2}}
	items2 := [][]int{{7, 1}, {2, 2}, {1, 4}}
	ans := mergeSimilarItems(items1, items2)
	fmt.Println(ans)
}
