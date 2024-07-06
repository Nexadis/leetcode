package main

func passThePillow(n int, time int) int {
	last := time % ((n - 1) * 2)
	if last > n-1 {
		return n - (last - n) - 1
	}
	return 1 + last
}
