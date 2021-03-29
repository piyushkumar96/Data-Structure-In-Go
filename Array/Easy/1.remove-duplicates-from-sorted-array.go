// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a sorted array nums, remove the duplicates in-place such that each
	element appears only once and returns the new length.
		Input: nums = [1,1,2]
		Output: 2, nums = [1,2]
		Explanation: Your function should return length = 2, with the first
		two elements of nums being 1 and 2 respectively. It doesn't matter what
		you leave beyond the returned length.

		Input: nums = [0,0,1,1,1,2,2,3,3,4]
		Output: 5, nums = [0,1,2,3,4]
*/
package main

import (
	"fmt"
)

func removeDuplicates(arr []int) int {
	larr := len(arr)
	if larr == 0 || larr == 1 {
		return larr
	}

	c := 1
	for i := 0; i < larr-1; i++ {
		if arr[i] != arr[i+1] {
			arr[c] = arr[i+1]
			c++
		}
	}
	return c
}

func main() {
	var n = 0
	fmt.Println("Enter the size of array")
	fmt.Scanln(&n)

	fmt.Println("Enter array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("The number of elements in array before:-")
	fmt.Println(len(arr))
	c := removeDuplicates(arr)
	fmt.Println("The number of elements in array after:-")
	fmt.Println(c)
}
