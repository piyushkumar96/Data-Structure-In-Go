// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given an array, rotate the array to the right by k steps, where k is non-negative.

	Note:- Do it in-place with O(1) extra space?

	Input: nums = [1,2,3,4,5,6,7], k = 3
	Output: [5,6,7,1,2,3,4]
	Explanation:
		rotate 1 steps to the right: [7,1,2,3,4,5,6]
		rotate 2 steps to the right: [6,7,1,2,3,4,5]
		rotate 3 steps to the right: [5,6,7,1,2,3,4]

	Input: nums = [-1,-100,3,99], k = 2
	Output: [3,99,-1,-100]
	Explanation:
		rotate 1 steps to the right: [99,-1,-100,3]
		rotate 2 steps to the right: [3,99,-1,-100]
*/

package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func rotate(nums []int, k int) {
	var n = len(nums)
	if n == 1 {
		return
	}
	k = k % n
	k = n - k

	for i := 0; i < gcd(k, n); i++ {
		var temp = nums[i]
		var j = i

		for true {
			var d = j + k
			if d >= n {
				d = d - n
			}

			if d == i {
				break
			}

			nums[j] = nums[d]
			j = d
		}
		nums[j] = temp
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

	var k int
	fmt.Println("Enter the rotate value:- \t")
	fmt.Scan(&k)

	fmt.Println("Array before:-")
	fmt.Println(arr)
	rotate(arr, k)
	fmt.Println("Array after:-")
	fmt.Println(arr)
}
