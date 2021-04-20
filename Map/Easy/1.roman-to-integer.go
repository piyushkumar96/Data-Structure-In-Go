// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
    Topic:- Roman to Integer
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
	Given a roman numeral, convert it to an integer.

	Input: s = "III"
	Output: 3

	Input: s = "LVIII"
	Output: 58

	Input: s = "MCMXCIV"
	Output: 1994
*/

package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var mp map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	slen := len(s)
	for i := 0; i < slen; i++ {
		if (i+1) < slen && mp[string(s[i])] < mp[string(s[i+1])] {
			sum += mp[string(s[i+1])] - mp[string(s[i])]
			i++
			continue
		}
		sum += mp[string(s[i])]
	}

	return sum
}

func main() {
	var str1 string
	fmt.Println("Enter the roman number:- \t")
	fmt.Scan(&str1)

	value := romanToInt(str1)
	fmt.Printf("The integer value of %s is %d", str1, value)
}
