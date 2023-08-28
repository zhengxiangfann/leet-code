package main

import "fmt"

func shuffle1(nums []int, n int) []int {
	ans := make([]int, n*2, n*2)
	for i := 0; i < n; i++ {
		ans[i*2] = nums[i]
		ans[i*2+1] = nums[i+n]
	}
	return ans
}

func checkIfPangram1(sentence string) bool {
	m := make(map[int32]int)
	count := 0
	for _, word := range sentence {
		if _, ok := m[word]; !ok {
			m[word] = 1
			count += 1
		}
	}
	return count == 26
}

func checkIfPangram3(sentence string) bool {
	state := 0
	for _, word := range sentence {
		state |= 1 << (word - 'a')
	}
	return state == 1<<26-1
}

func checkIfPangram2(sentence string) bool {
	exist := [26]bool{}
	for _, c := range sentence {
		exist[c-'a'] = true
	}
	for _, b := range exist {
		if !b {
			return false
		}
	}
	return true
}

func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i := 0; i < n; i++ {
		ans[i*2] = nums[i]
		ans[i*2+1] = nums[i+n]
	}
	return ans
}

func main() {
	//nums := []int{2, 5, 1, 3, 4, 7}
	//// 2, 5, 1, 3, 4, 7
	//// 2, 3, 5, 4, 1, 7
	//
	//ans := shuffle(nums, 3)
	//fmt.Println(ans)
	sentence := "thequickbrownfoxjumpsoverthelazyydog"
	fmt.Println(checkIfPangram1(sentence))
}
