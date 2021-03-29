// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	Input: s = "anagram", t = "nagaram"
	Output: true

	Input: s = "rat", t = "car"
	Output: false
*/

package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	var sl int = len(s)
	var tl int = len(t)

	if sl != tl {
		return false
	}

	arr := [26]int{}

	for i := 0; i < sl; i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	//fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	var str1 string
	fmt.Println("Enter the string1:- \t")
	fmt.Scan(&str1)

	var str2 string
	fmt.Println("Enter the string2:- \t")
	fmt.Scan(&str2)

	isAnagm := isAnagram(str1, str2)
	fmt.Println(isAnagm)
}
