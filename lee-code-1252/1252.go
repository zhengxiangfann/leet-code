package main

import "fmt"

func oddCells1(m int, n int, indices [][]int) int {
	mat := make([][]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = 0
		}
	}

	for _, v := range indices {

		for i := 0; i < n; i++ {
			mat[v[0]][i] += 1
		}

		for j := 0; j < m; j++ {
			mat[j][v[1]] += 1
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j]%2 == 1 {
				ans++
			}
		}
	}

	return ans
}

func oddCells(m int, n int, indices [][]int) int {
	rowCounts := make([]int, m)
	colCounts := make([]int, n)
	ans := 0

	for _, v := range indices {
		rowCounts[v[0]]++
		colCounts[v[1]]++
	}

	fmt.Println(rowCounts)
	fmt.Println(colCounts)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (rowCounts[i]+colCounts[j])%2 == 1 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	ans := oddCells(2, 3, [][]int{{0, 1}, {1, 1}})
	fmt.Println(ans)
}
