package main

import "fmt"

func main() {
	//board := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}

	var1 := '8'
	var2 := '8'

	fmt.Println(string(var1) + string(var2))

	fmt.Println(string(1) + string(2))

	//isValidSudoku(board)
}

//func isValidSudoku(board [][]byte) bool {
//	sudokuMap := make(map[byte]struct{})
//
//	for _, lineValue := range board {
//		for _, value := range lineValue {
//
//			if value != 46 {
//				_, exists := sudokuMap[]
//
//			}
//
//		}
//	}
//	return false
//}
