package main

import (
	"fmt"
)

func decodeMessage2(key string, message string) string {
	mapKey := make(map[int32]int)
	i := 0
	for _, v := range key {
		if _, ok := mapKey[v]; !ok && v != ' ' {
			mapKey[v] = 'a' + i
			i++
		}
	}

	ans := ""
	for _, v := range message {
		if value, ok := mapKey[v]; ok {
			ans += string(value)
		} else if v == ' ' {
			ans += " "
		}
	}
	return ans
}

func decodeMessage1(key string, message string) string {
	mapKey := make(map[int32]int32)
	var i int32 = 0
	for _, v := range key {
		if _, ok := mapKey[v]; !ok && v != ' ' {
			mapKey[v] = 'a' + i
			i++
		}
	}

	ans := make([]int32, len(message), len(message))
	for j, v := range message {
		if value, ok := mapKey[v]; ok {
			ans[j] = value
		} else if v == ' ' {
			ans[j] = ' '
		}
	}
	return string(ans)
}

func decodeMessage(key string, message string) string {
	mapKey := make(map[int32]int32)
	var i int32 = 0
	for _, v := range key {
		if _, ok := mapKey[v]; !ok && v != ' ' {
			mapKey[v] = 'a' + i
			i++
		}
	}
	mapKey[' '] = ' '
	ans := make([]int32, len(message), len(message))
	for j, v := range message {
		ans[j] = mapKey[v]
	}
	return string(ans)
}

func defangIPaddr(address string) string {
	var ans []int32
	for _, v := range address {
		if v == '.' {
			ans = append(ans, '[', '.', ']')
		} else {
			ans = append(ans, v)
		}

	}
	return string(ans)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil {
		return false
	}
	if root.Right == nil {
		return false
	}

	return root.Val == root.Left.Val+root.Right.Val
}

/*
https://leetcode.cn/submissions/detail/447003216/
*/
func xorOperation(n int, start int) int {
	sum := 0
	for i := 0; i < n; i++ {

		sum = sum ^ (start + 2*i)
	}
	return sum
}

func main() {

	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	//key := "eljuxhpwnyrdgtqkviszcfmabo"
	//message := "zwx hnfx lqantp mnoeius ycgk vcnjrdb"

	ans := decodeMessage(key, message)
	fmt.Println(ans)
}
