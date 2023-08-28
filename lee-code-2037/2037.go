package main

import (
	"fmt"
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i := 0; i < len(students); i++ {
		ans += int(math.Abs(float64(students[i] - seats[i])))
	}
	return ans
}

func sum(nums []int) (ret int) {
	for _, v := range nums {
		ret += v
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func leftRightDifference(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = abs(sum(nums[:i]) - sum(nums[i+1:]))
	}
	return ans
}

func mostWordsFound(sentences []string) int {
	maxCount := 0
	for _, sentence := range sentences {
		count := 0
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == ' ' {
				count++
			}
		}
		count++
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	//seats := []int{4, 1, 5, 9}
	//students := []int{1, 3, 2, 6}
	//ans := minMovesToSeat(seats, students)
	//fmt.Println(ans)

	//nums := []int{10, 4, 8, 3}
	//ans := leftRightDifference(nums)
	//fmt.Println(ans)
	//sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	sentences := []string{"please wait", "continue to fight", "continue to win"}
	ans := mostWordsFound(sentences)
	fmt.Println(ans)

}
