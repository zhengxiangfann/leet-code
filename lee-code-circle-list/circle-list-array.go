package main

//
//import (
//	"fmt"
//)
//
//type MyCircularQueue struct {
//	length int
//	head   int
//	tail   int
//	data   []int
//}
//
//func Constructor(k int) MyCircularQueue {
//	return MyCircularQueue{
//		length: k,
//		head:   -1,
//		tail:   -1,
//		data:   make([]int, k),
//	}
//}
//
//func (this *MyCircularQueue) EnQueue(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//
//	if this.IsEmpty() {
//		this.head = 0
//	}
//
//	this.tail = (this.tail + 1) % this.length
//	this.data[this.tail] = value
//	return true
//}
//
//func (this *MyCircularQueue) DeQueue() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	if this.head == this.tail {
//		this.head = -1
//		this.tail = -1
//		return true
//	}
//	this.head = (this.head + 1) % this.length
//	return true
//}
//
//func (this *MyCircularQueue) Front() int {
//	if this.IsEmpty() {
//		return -1
//	} else {
//		return this.data[this.head]
//	}
//}
//
//func (this *MyCircularQueue) Rear() int {
//	if this.IsEmpty() {
//		return -1
//	} else {
//		return this.data[this.tail]
//	}
//}
//
//func (this *MyCircularQueue) IsEmpty() bool {
//	return this.head == -1
//}
//
//func (this *MyCircularQueue) IsFull() bool {
//	return (this.tail+1)%this.length == this.head
//}
//
//func main01() {
//	obj := Constructor(3)
//	param_1 := obj.EnQueue(1)
//	obj.EnQueue(2)
//	obj.EnQueue(3)
//	fmt.Println(param_1)
//	param_2 := obj.DeQueue()
//	fmt.Println(param_2)
//	param_3 := obj.Front()
//	fmt.Println(param_3)
//	param_4 := obj.Rear()
//	fmt.Println(param_4)
//	param_5 := obj.IsEmpty()
//	fmt.Println(param_5)
//	param_6 := obj.IsFull()
//	fmt.Println(param_6)
//}
