package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets1(s string) string {
	m := map[string]string{
		"1":   "a",
		"2":   "b",
		"3":   "c",
		"4":   "d",
		"5":   "e",
		"6":   "f",
		"7":   "g",
		"8":   "h",
		"9":   "i",
		"10#": "j",
		"11#": "k",
		"12#": "l",
		"13#": "m",
		"14#": "n",
		"15#": "o",
		"16#": "p",
		"17#": "q",
		"18#": "r",
		"19#": "s",
		"20#": "t",
		"21#": "u",
		"22#": "v",
		"23#": "w",
		"24#": "x",
		"25#": "y",
		"26#": "z",
	}

	ans := ""

	for i := 2; i < len(s); i++ {
		if string(s[i]) == "#" {
			tmp := 0
			tmp, _ = strconv.Atoi(s[i-2 : i])
			if tmp > 26 {
				tmp, _ = strconv.Atoi(s[i-1 : i])
			}
			fmt.Println(tmp, s[i-2:i+1])
			ans += m[s[i-2:i+1]]

		} else {

			ans += m[string(s[i-2])] + m[string(s[i-1])] + m[string(s[i])]
		}
	}
	fmt.Println(ans)
	return ans
}

func freqAlphabets3(s string) string {
	m := map[string]string{
		"1":   "a",
		"2":   "b",
		"3":   "c",
		"4":   "d",
		"5":   "e",
		"6":   "f",
		"7":   "g",
		"8":   "h",
		"9":   "i",
		"10#": "j",
		"11#": "k",
		"12#": "l",
		"13#": "m",
		"14#": "n",
		"15#": "o",
		"16#": "p",
		"17#": "q",
		"18#": "r",
		"19#": "s",
		"20#": "t",
		"21#": "u",
		"22#": "v",
		"23#": "w",
		"24#": "x",
		"25#": "y",
		"26#": "z",
	}

	ans := ""
	i := 0

	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			num, _ := strconv.Atoi(s[i : i+2])
			ans += m[strconv.Itoa(num)]
			i += 3 // Skip the number and the '#' symbol
		} else {
			num, _ := strconv.Atoi(string(s[i]))
			ans += m[strconv.Itoa(num)]
			i++
		}
	}

	return ans
}

func freqAlphabets(s string) string {
	ans := ""
	i := 0

	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			num, _ := strconv.Atoi(s[i : i+2])
			ans += string('a' + num - 1)
			i += 3
		} else {
			num, _ := strconv.Atoi(string(s[i]))
			ans += string('a' + num - 1)
			i++
		}
	}

	return ans
}

func main() {
	ans := freqAlphabets("10#11#12")
	fmt.Println(ans)
}
