package Offer
func findNumsAppearOnce(nums []int) []int {
	m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok := m[nums[i]];!ok{
			m[nums[i]] = 1
		}else{
			m[nums[i]] += 1
		}
	}
	ans := make([]int,0)
	for i:=0;i<len(nums);i++{
		if m[nums[i]] ==1{
			ans = append(ans,nums[i])
		}
	}
	return ans
}