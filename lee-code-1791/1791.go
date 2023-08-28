package main

import "fmt"

func findCenter(edges [][]int) (ans int) {

	n := len(edges) + 1
	count := make([]int, n+1)

	for _, v := range edges {
		count[v[0]]++
		count[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if count[i] == n-1 {
			return i
		}
	}

	return -1
}

func main() {

	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	ans := findCenter(edges)
	fmt.Println(ans)
}
