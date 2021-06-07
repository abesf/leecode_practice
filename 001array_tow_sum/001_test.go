package test1

import (
	"testing"
)
func TestTwoSum(t *testing.T)  {
	a:=[]int{2,7,11,15}
	b:=9
	c:=TwoSum(a,b)
	t.Log(c)
}

func TwoSum(nums []int,target int) []int {
	tmp:=make(map[int]int)
	for i,v:=range nums{
		diff:=target-v
		if _,ok:=tmp[diff];ok{
			return []int{i,tmp[diff]}
		}
		tmp[v]=i
	}
	return []int{}
}