package main

func max(nums []int) int {
	ans := 0
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxCandies := max(candies)

	ans := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCandies {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans

}

func sumOfSquares(nums []int) int {
	ans := 0
	n := len(nums)
	for i, v := range nums {
		if n%(i+1) == 0 {
			ans += v * v
		}
	}
	return ans
}

func main() {
	//num := []int{12, 1, 12}
	//ans := kidsWithCandies(num, 10)
	//fmt.Println(ans)

}
