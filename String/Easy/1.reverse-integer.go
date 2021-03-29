// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a signed 32-bit integer x, return x with its digits reversed.
    If reversing x causes the value to go outside the signed 32-bit integer
    range [-231, 231 - 1], then return 0.

	Input: x = 123
    Output: 321

	Input: x = -123
    Output: -321

    Input: x = 120
	Output: 21
*/

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	d := 0
	var sign int = -1
	if x > 0 {
		sign = 1
	}

	x = int(math.Abs(float64(x)))
	for x > 0 {
		d = d*10 + (x % 10)
		x = x / 10
	}
	if d > math.MaxInt32 {
		return 0
	}

	return d * sign
}

func main() {
	var num int
	fmt.Println("Enter the integer:- \t")
	fmt.Scan(&num)

	fmt.Println("String before:-")
	fmt.Println(num)

	revNum := reverse(num)
	fmt.Println("String after:-")
	fmt.Println(revNum)
}
