package main

func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum+nums[i], nums[i])
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
