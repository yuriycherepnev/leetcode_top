/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}

	sum := twoSum(nums, 6)

	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)

	for i, val := range nums {
		_, exists := numbers[val]
		remainder := target - val

		if exists {
			return []int{i, numbers[val]}
		} else {
			numbers[remainder] = i
		}
	}
	return nil
}
