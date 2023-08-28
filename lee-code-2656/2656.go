package main

/*
m+(m+1)+(m+2)+⋯+(m+k−1)=
2
(2m+k−1)⋅k
​

*/

func maximizeSum(nums []int, k int) int {
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	//(max*2 + k - 1) * k / 2

	//max + 0
	//max + 1
	//max + 2
	//max + (k -1)
	// ans  = (max + max+k-1) * k /2

	ans := 0
	for i := maxNum; i < maxNum+k; i++ {
		ans += i
	}
	return ans
}

func main() {

}
