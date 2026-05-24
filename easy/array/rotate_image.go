package main

import "fmt"

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
	//input: [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
	//output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		j := i
		for j < n {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
			j++
		}
	}

	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			right--
			left++
		}
	}
}
