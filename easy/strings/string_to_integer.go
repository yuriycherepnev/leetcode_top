package main

import (
	"fmt"
	"math"
)

func main() {
	str := "lа"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	//result := myAtoi(str)
	//
	//fmt.Println(result)
}

func myAtoi(s string) int {
	resultNumber := 0
	const MaxInt32 = math.MaxInt32 // 2147483647
	const MinInt32 = math.MinInt32 // 2147483647

	for _, value := range s {
		if value >= 48 && value <= 57 {
			digit := int(value - '0')
			if resultNumber > (MaxInt32-digit)/10 {
				return MaxInt32
			}
			resultNumber = resultNumber*10 + digit
		} else {

		}
	}
	return resultNumber
}
