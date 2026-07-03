package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "issip"

	number := strStr(haystack, needle)

	fmt.Println(number)
}

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	if n == 0 {
		return 0
	}
	if n > h {
		return -1
	}
	for i := 0; i <= h-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
