package Offer

func maxProductAfterCutting(length int) int {
	dp := make([]int,length+1)
	for i:=2;i<len(dp);i++{
		for j:=1;j<i;j++{
			dp[i] = max(dp[i],max(j*dp[i-j],j*(i-j)))
		}
	}
	return dp[length]
}

func max(a,b int)int{
	if a>b{
		return a
	}else {
		return b
	}
}


func maxProductAfterCutting(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length ==3 {
		return 2
	}
	time := length/3
	res := 1
	for i:=0;i<time;i++{
		res *= 3
	}
	//fmt.Println(time,res)
	if length%3==1{
		res = res*4/3
	}else if length%3==2{
		res = res*2
	}

	return res
}
