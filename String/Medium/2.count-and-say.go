// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	Given a positive integer n, return the nth term of the count-and-say sequence.

	The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
	countAndSay(1) = "1"
	countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is
    then converted into a different digit string.

	To determine how you "say" a digit string, split it into the minimal number of groups so that
    each group is a contiguous section all of the same character. Then for each group, say the number
    of characters, then say the character. To convert the saying into a digit string, replace the
    counts with a number and concatenate every saying.

	Input: n = 1
	Output: "1"
	Explanation: This is the base case.

	Input: n = 4
	Output: "1211"
	Explanation:
	countAndSay(1) = "1"
	countAndSay(2) = say "1" = one 1 = "11"
	countAndSay(3) = say "11" = two 1's = "21"
	countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
*/
package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	str := "11"
	for i := 3; i <= n; i++ {
		str = str + "$"
		slen := len(str)

		count := 1
		tmpStr := ""
		for j := 0; j < (slen - 1); j++ {
			if str[j] == str[j+1] {
				count = count + 1
			} else {
				tmpStr = tmpStr + string(strconv.Itoa(count)) + string(str[j])
				count = 1
			}
		}

		str = tmpStr
	}
	return str
}
func main() {
	var n int = 0
	fmt.Println("Enter the value of n")
	fmt.Scan(&n)

	str := countAndSay(n)
	fmt.Println("The nth element of countAndSay is:- ")
	fmt.Println(str)
}
