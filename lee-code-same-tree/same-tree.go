package main

import "fmt"

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 垃圾写法
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	} else if p.Left == nil && q.Left != nil {
		return false
	} else if p.Left != nil && q.Left == nil {
		return false
	} else if p.Right == nil && q.Right != nil {
		return false
	} else if p.Right != nil && q.Right == nil {
		return false
	} else if p.Left != nil && q.Left != nil {
		return isSameTree(p.Left, q.Left)
	} else if p.Right != nil && q.Right != nil {
		return isSameTree(p.Right, q.Right)
	} else if p.Val == q.Val && p.Left == nil && q.Left == nil && p.Right == nil && q.Right == nil {
		return true
	} else {
		return false
	}
}

// chat-GPT的回答
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{
		1, nil, &TreeNode{3, nil, nil},
	}
	q := &TreeNode{
		1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil},
	}

	ret := isSameTree(p, q)
	fmt.Println(ret)
}
