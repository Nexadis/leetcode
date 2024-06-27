package main

func findTheDifference(s string, t string) byte {
	var bs byte
	ls := len(s)
	lt := len(t)

	for i := 0; i < lt; i++ {
		bs += t[i]
	}
	for i := 0; i < ls; i++ {
		bs -= s[i]
	}
	return bs
}
