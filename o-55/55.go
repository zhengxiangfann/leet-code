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

func maxDepth(root *TreeNode) (depth int) {
	if root == nil {
		return
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	depth := maxDepth(tree)
	fmt.Println(depth)
}
