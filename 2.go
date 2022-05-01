package main

func missingNumber(nums []int) int {
	var s1, s2, i int

	for i < len(nums) {
		s1 += i
		s2 += nums[i]
		i++
	}
	return s1 - s2 + i
}
