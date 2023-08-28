package main

func minCount(coins []int) (ans int) {

	for _, v := range coins {

		if v%2 == 0 {
			ans += v / 2
		} else {
			ans += v/2 + 1
		}

	}
	return ans
}

func main() {

}
