package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "issip"

	number := strStr(haystack, needle)

	fmt.Println(number)
}

func strStr(haystack string, needle string) int {
	needleRune := []rune(needle)
	i := 0

	for index, value := range haystack {
		if value == needleRune[i] {
			i++
		} else {
			i = 0
		}
		if i == len(needleRune) {
			return (index - len(needleRune)) + 1
		}
	}
	return -1
}
