package main

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
	ret := make([]int64, 0)
	if finalSum%2 == 1 {
		return ret
	}
	var startNum int64 = 2
	for finalSum >= startNum {
		if finalSum-startNum > startNum {
			ret = append(ret, startNum)
			finalSum = finalSum - startNum
			startNum += 2
		} else {
			ret = append(ret, finalSum)
			break
		}
	}
	return ret
}

func main() {
	ret := maximumEvenSplit(4)
	fmt.Println(ret)
}
