package main

import "fmt"

func countDigits(num int) (ans int) {
	newNum := num
	for num > 0 {
		digit := num % 10
		if newNum%digit == 0 && digit != 0 {
			ans++
		}
		num = num / 10
	}
	return ans
}

func subtractProductAndSum(n int) int {
	product := 0
	sum := 0
	for n > 0 {
		digit := n % 10
		product *= digit
		sum += digit
		n = n / 10
	}
	return product - sum
}

func main() {
	ans := countDigits(1248)
	fmt.Println(ans)
}
