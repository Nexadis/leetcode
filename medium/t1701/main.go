package main

func averageWaitingTime(customers [][]int) float64 {
	waits := make([]int, 0, len(customers))
	finish := 0
	for _, c := range customers {
		a, t := c[0], c[1]
		if a >= finish {
			finish = a
		}
		finish += t
		waits = append(waits, finish-a)
	}
	avg := float64(0)
	for _, w := range waits {
		avg += float64(w)
	}
	return avg / float64(len(waits))
}
