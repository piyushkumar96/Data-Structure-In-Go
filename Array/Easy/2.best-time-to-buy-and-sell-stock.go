// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
	You are given an array prices where prices[i] is the price of a given stock on the ith day.

	Find the maximum profit you can achieve. You may complete as many
	transactions as you like (i.e., buy one and sell one share of the
	stock multiple times).

	Note:- You may not engage in multiple transactions simultaneously
		  (i.e., you must sell the stock before you buy again).

	Input: prices = [7,1,5,3,6,4]
	Output: 7
	Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5),
	profit = 5-1 = 4. Then buy on day 4 (price = 3) and sell on day 5 (price = 6),
	profit = 6-3 = 3.

	Input: prices = [1,2,3,4,5]
	Output: 4
	Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.

	Note:- that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging
           multiple transactions at the same time. You must sell before buying again.
*/
package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var l = len(prices)
	if l <= 1 {
		return 0
	}

	var i = 0
	var maxProfit = 0
	for i < l-1 {
		for i < l-1 && prices[i+1] <= prices[i] {
			i++
		}
		if i == l-1 {
			break
		}

		i++
		maxProfit = maxProfit - prices[i-1]
		for i < l && prices[i] >= prices[i-1] {
			i++
		}
		maxProfit = maxProfit + prices[i-1]
	}
	return maxProfit
}

func main() {
	var n = 0
	fmt.Println("Enter the size of array")
	fmt.Scanln(&n)

	fmt.Println("Enter the prices:- \t")
	var prices = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	profit := maxProfit(prices)
	fmt.Println("The maximum profit:-")
	fmt.Println(profit)
}
