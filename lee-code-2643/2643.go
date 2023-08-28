package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {

	m := make(map[int]int)

	for i := 0; i < len(mat); i++ {

		cnt := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				cnt++
			}
		}

		if _, ok := m[cnt]; !ok {
			m[cnt] = i
		}
	}

	fmt.Println(m)
	maxOne, idx := 0, 0
	for k, v := range m {
		if k > maxOne {
			maxOne = k
			idx = v
		}
	}
	return []int{idx, maxOne}
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 1}}
	ans := rowAndMaximumOnes(mat)
	fmt.Println(ans)
}
