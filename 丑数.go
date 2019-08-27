package Offer
func getUglyNumber(n int) int {
	if n<=0{
		return 0
	}
	res := make([]int,n)
	res[0] = 1
	next,minn := 1,0
	u2,u3,u5 := 0,0,0
	for next<n{
		minn = min(res[u2]*2,min(res[u3]*3,res[u5]*5))
		res[next] = minn
		next++
		for res[u2]*2<=minn{
			u2++
		}
		for res[u3]*3<=minn{
			u3++
		}
		for res[u5]*5<=minn{
			u5++
		}
	}
	return res[n-1]
}
func min(a,b int)int{
	if a<b {
		return a
	}else {
		return b
	}
}