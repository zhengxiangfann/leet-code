package main

import (
	"fmt"
	"strings"
)

type MovingAverage struct {
	Size   int
	Buffer []int
	Index  int
	Sum    int
	Count  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size:   size,
		Buffer: make([]int, size),
		Index:  0,
		Sum:    0,
		Count:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Count++
	this.Sum += val

	// Calculate the actual window size based on the number of elements added
	windowSize := this.Size
	if this.Count < this.Size {
		windowSize = this.Count
	}

	// Update the buffer and index
	this.Sum -= this.Buffer[this.Index]
	this.Buffer[this.Index] = val
	this.Index = (this.Index + 1) % this.Size

	// Calculate and return the moving average
	return float64(this.Sum) / float64(windowSize)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func generateTheString(n int) string {
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}

func main() {
	c := Constructor(3)

	fmt.Println(c.Next(1))
	fmt.Println(c.Next(10))
	fmt.Println(c.Next(3))
	fmt.Println(c.Next(5))
}
