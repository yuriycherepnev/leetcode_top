package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4}

	concat := getConcatenation(number)

	fmt.Println(concat)
}

func getConcatenation(numbers []int) []int {
	length := len(numbers)

	for i := 0; i < length; i++ {
		numbers = append(numbers, numbers[i])
	}

	return numbers
}
