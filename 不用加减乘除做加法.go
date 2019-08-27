package Offer
//  A+B=A^B+(A&B)<<1;
// (1)  A'=A^B;

//       B'=(A&B)<<1;

// (2)  A=A'

//       B=B'

// (3)  判断B是否为0，不为0则重新执行(1)
func add(num1 int, num2 int) int {
	var a,c int
	for num2!=0{
		a = num1^num2
		c = (num1&num2)<<1
		num1 = a
		num2 = c
	}
	return num1
}