package main

func main()  {
	moveZero([]int{0,1,0,3,12})
}
func moveZero(nums []int)  {
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
