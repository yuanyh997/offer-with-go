package Offer
func reOrderArray(array []int){
	l := 0
	r := len(array)-1
	for l<r{
		for array[l]%2==1{
			l++
		}
		for array[r]%2==0{
			r--
		}
		if l<r{
			array[l],array[r] = array[r],array[l]
			//fmt.Println(l,r,array)
		}
	}
}