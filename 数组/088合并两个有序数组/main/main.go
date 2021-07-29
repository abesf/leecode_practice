package main

import "sort"
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
func main(){

}
//暴力解法，合并后排序
func merge(nums1 []int, m int, nums2 []int, n int)  {
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}
//指针+中间值，为2个数组各自定义一个指针,判定每个数组所在指针值，按序插入新数组
//最后拷贝一下
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	sorted:=make([]int,m+n)
	p1:=0
	p2:=0
	for {
		if p1==m{
			sorted=append(sorted,nums2[p2:]...)
			break
		}
		if p2==n{
			sorted=append(sorted,nums1[p1:]...)
			break
		}
		if nums1[p1]<nums2[p2]{
			sorted=append(sorted,nums1[p1])
			p1++
		}else{
			sorted=append(sorted,nums2[p1])
			p2++
		}
	}
	copy(nums1,sorted)
}