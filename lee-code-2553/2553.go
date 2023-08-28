package main

func separateDigits(nums []int) []int {
	ans := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		for nums[i] > 0 {
			digit := nums[i] % 10
			nums[i] = nums[i] / 10
			ans = append([]int{digit}, ans...)
		}
	}
	return ans
}

func main() {
	nums := []int{13, 25, 83, 7}
	separateDigits(nums)
}
