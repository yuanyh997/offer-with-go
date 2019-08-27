package Offer

import "sort"

func getLeastNumbers_Solution(input []int , k int) []int{
	sort.Ints(input)
	return input[:k]
}