package main

/*
*
  - Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next

		if k < 0 {
			slow = slow.Next
		}
	}
	return slow.Val
}

func main() {

}
