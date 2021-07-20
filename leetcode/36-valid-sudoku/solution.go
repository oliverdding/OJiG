package leetcode

/*
  看了题解，最重点的一个式子：
  box = (row/3)*3+col/3
  其实也就是空间换时间了
*/
func isValidSudoku(board [][]byte) bool {
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == '.' {
				continue
			}
			value := board[x][y] - '1'
			if rows[y][value] == true {
				return false
			}
			rows[y][value] = true
			if cols[x][value] == true {
				return false
			}
			cols[x][value] = true
			if boxes[(y/3)*3+x/3][value] == true {
				return false
			}
			boxes[(y/3)*3+x/3][value] = true
		}
	}
	return true
}
