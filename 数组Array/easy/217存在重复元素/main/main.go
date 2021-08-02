package main

import (
	"sort"
)
func main()  {
	//a:=containsDuplicate([]int{1,2,3,1})
	//a:=containsDuplicate2([]int{1,2,3,1})
	//fmt.Println(a)
}
//1排序后判等 时间复杂度：O(N logN) 空间复杂度 O(logN)
//2set  时间复杂度：O(N)  空间复杂度O(N)
//出现位置
//1 leecode-217
//2 剑指offer03
func containsDuplicate(nums []int) bool {
	return false
}
func containsDuplicate1(nums []int) bool {
	container:=make(map[int]int)
	l:=len(nums)
	for i:=0;i<l;i++{
		if _,ok:=container[i];ok{
			return true
		}
		container[i]=nums[i]
	}
	return false
}
func containsDuplicate2(nums []int) bool {

	l:=len(nums)
	sort.Ints(nums)
	for i:=1;i<l;i++{
		if nums[i]==nums[i-1]{
			return true
		}
	}
	return false
}