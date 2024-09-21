package main

/**
PROBLEM:
leetcode -> medium -> Set matrix zeroes

TAGS:
matrix , arrays
**/
func setZeroes(matrix [][]int) {
	colLen := len(matrix[0])
	var rows []int
	var cols []int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				rows = append(rows, row)
				cols = append(cols, col)
			}
		}
	}

	for row := 0; row < len(rows); row++ {
		matrix[rows[row]] = make([]int, colLen)
	}
	for col := 0; col < len(cols); col++ {
		for row := 0; row < len(matrix); row++ {
			matrix[row][cols[col]] = 0
		}
	}

}
