package main

import "fmt"

func main() {
	number := 10001

	result := isPalindrome(number)

	fmt.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	divider := 1

	for temp := x / 10; temp != 0; temp /= 10 {
		divider *= 10
	}

	for x > 0 {
		first := x / divider
		last := x % 10

		if first != last {
			return false
		}

		x = (x % divider) / 10
		divider /= 100
	}

	return true
}
