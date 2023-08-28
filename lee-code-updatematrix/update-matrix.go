package main

import (
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	rows, cols := len(mat), len(mat[0])
	dist := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	queue := make([][2]int, 0)

	// 将所有的 0 元素加入队列，并将它们的距离设为 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				dist[i][j] = 0
			}
		}
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				if dist[newRow][newCol] > dist[cell[0]][cell[1]]+1 {
					dist[newRow][newCol] = dist[cell[0]][cell[1]] + 1
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
	}

	return dist
}

func main() {
	//mat = [[0,0,0],[0,1,0],[1,1,1]]
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	newMat := updateMatrix(mat)
	fmt.Println(newMat)
}
