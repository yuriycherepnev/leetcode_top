/*
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Input: nums = [1,2,3,1]
Output: true
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}

	duplicates := containsDuplicate(nums)

	fmt.Println(duplicates)

}

func containsDuplicate(nums []int) bool {
	keyNumbers := make(map[int]bool)

	for _, value := range nums {
		if keyNumbers[value] {
			return true
		}
		keyNumbers[value] = true
	}
	return false
}
