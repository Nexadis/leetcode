package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	narr := make([]byte, 0, 10)
	for x > 0 {
		narr = append(narr, byte(x%10))
		x = x / 10
	}
	l := 0
	r := len(narr) - 1
	for l < r {
		if narr[l] != narr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
