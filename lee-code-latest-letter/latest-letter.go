package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {

	if target >= letters[len(letters)-1] || target < letters[0] {
		return letters[0]
	}

	left, right := 0, len(letters)-1
	for left-right <= 1 {
		mid := (right + left) / 2
		//if letters[mid] == target {
		//	return letters[mid+1]
		//} else
		if left-right == 1 {
			return letters[left]
		} else if letters[mid] > target {
			right = mid - 1
		} else if letters[mid] <= target {
			left = mid + 1
		}
	}
	return letters[0]
}

func main() {
	letters := []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}
	ret := nextGreatestLetter(letters, 'e')
	fmt.Println(string(ret))
}
