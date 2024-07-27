package main

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	left, right := 0, len(nums)-1
	sLeft, sRight := -1, -1

	for i := len(nums) - 1; i >= 0; i-- {
		if sLeft < 0 {
			sLeft = nums[left] * nums[left]
		}

		if sRight < 0 {
			sRight = nums[right] * nums[right]
		}

		if left == right {
			result[i] = sLeft
			break
		}

		if sLeft > sRight {
			result[i] = sLeft
			sLeft = -1
			left++
		} else if sRight >= sLeft {
			result[i] = sRight
			sRight = -1
			right--
		}
	}

	return result
}
