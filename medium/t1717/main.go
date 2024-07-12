package main

func maximumGain(s string, x int, y int) int {
	score := 0
	step := 1
	start := 0
	stop := len(s) - 1

	if x < y {
		step = -1
		start, stop = stop, start
		x, y = y, x
	}

	aCount := 0
	bCount := 0

	for i := start; i != stop+step; i += step {
		c := s[i]
		switch c {
		case 'a':
			aCount++
		case 'b':
			if aCount > 0 {
				aCount--
				score += x
			} else {
				bCount++
			}
		default:
			score += y * min(aCount, bCount)
			aCount = 0
			bCount = 0
		}

	}
	score += y * min(aCount, bCount)

	return score
}
