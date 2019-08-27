package Offer

func NumberOf1(n int) int {
	count := 0
	flag := 1
	for flag<=2147483648{
		if n&flag==0{
			count++
			fmt.Println(flag)
		}
		fmt.Println(flag)
		flag = flag<<1
	}
	return count

}
