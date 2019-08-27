package Offer

func replaceSpaces(str string) string {
	l := len(str)
	ans := ""
	for i:=0;i<l;i++{
		if str[i]==' '{
			ans += "%20"
		}else{
			ans += string(str[i])
		}
	}
	return ans
}
