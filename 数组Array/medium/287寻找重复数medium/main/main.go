package main

func main() {
}

//查找数组里的重复数
func findDuplicate(nums []int) int { return 0 }

//暴力解，查找出每个元素出现次数，在查找出现次数>1的
func findDuplicate1(nums []int) int {
	l := len(nums)
	collection := make(map[int]int)
	for i := 0; i < l; i++ {
		if _, ok := collection[nums[i]]; ok {
			return nums[i]
		}
		collection[nums[i]]++
	}

	return 0
}

//todo 1二分
func findDuplicate2(nums []int) int {
	left,right:=0,len(nums)-1
	l:=len(nums)
	for left<right{
		mid:=(left+right)>>1
		counter:=0
		for i:=0;i<l;i++{
			if nums[i]<=mid{
				counter++
			}
		}
		if counter>mid{
			right=mid
		}else{
			left=mid+1
		}
	}
	return left
}

//todo 1快慢指针
func findDuplicate3(nums []int) int { return 0 }