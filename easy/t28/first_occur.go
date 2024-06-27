package main

func strStr(haystack string, needle string) int {
	ln := len(needle)
	lh := len(haystack)
	for i := 0; i <= lh-ln; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}
