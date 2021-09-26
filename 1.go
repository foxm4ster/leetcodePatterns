package main

func ContainsDuplicate(nums []int) bool {
	seen := make([]int, 0, len(nums))

	for _, i := range nums {
		if inArray(&i, &seen) {
			return true
		}
		seen = append(seen, i)
	}

	return false
}

func inArray(n *int, seen *[]int) bool {
	for _, i := range *seen {
		if *n == i {
			return true
		}
	}

	return false
}
