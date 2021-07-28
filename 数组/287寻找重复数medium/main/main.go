package main

func main()  {
	
}
//暴力解，超时了
func findDuplicate(nums []int) int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]==nums[j]{
				return nums[i]
			}
		}
	}
	return 0
}