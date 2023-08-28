package main

import "strconv"

/*
给你一个下标从 0开始的字符串details。details中每个元素都是一位乘客的信息，信息用长度为 15的字符串表示，表示方式如下：

前十个字符是乘客的手机号码。
接下来的一个字符是乘客的性别。
接下来两个字符是乘客的年龄。
最后两个字符是乘客的座位号。
请你返回乘客中年龄 严格大于 60 岁的人数。

示例 1：

输入：details = ["7868190130M7522","5303914400F9211","9273338290F4010"]
输出：2
解释：下标为 0 ，1 和 2 的乘客年龄分别为 75 ，92 和 40 。所以有 2 人年龄大于 60 岁。
*/
func countSeniors(details []string) (ans int) {
	for _, v := range details {
		if age, _ := strconv.Atoi(v[len(v)-4 : len(v)-2]); age > 60 {
			ans++
		}
	}
	return
}

func main() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	countSeniors(details)
}
