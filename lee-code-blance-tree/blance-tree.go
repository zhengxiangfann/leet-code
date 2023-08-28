package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 是否平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := (right + left) / 2
		fmt.Println(mid, nums[mid])
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	//tree := TreeNode{
	//	1, nil, nil,
	//}
	//ret := isBalanced(&tree)
	//fmt.Println(ret)

	arr := []int{1, 2, 5, 9, 88, 190, 300, 400, 789, 9876, 9999, 98765}
	ret := search(arr, 1)
	fmt.Println(ret)

}
