package main

import (
	"fmt"
)

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func intToRoman(num int) (ans string) {

	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	for num > 0 {
		maxRoman := 1
		for v := range romanMap {
			if v > maxRoman && num >= v {
				maxRoman = v
			}
		}
		
		ans = ans + romanMap[maxRoman]
		num = num - maxRoman
	}

	return
}

func main() {
	ans := intToRoman(1500)
	fmt.Println(ans)
}
