package Offer
func findNumberAppearingOnce(nums []int) int {
	m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok := m[nums[i]];!ok{
			m[nums[i]] = 1
		}else{
			m[nums[i]] += 1
		}
	}

	for i:=0;i<len(nums);i++{
		if m[nums[i]] ==1{
			return nums[i]
		}
	}
	return -1
}
