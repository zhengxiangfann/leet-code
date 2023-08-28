package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head

	for p1 != nil {
		if k == 0 {
			p2 = p2.Next
		} else {
			k--
		}
		p1 = p1.Next
	}
	return p2
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ans := getKthFromEnd(l, 2)
	fmt.Println(ans.Val)
}
