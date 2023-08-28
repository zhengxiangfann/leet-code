package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) (ans int) {
	arr := make([]int, 4)
	for i := 3; i >= 0; i-- {
		arr[i] = num % 10
		num = num / 10
	}

	sort.Ints(arr)

	return arr[0]*10 + arr[1]*10 + arr[2] + arr[3]
}

func main() {

	num := 1000
	ans := minimumSum(num)
	fmt.Println(ans)

}
