package Offer

import "sort"

func moreThanHalfNum_Solution(nums []int) int {
	if len(nums)==0{
		return -1
	}
	if len(nums)==1{
		return nums[0]
	}
	res := nums[0]
	n :=1
	for i:=1;i<len(nums);i++{
		if nums[i] == res{
			n++
		}else{
			n--
		}
		if n<=0{
			res = nums[i]
			n = 1
		}
	}
	return res
}

func moreThanHalfNum_Solution(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func moreThanHalfNum_Solution(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok := m[nums[i]];!ok{
			m[nums[i]] = 1
		}else{
			m[nums[i]] += 1
			if m[nums[i]]>=len(nums)/2{
				return nums[i]
			}
		}
	}

	return -1
}