package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	ret := false
	visited := make(map[int]bool)
	for {
		visited[n] = true
		nToStr := strconv.Itoa(n)
		sum := 0
		for i := 0; i < len(nToStr); i++ {
			b, _ := strconv.Atoi(string(nToStr[i]))
			sum += (b * b)
		}
		if sum == 1 {
			ret = true
			break
		}
		if _, ok := visited[sum]; ok {
			break
		}
		n = sum
	}
	return ret
}

func isHappy1(n int) bool {
	visited := make(map[int]bool) // 用于记录已访问过的数字
	for n != 1 && !visited[n] {   // 当 n 不等于 1 且 n 未被访问过时执行循环
		visited[n] = true // 将当前数字 n 标记为已访问
		sum := 0
		for n > 0 { // 计算当前数字 n 的每个位上数字的平方和
			digit := n % 10      // 取最低位的数字
			sum += digit * digit // 平方和累加
			n /= 10              // 去掉最低位的数字
		}
		n = sum // 将计算得到的平方和赋值给 n，继续下一轮循环
	}
	return n == 1 // 如果最终 n 等于 1，则为快乐数，返回 true；否则为非快乐数，返回 false
}

func main() {
	fmt.Println(isHappy(100))
}
