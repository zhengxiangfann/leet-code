package main

import "fmt"

type NumArray struct {
	Buf  []int
	Size int
}

func Constructor(nums []int) NumArray {
	arr := &NumArray{
		Buf:  make([]int, 0),
		Size: 0,
	}
	arr.Buf = append(arr.Buf, nums...)

	arr.Size = len(nums)
	return *arr
}

func (this *NumArray) SumRange(left int, right int) int {
	if this.Size > 0 {
		resp := 0
		for i := left; i <= right; i++ {
			resp += this.Buf[i]
		}
		return resp
	}
	return 0
}

func main() {
	/**
	 * Your NumArray object will be instantiated and called as such:

	 */
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	param_1 := obj.SumRange(0, 2)
	fmt.Println(param_1)
}
