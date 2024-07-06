package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
