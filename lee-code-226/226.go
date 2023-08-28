package main

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {

		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {

}
