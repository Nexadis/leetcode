package main

func threeConsecutiveOdds(arr []int) bool {
	odds := 0
	for i := 0; i < len(arr); i++ {
		if odds == 3 {
			return true
		}
		if arr[i]%2 == 0 {
			odds = 0
		} else {
			odds++
		}
	}
	return odds == 3
}
