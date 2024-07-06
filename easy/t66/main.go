package main

func plusOne(digits []int) []int {
	r := len(digits) - 1

	for r >= 0 {
		if digits[r] != 9 {
			digits[r]++
			return digits
		}
		digits[r] = 0
		r--
	}
	return append([]int{1}, digits...)
}
