package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	pointer := head
	curPointer := head
	for pointer != nil {
		pointer = pointer.Next
		if pointer != nil {
			if curPointer.Val == pointer.Val {
				curPointer.Next = pointer.Next
			} else {
				curPointer = pointer
			}
		}
	}
	return head
}

func main() {
	//list3 := &ListNode{4, nil}
	//list33 := &ListNode{3, list3}
	list22 := &ListNode{1, nil}
	list2 := &ListNode{1, list22}
	//list22 := &ListNode{2, list2}
	list1 := &ListNode{1, list2}

	ret := deleteDuplicates(list1)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
