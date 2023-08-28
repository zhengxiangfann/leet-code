package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * Definition for a Node.
 */

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {

	if root == nil {
		return nil
	}

	ans := []int{}

	for _, node := range root.Children {
		ans = append(ans, postorder(node)...)
	}
	ans = append(ans, root.Val)
	return ans
}

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	return nums[1]
}

func replaceDigits(s string) string {
	var buf strings.Builder
	for i := 1; i < len(s); i += 2 {
		buf.WriteByte(s[i-1])
		fmt.Println(string(s[i-1]))
		buf.WriteByte((s[i-1] + s[i] - 48))
		fmt.Println(string(s[i-1] + s[i] - 48))
	}
	if len(s)%2 == 1 {
		buf.WriteByte(s[len(s)-1])
	}
	return buf.String()
}

func main() {
	ans := replaceDigits("a1b2c3d4e")
	fmt.Println(ans)
}
