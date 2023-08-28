package main

import "fmt"

func numIslands(grid [][]byte) int {
	var count int
	if grid == nil {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				bfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func bfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	n := len(grid)
	m := len(grid[0])
	queue := make([]int, 0)
	code := x*m + y
	queue = append(queue, code)
	for len(queue) > 0 {
		code = queue[0]
		queue = queue[1:]
		i := code / m
		j := code % m

		if i > 0 && grid[i-1][j] == '1' { // 上
			grid[i-1][j] = '0'
			queue = append(queue, (i-1)*m+j)
		}

		if i < n-1 && grid[i+1][j] == '1' { // 下
			grid[i+1][j] = '0'
			queue = append(queue, (i+1)*m+j)
		}

		if j > 0 && grid[i][j-1] == '1' {
			grid[i][j-1] = '0'
			queue = append(queue, i*m+j-1)
		}

		if j < m-1 && grid[i][j+1] == '1' {
			grid[i][j+1] = '0'
			queue = append(queue, i*m+j+1)
		}
	}
}

func bfs1(grid [][]byte, x int, y int) {
	//把当前格子先置为0
	grid[x][y] = '0'
	n := len(grid)
	m := len(grid[0])
	//使用队列，存储的是格子坐标转化的值
	queue := make([]int, 0)
	//我们知道平面坐标是两位数字，但队列中存储的是一位数字，
	//所以这里是把两位数字转化为一位数字
	code := x*m + y
	//坐标转化的值存放到队列中
	queue = append(queue, code)

	for len(queue) > 0 {
		code = queue[0]
		queue = queue[1:]
		//在反转成坐标值（i，j）
		i := code / m
		j := code % m
		if i > 0 && grid[i-1][j] == '1' { //上
			//如果上边格子为1，把它置为0，然后加入到队列中
			//下面同理
			grid[i-1][j] = '0'
			queue = append(queue, (i-1)*m+j)

		}
		if i < n-1 && grid[i+1][j] == '1' { //下
			grid[i+1][j] = '0'
			queue = append(queue, (i+1)*m+j)
		}
		if j > 0 && grid[i][j-1] == '1' { //左
			grid[i][j-1] = '0'
			queue = append(queue, i*m+j-1)

		}
		if j < m-1 && grid[i][j+1] == '1' { //右
			grid[i][j+1] = '0'
			queue = append(queue, i*m+j+1)
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	n := numIslands(grid)
	fmt.Println(n)

	n = numIslands(grid1)
	fmt.Println(n)

}
