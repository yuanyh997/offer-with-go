package Offer

func isPopOrder(pushV []int , popV []int) bool{
	if pushV==nil || popV ==nil || len(pushV)!=len(popV){
		return false
	}
	s := make([]int,0)
	pid := 0
	for i:=0;i<len(popV);i++{
		s = append(s,pushV[i])
		for len(s)!=0 && s[len(s)-1]==popV[pid]{
			s = s[:len(s)-1]
			//fmt.Println(s)
			pid++
		}
	}
	//fmt.Println(s)
	return len(s)==0
}