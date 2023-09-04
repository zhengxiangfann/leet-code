package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {

	targetCounts := make(map[int]int)
	arrCounts := make(map[int]int)

	for i := 0; i < len(target); i++ {
		targetCounts[target[i]]++
	}

	fmt.Println(targetCounts)

	for i := 0; i < len(arr); i++ {
		arrCounts[arr[i]]++
	}

	for num, count := range targetCounts {
		if arrCounts[num] != count {
			return false
		}
	}

	return true
}

func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	ans := canBeEqual(target, arr)
	fmt.Println(ans)
}
