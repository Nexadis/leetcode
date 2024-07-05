package main

func romanToInt(s string) int {
	sum := 0
	nums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	for i := 0; i < l-1; i++ {
		c := s[i]
		if nums[c] < nums[s[i+1]] {
			sum -= nums[c]
		} else {
			sum += nums[c]
		}
	}
	sum += nums[s[l-1]]

	return sum
}
