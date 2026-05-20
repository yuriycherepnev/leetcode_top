/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.
*/

package main

func main() {
	digits := []int{2, 1, 0, 3, 0, 12, 9, 67}

	moveZeroes(digits)
}

func moveZeroes(nums []int) {
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++
		}
	}
}
