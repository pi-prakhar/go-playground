package main

import "fmt"

/**
Problem:
leetcode -> medium -> Spiral Matrix

Tags:
array, matrix, simulation
**/

func spiralOrder(matrix [][]int) []int {
	var res []int
	row := len(matrix)
	col := len(matrix[0])
	solve(row-1, col-1, 0, 0, &matrix, &res, 'R')
	fmt.Println(res)
	return res
}

func solve(row int, col int, crow int, ccol int, matrix *[][]int, res *[]int, dir rune) {
	if len(*res) == ((row + 1) * (col + 1)) {
		return
	}

	switch dir {
	case 'R':
		if (ccol <= col) && ((*matrix)[crow][ccol] != -101) {
			fmt.Println((*matrix)[crow][ccol])
			(*res) = append((*res), (*matrix)[crow][ccol])
			(*matrix)[crow][ccol] = -101
			ncol := ccol + 1
			solve(row, col, crow, ncol, matrix, res, dir)
		} else {
			ncol := ccol - 1
			nrow := crow + 1
			solve(row, col, nrow, ncol, matrix, res, 'D')
		}
	case 'D':
		if (crow <= row) && ((*matrix)[crow][ccol] != -101) {
			fmt.Println((*matrix)[crow][ccol])
			(*res) = append((*res), (*matrix)[crow][ccol])
			(*matrix)[crow][ccol] = -101
			nrow := crow + 1
			solve(row, col, nrow, ccol, matrix, res, dir)
		} else {
			nrow := crow - 1
			ncol := ccol - 1
			solve(row, col, nrow, ncol, matrix, res, 'L')
		}
	case 'L':
		if (ccol >= 0) && ((*matrix)[crow][ccol] != -101) {
			fmt.Println((*matrix)[crow][ccol])
			(*res) = append((*res), (*matrix)[crow][ccol])
			(*matrix)[crow][ccol] = -101
			ncol := ccol - 1
			solve(row, col, crow, ncol, matrix, res, dir)
		} else {
			ncol := ccol + 1
			nrow := crow - 1
			solve(row, col, nrow, ncol, matrix, res, 'U')
		}
	case 'U':
		if (crow >= 0) && ((*matrix)[crow][ccol] != -101) {
			fmt.Println((*matrix)[crow][ccol])
			(*res) = append((*res), (*matrix)[crow][ccol])
			(*matrix)[crow][ccol] = -101
			nrow := crow - 1
			solve(row, col, nrow, ccol, matrix, res, dir)
		} else {
			nrow := crow + 1
			ncol := ccol + 1
			solve(row, col, nrow, ncol, matrix, res, 'R')
		}
	default:
		break
	}
}
