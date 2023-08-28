package main

import "fmt"

/*
给定一个整数数组temperatures，
表示每天的温度，返回一个数组answer，
其中answer[i]是指对于第 i 天，
下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用0 来代替。
*/

func dailyTemperatures1(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return res
}

func dailyTemperatures2(T []int) []int {
	n := len(T)
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		j := i + 1
		for j < n {
			if T[j] > T[i] {
				res[i] = j - i
				break
			} else if res[j] == 0 {
				break
			} else {
				j += res[j]
			}
		}
	}

	return res
}

func main() {
	temperatures := []int{30, 60, 90}
	ret := dailyTemperatures2(temperatures)
	fmt.Println(ret)
}
