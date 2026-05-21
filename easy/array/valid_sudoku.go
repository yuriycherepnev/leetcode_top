/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'1', '.', '.', '3', '9', '5', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	result := isValidSudoku(board)

	fmt.Println(result)
}

func isValidSudoku(board [][]byte) bool {
	sudokuMap := make(map[uint32]struct{})
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}
			digit := uint32(cell - '0')
			rowKey := uint32(i+1)*10 + digit
			colKey := uint32(j+1)*100 + digit
			boxId := (i/3)*3 + j/3
			boxKey := uint32(boxId+1)*1000 + digit
			if _, exists := sudokuMap[rowKey]; exists {
				return false
			}
			if _, exists := sudokuMap[colKey]; exists {
				return false
			}
			if _, exists := sudokuMap[boxKey]; exists {
				return false
			}
			sudokuMap[rowKey] = struct{}{}
			sudokuMap[colKey] = struct{}{}
			sudokuMap[boxKey] = struct{}{}
		}
	}
	return true
}

/*
простое преобразование числа в строку и назад
digit := 7

str := string('0' + 7)
digit := int('7' - '0')
*/
