package main

func removeDuplicates(nums []int) int {
	l := 0
	r := 0
	for r < len(nums) {
		if nums[l] == nums[r] {
			r++
		} else {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}
