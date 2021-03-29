// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a string, find the first non-repeating character in it and return
    its index. If it doesn't exist, return -1.

	Input: s = "leetcode"
    Output: 0

	Input: s = "loveleetcode"
    Output: 2
*/

package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	p := [26]int{}
	for i := 0; i < len(s); i++ {
		if p[rune(s[i])-97] == 0 {
			p[rune(s[i])-97] = i + 1
		} else {
			p[rune(s[i])-97] = -1
		}

	}

	min := len(s) + 1
	for i := 0; i < 26; i++ {
		if p[i] > 0 && p[i] < min {
			min = p[i]
		}
	}

	if min == len(s)+1 {
		return -1
	}
	return min - 1
}

func main() {
	var str string
	fmt.Println("Enter the string:- \t")
	fmt.Scan(&str)

	index := firstUniqChar(str)
	fmt.Println("Unique char at index:-")
	fmt.Println(index)
}
