package main

func main()  {

}
//寻找数组中一个数据下标，使得下标左侧之和等于下标右侧之和，多值时返回最左侧,是最
// 左或最右时返回0,不存在返回-1
func pivotIndex(nums []int) int  {
	return -1
}
//时间复杂度：O(n)，其中 n 为数组的长度。
//空间复杂度：O(1)。
//左侧数据left=num[0]+nums[1]+....+nums[i-1]
//总和total
//右侧数据等于左侧的数据right=num[l-i-1]+nums[l-i-2]+...nums[i+1]
//判等left+right+nums[i]=total;left=right
//2*left+nums[i]=total存在数据
func pivotIndex1(nums []int) int  {
	total:=0
	l:=len(nums)
	for i:=0;i<l;i++{
		total+=nums[i]
	}
	left:=0
	for i:=0;i<l;i++{
		if nums[i]+2*left==total{
			return i
		}
		left+=nums[i]
	}
	return -1
}
type Solution interface {
	//region724 easy
    pivotIndex(nums []int) int
	//endregion
}