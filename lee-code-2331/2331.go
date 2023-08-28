package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree1(root *TreeNode) bool {
	if root != nil {
		if root.Val == 2 {
			return evaluateTree(root.Left) || evaluateTree(root.Right)
		} else if root.Val == 3 {
			return evaluateTree(root.Left) && evaluateTree(root.Right)
		} else if root.Val == 0 {
			return false
		} else if root.Val == 1 {
			return true
		}
	}
	return false
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case 3:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	default:
		return false
	}
}

func main() {
	tree := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}
	ans := evaluateTree(tree)
	fmt.Println(ans)
}
