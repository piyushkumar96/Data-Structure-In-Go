// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Implement strStr().
	Return the index of the first occurrence of needle in haystack, or -1
    if needle is not part of haystack.

	Note: What should we return when needle is an empty string? This is a great
          question to ask during an interview.

	For the purpose of this problem, we will return 0 when needle is an empty string.
    This is consistent to C's strstr() and Java's indexOf().

	Input: haystack = "hello", needle = "ll"
	Output: 2

	Input: haystack = "aaaaa", needle = "bba"
	Output: -1

	Input: haystack = "", needle = ""
	Output: 0

	Constraints:
		0 <= haystack.length, needle.length <= 5 * 104
		haystack and needle consist of only lower-case English characters.

*/
package main

import (
	"fmt"
)

// Using knuth morris prath Algo for pattern searching
func calLPS(pat string) []int {
	// length of the previous longest prefix suffix
	loplps := 0
	l := len(pat)
	lps := make([]int, l)
	lps[0] = 0
	for i := 1; i < l; {
		if pat[loplps] == pat[i] {
			loplps = loplps + 1
			lps[i] = loplps
			i++
		} else {
			if loplps > 0 {
				loplps = loplps - 1
				loplps = lps[loplps]
			} else {
				loplps = 0
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if nl == 0 {
		return 0
	}
	lps := calLPS(needle)
	j := 0
	for i := 0; i < hl; {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				j = 0
				i++
			}
		}

		if j == nl {
			return i - j
		}
	}

	return -1
}

func main() {
	var str1 string
	fmt.Println("Enter the string:- \t")
	fmt.Scan(&str1)

	var str2 string
	fmt.Println("Enter the sub string for search:- \t")
	fmt.Scan(&str2)

	isSubStr := strStr(str1, str2)
	fmt.Println(isSubStr)
}
