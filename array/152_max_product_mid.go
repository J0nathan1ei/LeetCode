package main

//func maxProduct(nums []int) int {
//	length := len(nums)
//	dp := make([]int, length+1)
//	dp[0],dp[1] = nums[0],
//}
//
//func max(x,y int)int{
//	if x > y {
//		return x
//	}
//	return y
//}

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
