package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

//标准动态规划1
//对每个值为该子序组最后一个值的所有子序组进行筛选，记录最优解
//状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0) ， dp[i] = nums[i] (dp[i-1] ≤ 0)
func maxSubArray(nums []int) int {
	sumMax := nums[0]
	dp := make([]int, len(nums))
	dp[0]=nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]>0{
			dp[i]=nums[i]+dp[i-1]
		}else{
			dp[i]=nums[i]
		}
		if sumMax<dp[i]{
			sumMax=dp[i]
		}
	}
	return sumMax
}
//简化动态规划,数组的每个值都重新定义为以该值为其所在子序组最后一位时，可达到最优子序组的和（以最大和为目标组成的子序组）
func maxSubArray1(nums []int) int {
	sumMax:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i-1]>0{
			nums[i]=nums[i]+nums[i-1]
		}
		if sumMax<nums[i]{
			sumMax =nums[i]
		}
	}
	return sumMax
}
