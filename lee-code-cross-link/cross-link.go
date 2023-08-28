package main

import "fmt"

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}
	var newHead *ListNode
	for headA != nil && headB != nil && newHead == nil {
		if headA.Val == headB.Val {
			newHead = headA
			break
		}
		if headA.Next != nil && headA.Next.Val == headB.Val {
			newHead = headA.Next
			break
		} else if headA.Next != nil && headA.Next.Val != headB.Val {
			headA = headA.Next
		} else {
			headB = headB.Next
		}
	}
	return newHead
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getLength(headA), getLength(headB)

	// 对齐链表的起点
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	// 依次比较节点是否相等
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

// 获取链表的长度
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func main() {
	headA := &ListNode{4, &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}}
	headB := &ListNode{1, headA.Next.Next}
	ret := getIntersectionNode(headA, headB)
	fmt.Println(ret)
}
