package Offer
func getNumberSameAsIndex(nums []int) int {
	for i:=0;i<len(nums);i++{
		if i==nums[i]{
			return i
		}
	}
	return -1
}