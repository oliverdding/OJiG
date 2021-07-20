package leetcode

/*
  对角线，左右镜面旋转
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	for row := 0; row < n; row++ {
		for col := 0; col < n-row; col++ {
			matrix[row][col], matrix[n-col-1][n-row-1] = matrix[n-col-1][n-row-1], matrix[row][col]
		}
	}
	for col := 0; col < n; col++ {
		for row := 0; row < n/2; row++ {
			matrix[row][col], matrix[n-row-1][col] = matrix[n-row-1][col], matrix[row][col]
		}
	}
}
