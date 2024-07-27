package max_consecutive_ones

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Constraints:

//    1 <= nums.length <= 105
//    nums[i] is either 0 or 1.

// findMaxConsecutiveOnes решение с помощью окошка
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	left := 0

	for i, v := range nums {
		// если текущая цифра 0
		if v == 0 {
			if i-left > max {
				max = i - left
			}
			// ставим на следующую позицию, предполагая, что там 1
			left = i + 1
		}

		if len(nums)-1 == i {
			if len(nums)-left > max {
				max = len(nums) - left
			}
		}
	}

	return max
}

// findMaxConsecutiveOnes2 решение простое с помощью каунтеров
func findMaxConsecutiveOnes2(nums []int) int {
	max, set := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			set++
		} else {
			set = 0
		}
		if set > max {
			max = set
		}
	}
	return max
}
