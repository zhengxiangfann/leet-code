package main

//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//type MyCircularQueue struct {
//	head, tail     *ListNode
//	capacity, size int
//}
//
//func Constructor(k int) MyCircularQueue {
//	return MyCircularQueue{capacity: k}
//}
//
//func (q *MyCircularQueue) EnQueue(value int) bool {
//	if q.IsFull() {
//		return false
//	}
//	node := &ListNode{Val: value}
//	if q.head == nil {
//		q.head = node
//		q.tail = node
//	} else {
//		q.tail.Next = node
//		q.tail = node
//	}
//	q.size++
//	return true
//}
//
//func (q *MyCircularQueue) DeQueue() bool {
//	if q.IsEmpty() {
//		return false
//	}
//	q.head = q.head.Next
//	q.size--
//	return true
//}
//
//func (q MyCircularQueue) Front() int {
//	if q.IsEmpty() {
//		return -1
//	}
//	return q.head.Val
//}
//
//func (q MyCircularQueue) Rear() int {
//	if q.IsEmpty() {
//		return -1
//	}
//	return q.tail.Val
//}
//
//func (q MyCircularQueue) IsEmpty() bool {
//	return q.size == 0
//}
//
//func (q MyCircularQueue) IsFull() bool {
//	return q.size == q.capacity
//}
//
//func main() {
//	obj := Constructor(3)
//	fmt.Println(obj)
//}
