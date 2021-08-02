package main

func main()  {

}
//返回2个数组里的共有数，重复数取最小重复位
//思路1 使用map记录其中每个数的重复
//时间复杂度map O(m+n)
//空间复杂度 O(min(m,n))
func intersect(nums1 []int, nums2 []int) []int {
return []int{}
}
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1)>len(nums2){
		return intersect(nums2,nums1)
	}
	collection:=make(map[int]int)
	result:=[]int{}
	for _,v:=range nums1{
		collection[v]++
	}
	for _,v2:=range nums2{
		if collection[v2]>0{
			result=append(result,v2)
			collection[v2]--
		}
	}
	return result
}
//思路2 排序+双指针
//两个指针分别指向两个数组的头部。每次比较两个指针指向的两个数组中的数字，
//如果两个数字不相等，则将指向较小数字的指针右移一位，如果两个数字相等，
//将该数字添加到答案，并将两个指针都右移一位。当至少有一个指针超出数组范围时，遍历结束。

func intersect2(nums1 []int, nums2 []int) []int {
}