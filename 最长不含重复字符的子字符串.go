package Offer
//dp f(i)表示以i个字符结尾不包含重复子字符串的最长长度
//1. 如果s[i]未出现过 f(i) = f(i-1)+1
//2. 如果在之前出现过，开始计算s[i]距离上次出现的距离为d
//2.1 如果d<=f(i-1),则说明第i个字符上次出现在f(i-1)对应的不重复字符串之内，那么更新f(i)=d
//2.2 如果d>f(i-1),则无影响,f(i)=f(i-1)+1
func longestSubstringWithoutDuplication(s string) int {
	dp := make([]int,len(s))
	m := make(map[byte]int)
	if len(s)==0{
		return 0
	}
	dp[0] = 1
	m[s[0]]= 0
	res :=1
	for i:=1;i<len(s);i++{
		//fmt.Println(maxLen,s[i])
		if preIndex,ok := m[s[i]];!ok{//未出现
			//fmt.Println(1)
			dp[i] = dp[i-1]+1
		}else if i - preIndex >dp[i-1]{ //d > f(-1)
			dp[i] = dp[i-1]+1
		}else{ //d<= f(i-1)
			dp[i] = i-preIndex
		}
		m[s[i]] = i
		res = max(res,dp[i])
	}
	//fmt.Println(dp)
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}


//dp f(i)表示以i个字符结尾不包含重复子字符串的最长长度
//1. 如果s[i]未出现过 f(i) = f(i-1)+1
//2. 如果在之前出现过，开始计算s[i]距离上次出现的距离为d
//2.1 如果d<=f(i-1),则说明第i个字符上次出现在f(i-1)对应的不重复字符串之内，那么更新f(i)=d
//2.2 如果d>f(i-1),则无影响,f(i)=f(i-1)+1
func longestSubstringWithoutDuplication(s string) int {
	m := make(map[byte]int)
	curLen,maxLen := 0,0
	for i:=0;i<len(s);i++{
		//fmt.Println(maxLen,s[i])
		if preIndex,ok := m[s[i]];!ok{//未出现
			curLen++
		}else if i - preIndex >curLen{ //d > f(-1)
			curLen++
		}else{ //d<= f(i-1)
			if curLen > maxLen{
				maxLen = curLen
			}
			curLen = i-preIndex
		}
		if curLen > maxLen{
			maxLen = curLen
		}
		m[s[i]] = i
	}
	return maxLen
}