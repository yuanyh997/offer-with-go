package Offer
func firstNotRepeatingChar(s string) byte {
	m := make(map[byte]int)
	for i:=0;i<len(s);i++{
		if _,ok := m[s[i]];ok{
			m[s[i]] = m[s[i]]+1
		}else{
			m[s[i]] = 1
		}
	}
	for i:=0;i<len(s);i++{
		if m[s[i]]==1{
			return s[i]
		}
	}
	return '#'

}