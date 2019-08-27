package Offer

import (
	"sort"
	"strconv"
)

type Cs struct{
	num string
}
type Csslice []Cs
func (slice Csslice) Less(i,j int) bool{

	return slice[i].num+slice[j].num < slice[j].num+slice[i].num
}
func (slice Csslice) Len() int{
	return len(slice)
}

func (slice Csslice) Swap(i,j int){
	slice[i],slice[j] = slice[j],slice[i]
}
func printMinNumber(nums []int) string {
	ans := make(Csslice,len(nums))
	for i:=0;i<len(nums);i++{
		ans[i].num = strconv.Itoa(nums[i])
	}
	sort.Sort(ans)
	//fmt.Println(ans)
	res := ""
	for i:=0;i<len(nums);i++{
		res = res+ans[i].num
	}
	return res
}