package Offer
func findNumbersWithSum(nums []int, target int) []int {
	ans := make([]int,0)
	m := make(map[int]int)
	for _,i := range nums{
		if _,ok:=m[target-i];ok{
			ans = append(ans,i)
			ans = append(ans,target-i)
		}else{
			m[i] = target-i
		}
	}
	return ans
}