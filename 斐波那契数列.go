package Offer
func Fibonacci(n int) int {
	if n==0{
		return 0
	}
	if n==1 || n==2{
		return 1
	}else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}