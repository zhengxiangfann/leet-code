package main

import (
	"fmt"
	"strings"
)

func countMatches1(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	for _, item := range items {
		switch ruleKey {
		case "type":
			if item[0] == ruleValue {
				ans++
			}
		case "color":
			if item[1] == ruleValue {
				ans++
			}
		case "name":
			if item[2] == ruleValue {
				ans++
			}
		}
	}
	return ans
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	var m = map[string]int{"type": 0, "color": 1, "name": 2}
	for _, item := range items {
		if item[m[ruleKey]] == ruleValue {
			ans++
		}
	}
	return ans
}

func getConcatenation(nums []int) []int {
	l := len(nums)
	ans := make([]int, len(nums)*2, len(nums)*2)
	for i := 0; i < 2*l; i++ {
		ans[i] = nums[i%l]
	}
	return ans
}

func interpret(command string) string {
	ans := ""
	l, r := 0, 0
	for i, v := range command {
		if v == 'G' {
			ans += "G"
		} else if v == '(' && l == 0 {
			l = i
		} else if v == ')' {
			r = i
			if r-l == 1 {
				ans += "o"
				r, l = 0, 0
			} else {
				ans += command[l+1 : r]
				r, l = 0, 0
			}
		}
	}
	return ans
}

func countAsterisks1(s string) int {

	newS := strings.Split(s, "|")

	fmt.Println(newS)
	ans := 0
	for i := 0; i < len(newS); i++ {
		if i%2 == 0 {
			for _, v := range newS[i] {

				if v == '*' {
					ans++
				}
			}
		}
	}
	return ans
}

func countAsterisks(s string) int {
	count := 0
	inPair := false
	for _, v := range s {
		if v == '*' && !inPair {
			count++
		}
		if v == '|' {
			inPair = !inPair
		}
	}
	return count
}

func main() {

	ans := countAsterisks("*||||**||*||**|*|**||||**")
	fmt.Println(ans)
	//items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	//ans := countMatches(items, "name", "computer")
	//fmt.Println(ans)
	//
	//nums := []int{1, 2, 1}
	//
	//ret := getConcatenation(nums)
	//
	//fmt.Print(ret)

	//s := "abc"
	//fmt.Println(s[1:2])
	//ret := interpret("(al)G(al)()()G")
	//fmt.Println(ret)
}
