package main

import "fmt"

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	return (num-1)%9 + 1
}

func addDigits1(num int) int {
	var sum int
	for {
		sum = 0
		for num > 0 {
			digit := num % 10
			num = num / 10
			sum += digit
		}
		num = sum
		if num/10 == 0 {
			return sum
		}
	}
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	arr := []int{5, 3, 2}
	for _, v := range arr {
		for n%v == 0 {
			n = n / v
		}
	}
	return n == 1
}

func main() {

	ret := isUgly(8)
	fmt.Println(ret)

}
