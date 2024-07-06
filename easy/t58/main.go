package main

func lengthOfLastWord(s string) int {
	lastLen := 0
	nowLen := 0
	for _, c := range s {
		if c != ' ' {
			nowLen++
			continue
		}
		if nowLen != 0 {
			lastLen = nowLen
		}
		nowLen = 0
	}
	if nowLen != 0 {
		lastLen = nowLen
	}

	return lastLen
}
