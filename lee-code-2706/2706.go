package main

import "fmt"

func buyChoco(prices []int, money int) int {

	if len(prices) < 2 {
		return money
	}

	smallest1, smallest2 := 100, 100

	for _, price := range prices {
		if price < smallest1 {
			smallest2 = smallest1
			smallest1 = price
		} else if price < smallest2 {
			smallest2 = price
		}
	}

	fmt.Println(smallest1, smallest2)

	m := money - smallest1 - smallest2

	if m >= 0 {
		return m
	} else {
		return money
	}
}

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	if root == nil {
		return 0
	}

	colors := make(map[int]struct{})

	colors[root.Val] = struct{}{}
	countColors(root.Left, colors)
	countColors(root.Right, colors)

	return len(colors)
}

func countColors(root *TreeNode, colors map[int]struct{}) {
	if root == nil {
		return
	}

	colors[root.Val] = struct{}{}
	countColors(root.Left, colors)
	countColors(root.Right, colors)
}

func main() {
	p := []int{98, 54, 6, 34, 66, 63, 52, 39}
	ans := buyChoco(p, 62)
	fmt.Println(ans)
}
