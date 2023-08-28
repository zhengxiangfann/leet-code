package main

import (
	"fmt"
)

func rotate1(matrix [][]int) {
	row := len(matrix)
	height := len(matrix[0])
	for i := 0; i < row/2-1; i++ {
		for j := 0; j < height/2-1; i++ {
			matrix[i][j], matrix[row][j] = matrix[i][height], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}

func rotate2(matrix [][]int) {
	n := len(matrix)
	// 转置操作
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 逆序操作
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	limit := (n - 1) / 2
	for i := 0; i <= limit; i++ {
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

func main() {

	arr := [][]int{{1, 2}, {3, 4}}
	rotate(arr)
}
