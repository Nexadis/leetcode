package main

func climbStairs(n int) int {
	return countSteps(nil, n)
}

func countSteps(steps []int, want int) int {
	if want <= 0 {
		return 0
	}
	if steps == nil {
		steps = make([]int, max(3, want+1))
		steps[1] = 1
		steps[2] = 2
	}

	if steps[want] != 0 {
		return steps[want]
	}

	b2 := countSteps(steps, want-2)
	b1 := countSteps(steps, want-1)
	steps[want] = b1 + b2
	return b1 + b2
}
