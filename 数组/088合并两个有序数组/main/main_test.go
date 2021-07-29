package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T){
	//region1
	nums1 := []int{1,2,3,0,0,0}
	m:= 3
	nums2:=[]int{2,5,6}
	n := 3
	//endregion
	merge1(nums1,m,nums2,n)
	fmt.Println()
}