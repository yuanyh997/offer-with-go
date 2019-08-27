package Offer
func getNumberOfK(nums []int , k int) int{
	if len(nums)==0{
		return 0
	}
	ans:=0
	l:= 0
	r:= len(nums)-1
	mid := 0
	for l<r{
		mid = (l+r)/2
		//fmt.Println(mid)
		if nums[mid]>k{
			r = mid
		}else if nums[mid]<k{
			l = mid+1
		}else{
			break
		}
	}
	l = mid
	r = mid+1
	for l>=0 && nums[l]==k{
		ans++
		l--
	}
	for r<len(nums) && nums[r]==k{
		ans++
		r++
	}
	return ans
}

func getNumberOfK(nums []int , k int) int{
	ans:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==k{
			ans++
		}
		if nums[i]>k{
			break
		}
	}
	return ans
}