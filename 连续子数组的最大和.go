package Offer

import "sort"

func maxSubArray(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	dp := make([]int,len(nums)+1)
	for i:=1;i<len(nums)+1;i++{
		for j:=0;j<i;j++{
			dp[i] = max(nums[j],dp[i-1]+nums[j])
		}
		//fmt.Println(dp)
	}
	dp = dp[1:]
	sort.Ints(dp)
	return dp[len(nums)-1]
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}