package Offer

import "sort"

func isContinuous(numbers []int) bool {
	if len(numbers)==0{
		return false
	}
	sort.Ints(numbers)
	cnt := 0
	//flag := true
	for i:=0;i<len(numbers)-1;i++{
		if numbers[i]==0{
			cnt ++
			continue
		}else if numbers[i+1]-numbers[i]>1{
			//flag = false
			//fmt.Println(numbers[i],cnt)
			cnt = cnt-(numbers[i+1]-numbers[i])+1
			//fmt.Println(numbers[i],cnt)
		}else if numbers[i+1] == numbers[i]{
			cnt = -1
			break
		}
	}
	//fmt.Println(cnt)
	return cnt>=0
}