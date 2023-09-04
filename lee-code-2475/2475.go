package main

import "fmt"

func unequalTriplets(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				for k := j + 1; k < len(nums); k++ {

					if nums[j] != nums[k] && nums[i] != nums[k] {
						ans++
					}
				}
			}

		}
	}
	return ans
}

func main() {
	nums := []int{4, 4, 2, 4, 3}
	ans := unequalTriplets(nums)
	fmt.Println(ans)
}
