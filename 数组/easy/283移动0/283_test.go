package _83移动0

import "testing"

func TestMoveZero(t *testing.T)  {

}

func moveZeroes(nums []int)  {
  j:=0
  for i,_:=range nums{
  	if nums[i]!=0{
  		nums[j]=nums[i]
  		if i!=j{
  			nums[i]=0
		}
  		j++
	}
  }
}