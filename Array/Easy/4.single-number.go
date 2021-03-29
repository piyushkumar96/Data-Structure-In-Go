// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

	Note:- Implement a solution with a linear runtime complexity and without using extra memory.

	Input: nums = [2,2,1]
	Output: 1

	Input: nums = [4,1,2,1,2]
	Output: 4

	Input: nums = [1]
	Output: 1
*/

package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var n = len(nums)
	for i := 1; i < n; i++ {
		// XOR ing of number
		nums[0] = nums[0] ^ nums[i]
	}

	return nums[0]
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

	num := singleNumber(arr)
	fmt.Println("Single number in array:-")
	fmt.Println(num)
}
