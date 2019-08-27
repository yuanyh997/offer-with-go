package Offer

func duplicateInArray(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

func duplicateInArray(nums []int) int {
	m := make(map[int]int)
	if len(nums) == 0 {
		return -1
	}
	for i, _ := range nums {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			return -1
		}
	}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}
