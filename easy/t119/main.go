package main

func getRow(rowIndex int) []int {
	res := make([]int, 0, rowIndex+1)
	res = append(res, 1)
	prev := 1
	for i := 1; i <= rowIndex; i++ {
		next := prev * (rowIndex - i + 1) / i
		res = append(res, next)
		prev = next
	}
	return res
}
