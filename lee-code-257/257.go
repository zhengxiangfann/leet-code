package main

import "fmt"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []int {

	ans := make([]int, 0)

	if root != nil {
		fmt.Println(root.Val)
		ans = append(ans, root.Val)
	}

	if root.Left != nil {
		ans = append(ans, binaryTreePaths(root.Left)...)
	}

	if root.Right != nil {
		ans = append(ans, binaryTreePaths(root.Right)...)
	}
	fmt.Println(ans)
	return ans
}

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, &TreeNode{8, nil, nil}}}
	binaryTreePaths(tree)
}
