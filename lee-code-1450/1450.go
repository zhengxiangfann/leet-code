package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			ans++
		}
	}
	return
}

func main() {
	s := []int{1, 1, 1, 1}
	e := []int{1, 3, 2, 4}
	ans := busyStudent(s, e, 7)
	fmt.Println(ans)
}
