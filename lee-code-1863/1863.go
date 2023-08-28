package main

import "fmt"

func getSubset(nums []int) [][]int {
	subsets := [][]int{{}}

	for _, num := range nums {
		newSubsets := make([][]int, 0, len(subsets)*2)
		for _, subset := range subsets {
			newSubset := append([]int(nil), subset...)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		subsets = append(subsets, newSubsets...)
	}

	return subsets
}

func subsetXORSum(nums []int) int {
	subsets := [][]int{{}}

	for _, num := range nums {
		newSubsets := make([][]int, 0, len(subsets)*2)
		for _, subset := range subsets {
			newSubset := append([]int(nil), subset...)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		subsets = append(subsets, newSubsets...)
	}

	ans := 0
	for _, value := range subsets {
		if len(value) == 0 {
			ans += 0
		} else {
			subAns := 0
			for _, v := range value {
				subAns ^= v
			}
			ans += subAns
		}
	}

	return ans
}

func main() {
	arr := []int{3, 4, 5, 6, 7, 8}
	ret := subsetXORSum(arr)
	fmt.Println(ret)
}
