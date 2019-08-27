package Offer
func maxDiff(nums []int) int {
	min := make([]int,len(nums))
	mp := 9999
	for i:=0;i<len(min);i++{
		if nums[i]<mp{
			min[i] = nums[i]
			mp = nums[i]
		}else{
			min[i] = mp
		}
	}
	res := 0
	for i:=0;i<len(min);i++{
		if nums[i]-min[i]>res{
			res = nums[i]-min[i]
		}

	}
	//fmt.Println(min)
	return res
}