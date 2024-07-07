package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	cnts := make(map[int]int, len(nums2))
	for _, v := range nums2 {
		cnts[v]++
	}
	i := 0
	for _, v := range nums1 {
		if c := cnts[v]; c > 0 {
			nums1[i] = v
			i++
			cnts[v]--
		}
	}
	return nums1[:i]
}
