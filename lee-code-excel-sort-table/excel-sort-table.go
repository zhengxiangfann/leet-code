package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {

	//l := len(columnTitle) - 1
	//sum := 0
	//for i := 0; i < l+1; i++ {
	//	sum += (int(columnTitle[l-i]-'A') + (i * 26))
	//}
	//return sum

	sum := 0
	for i := 0; i < len(columnTitle); i++ {
		sum = sum*26 + int(columnTitle[i]-'A'+1)
	}
	return sum
}

func convertToTitle1(columnNumber int) string {

	result := ""
	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		columnNumber = columnNumber / 26

		result = string('A'+remainder) + result
	}
	return result
}

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		columnNumber /= 26
		result = string('A'+remainder) + result
	}
	return result
}

func main() {
	in := "ABC"
	out := convertToTitle1(titleToNumber(in))
	fmt.Println(in == out)
}
