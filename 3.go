package main

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0, len(nums))

	i := 0
	pos := 0
	for i < len(nums) {
		pos = nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
