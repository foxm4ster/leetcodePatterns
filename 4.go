package main

func singleNumber(nums []int) int {
	var mask int
	for i := range nums {
		mask ^= nums[i]
	}

	return mask
}
