package main

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		_, ok := seen[n]
		if ok {
			return true
		}
		seen[n] = struct{}{}
	}

	return false
}
