package Offer

func Power(base float64, exponent int) float64 {
	ans := 1.0
	if exponent>=0{
		for i:=0;i<exponent;i++{
			ans *= base
		}
	}
	if exponent<0{
		exponent = -exponent
		for i:=0;i<exponent;i++{
			ans /= base
		}
	}
	return ans
}