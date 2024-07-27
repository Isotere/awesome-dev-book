package main

func findNumbers(nums []int) int {
	isEvenDigits := func(i int) bool {
		x, count := 10, 1
		for x <= i {
			x *= 10
			count++
		}

		return count%2 == 0
	}

	counter := 0
	for _, v := range nums {
		if isEvenDigits(v) {
			counter++
		}
	}

	return counter
}
