package main

func removeElement(nums []int, val int) int {
	l := 0
	for _, n := range nums {
		if n == val {
			continue
		}
		nums[l] = n
		l++
	}
	return l
}
