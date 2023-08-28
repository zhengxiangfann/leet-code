package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, row, col, oldColor, newColor int) {
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) || image[row][col] != oldColor {
		return
	}
	image[row][col] = newColor
	dfs(image, row-1, col, oldColor, newColor) // 上
	dfs(image, row+1, col, oldColor, newColor) // 下
	dfs(image, row, col-1, oldColor, newColor) // 左
	dfs(image, row, col+1, oldColor, newColor) // 右
}

func main() {
	//image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	newImage := floodFill(image, 0, 0, 2)
	fmt.Println(newImage)
}
