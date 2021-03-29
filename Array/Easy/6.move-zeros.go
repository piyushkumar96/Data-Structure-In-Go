// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given an integer array nums, move all 0's to the end of it while maintaining
    the relative order of the non-zero elements.

	Note:- that you must do this in-place without making a copy of the array.

	Input: nums = [0,1,0,3,12]
    Output: [1,3,12,0,0]

	Input: nums = [0]
    Output: [0]
*/

package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	var l int = len(nums)
	if l <= 1 {
		return
	}
	k := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			if i != k {
				nums[i] = 0
			}
			k = k + 1
		}
	}
}

func main() {
	var n = 0
	fmt.Println("Enter the size of array")
	fmt.Scan(&n)

	fmt.Println("Enter the array:- \t")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Array before:-")
	fmt.Println(arr)
	moveZeroes(arr)
	fmt.Println("Array after:-")
	fmt.Println(arr)
}
