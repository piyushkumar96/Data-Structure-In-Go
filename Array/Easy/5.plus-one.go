// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a non-empty array of decimal digits representing a non-negative integer,
    increment one to the integer.

	The digits are stored such that the most significant digit is at the head of the list,
    and each element in the array contains a single digit.

    You may assume the integer does not contain any leading zero, except the number 0 itself.

	Input: digits = [1,2,3]
	Output: [1,2,4]
	Explanation: The array represents the integer 123.

	Input: digits = [4,3,2,1]
	Output: [4,3,2,2]
	Explanation: The array represents the integer 4321.

	Input: digits = [0]
	Output: [1]
*/

package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var l = len(digits)

	var c = 1
	i := l - 1
	for ; i >= 0; i-- {
		a := digits[i] + 1
		rem := a % 10
		if a <= 9 {
			digits[i] = digits[i] + c
			break
		} else {
			digits[i] = rem
			c = a / 10
		}
	}

	if c > 0 && i == (-1) {
		digits = append([]int{c}, digits...)
	}

	return digits
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

	newNum := plusOne(arr)
	fmt.Println("Number after adding 1 :-")
	fmt.Println(newNum)
}
