package Offer

func searchArray(array [][]int, target int) bool {
	for i:=0;i<len(array);i++{
		for j:=0;j<len(array[i]);j++{
			if array[i][j]==target{
				return true
			}
		}
	}
	return false
}