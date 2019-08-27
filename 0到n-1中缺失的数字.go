package Offer
func getMissingNumber(nums []int) int {
	for i:=0;i<len(nums);i++{
		if nums[i]!=i{
			return i
		}
	}
	return len(nums)
}

func getMissingNumber(nums []int) int {
	if len(nums)==0{
		return 0
	}
	l,r,mid := 0,len(nums)-1,0
	for l<r{
		mid = (l+r)/2
		if nums[mid] != mid{
			r = mid
		}else{
			l = mid+1
		}
	}
	//fmt.Println(mid,l,r)
	if (nums[r]==r){
		r++
	}
	return r
}