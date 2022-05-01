package main

import (
	"log"
)

func singleNumber(nums []int) int {
	var mask int
	for i := range nums {
		mask ^= nums[i]
	}

	return mask
}

func main() {
	log.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
