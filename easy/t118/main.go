package main

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	rows[0] = []int{1}
	if numRows == 1 {
		return rows[:numRows]
	}
	rows[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i] = row
	}
	return rows
}
