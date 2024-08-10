func rotate(matrix [][]int) {
	n := len(matrix)
	tempMatrix := make([][]int, n)
	for i := range tempMatrix {
		tempMatrix[i] = make([]int, n)
	}
	for i := range matrix {
		copy(tempMatrix[i], matrix[i])
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[col][n-1-row] = tempMatrix[row][col]
		}
	}
}