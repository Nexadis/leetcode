package main

func minOperations(logs []string) int {
	cnt := 0
	for _, l := range logs {
		switch l {
		case "../":
			cnt--
		case "./":
		default:
			cnt++
		}
		if cnt < 0 {
			cnt = 0
		}
	}
	return cnt
}
